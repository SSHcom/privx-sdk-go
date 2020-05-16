//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package oauth

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/SSHcom/privx-sdk-go/pkce"
)

// Provider is a source for access token to api client
type Provider interface {
	Token() (string, error)
}

const (
	UserAgent = "privx-sdk-go"
)

type Config struct {
	ClientID        string `toml:"oauth_client_id"`
	ClientSecret    string `toml:"oauth_client_secret"`
	RedirectURI     string `toml:"oauth_redirect_uri"`
	APIClientID     string `toml:"api_client_id"`
	APIClientSecret string `toml:"api_client_secret"`
}

type Client struct {
	endpoint       string
	m              *sync.Mutex
	c              *sync.Cond
	authPending    bool
	config         Config
	http           *http.Client
	httpNoRedirect *http.Client
	token          *AccessToken
}

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    uint   `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	notAfter     time.Time
}

func NewClient(config Config, endpoint string, cert *x509.Certificate,
	verbose bool) (*Client, error) {

	tlsConfig := &tls.Config{}
	if cert != nil {
		pool := x509.NewCertPool()
		pool.AddCert(cert)
		tlsConfig.RootCAs = pool
	}

	m := new(sync.Mutex)

	return &Client{
		m:        m,
		c:        sync.NewCond(m),
		config:   config,
		endpoint: endpoint,
		http: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		},
		httpNoRedirect: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		},
	}, nil
}

// Token returns an OAuth token. If the client does not have a valid
// token, it will do the authentication flow to get a fresh token.
func (client *Client) Token() (string, error) {
	client.m.Lock()
	for client.authPending {
		client.c.Wait()
	}
	defer client.m.Unlock()
	now := time.Now()
	if client.token == nil || now.After(client.token.notAfter) {
		client.authPending = true
		client.token = nil
		client.m.Unlock()

		var token *AccessToken
		var err error

		if false {
			token, err = client.authorizationCodeGrant()
		} else {
			token, err = client.resourceOwnerPasswordCredentialsGrant()
		}

		client.m.Lock()
		client.authPending = false
		client.c.Broadcast()
		if err != nil {
			return "", err
		}
		client.token = token
	}
	return client.token.AccessToken, nil
}

func (client *Client) resourceOwnerPasswordCredentialsGrant() (
	*AccessToken, error) {

	form := url.Values{}
	form.Add("grant_type", "password")
	form.Add("username", client.config.APIClientID)
	form.Add("password", client.config.APIClientSecret)

	tokenURL := client.tokenURL()

	req, err := http.NewRequest(http.MethodPost, tokenURL,
		strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(client.config.ClientID, client.config.ClientSecret)

	resp, err := client.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		err = decodeError("", body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("unexpected response code: %s", resp.Status)
	}

	token := new(AccessToken)
	err = json.Unmarshal(body, token)
	if err != nil {
		return nil, err
	}
	token.notAfter = time.Now().Add(
		time.Duration(token.ExpiresIn) * time.Second)

	return token, nil
}

func (client *Client) authorizationCodeGrant() (*AccessToken, error) {

	codeVerifier, err := pkce.NewCodeVerifier()
	if err != nil {
		return nil, err
	}
	challenge, method := codeVerifier.ChallengeS256()

	// Random state.
	var buf [32]byte
	_, err = rand.Read(buf[:])
	if err != nil {
		return nil, err
	}
	state := base64.RawURLEncoding.EncodeToString(buf[:])

	token, err := client.authorizationEndpoint(challenge, method, state)
	if err != nil {
		return nil, err
	}

	code, st, err := client.login(token)
	if err != nil {
		return nil, err
	}
	if st != state {
		return nil, errors.New("invalid response state")
	}

	return client.tokenEndpoint(code, codeVerifier)
}

func (client *Client) authorizationURL(q string) string {
	return fmt.Sprintf("%s/auth/api/v1/oauth/authorize?%s", client.endpoint, q)
}

func (client *Client) tokenURL() string {
	return fmt.Sprintf("%s/auth/api/v1/oauth/token", client.endpoint)
}

func (client *Client) loginURL() string {
	return fmt.Sprintf("%s/auth/api/v1/login", client.endpoint)
}

func (client *Client) authorizationEndpoint(challenge, method, state string) (
	token string, err error) {

	// Call authorization endpoint.

	form := url.Values{}
	form.Add("response_type", "code")
	form.Add("client_id", client.config.ClientID)
	form.Add("redirect_uri", client.config.RedirectURI)
	form.Add("state", state)
	form.Add("user_agent", UserAgent)
	form.Add(pkce.ParamCodeChallenge, challenge)
	form.Add(pkce.ParamCodeChallengeMethod, method)

	authorizationURL := client.authorizationURL(form.Encode())

	req, err := http.NewRequest(http.MethodGet, authorizationURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.httpNoRedirect.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusTemporaryRedirect {
		err = decodeError(resp.Header.Get("Location"), body)
		if err != nil {
			return
		}
		return "", fmt.Errorf("unexpected response code: %s", resp.Status)
	}

	loginURL := resp.Header.Get("Location")
	if len(loginURL) == 0 {
		return "",
			errors.New("redirection URI missing from authorization response")
	}

	// Parse login token.
	parsed, err := url.Parse(loginURL)
	if err != nil {
		return "", err
	}
	token = parsed.Query().Get("token")
	if len(token) == 0 {
		return "", errors.New("no login token in authorization response")
	}

	return
}

func (client *Client) tokenEndpoint(code string, verifier pkce.CodeVerifier) (
	*AccessToken, error) {

	// Call token endpoint

	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("code", code)
	form.Add("redirect_uri", client.config.RedirectURI)
	form.Add("client_id", client.config.ClientID)
	form.Add(pkce.ParamCodeVerifier, verifier.String())

	tokenURL := client.tokenURL()

	req, err := http.NewRequest(http.MethodPost, tokenURL,
		strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if len(client.config.ClientSecret) > 0 {
		req.SetBasicAuth(client.config.ClientID, client.config.ClientSecret)
	}
	resp, err := client.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		err = decodeError("", body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("unexpected response code: %s", resp.Status)
	}

	token := new(AccessToken)
	err = json.Unmarshal(body, token)
	if err != nil {
		return nil, err
	}
	token.notAfter = time.Now().Add(
		time.Duration(token.ExpiresIn) * time.Second)

	return token, nil
}

func (client *Client) login(token string) (code, state string, err error) {
	var body []byte
	body, err = json.Marshal(map[string]string{
		"username": client.config.APIClientID,
		"password": client.config.APIClientSecret,
		"token":    token,
	})
	if err != nil {
		return
	}

	loginURL := client.loginURL()
	req, err := http.NewRequest(http.MethodPost, loginURL,
		bytes.NewReader(body))
	if err != nil {
		return "", "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", UserAgent)

	resp, err := client.httpNoRedirect.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = decodeError("", body)
		if err != nil {
			return "", "", err
		}
		return "", "", fmt.Errorf("unexpected response code: %s", resp.Status)
	}

	response := make(map[string]string)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}
	code, ok := response["code"]
	if ok {
		return code, response["state"], nil
	}
	err = errors.New("login did not return authorization code")

	return
}

// NewToken clears the possible cached token and runs the
// authentifation flow to get a new authentication token.
func (client *Client) NewToken() (string, error) {
	client.m.Lock()
	for client.authPending {
		client.c.Wait()
	}
	client.token = nil
	client.m.Unlock()

	return client.Token()
}

func decodeError(location string, body []byte) error {
	// Error in location's query parameters.
	if len(location) > 0 {
		parsed, err := url.Parse(location)
		if err != nil {
			return fmt.Errorf("invalid location URL: %s", err)
		}
		error := parsed.Query().Get("error")
		if len(error) > 0 {
			desc := parsed.Query().Get("error_description")
			if len(desc) > 0 {
				return fmt.Errorf("%s (%s)", error, desc)
			}
			return errors.New(error)
		}
		// No error attributes in the location query, let's check the
		// response body below.
	}

	// Response body as JSON.
	if len(body) == 0 {
		return nil
	}
	response := make(map[string]string)
	err := json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("invalid error response: %s", err)
	}

	error, ok := response["error"]
	if !ok {
		return fmt.Errorf("invalid error response: no 'error' attribute")
	}
	desc, ok := response["error_description"]
	if ok {
		return fmt.Errorf("%s (%s)", error, desc)
	}

	return errors.New(error)
}
