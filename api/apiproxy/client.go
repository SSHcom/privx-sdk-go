package apiproxy

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/v2/api/filters"
	"github.com/SSHcom/privx-sdk-go/v2/api/response"
	"github.com/SSHcom/privx-sdk-go/v2/restapi"
)

// ApiProxy is a api proxy client instance.
type ApiProxy struct {
	api restapi.Connector
}

// New host store client constructor.
func New(api restapi.Connector) *ApiProxy {
	return &ApiProxy{api: api}
}

// MARK: Status
// Status get api proxy microservice status.
func (c *ApiProxy) Status() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/api-proxy/api/v1/status").
		Get(status)

	return status, err
}

// MARK: Config
// GetApiProxyConfig get api proxy config.
func (c *ApiProxy) GetApiProxyConfig() (*ApiProxyAPIConf, error) {
	conf := &ApiProxyAPIConf{}

	_, err := c.api.
		URL("/api-proxy/api/v1/conf").
		Get(&conf)

	return conf, err
}

// MARK: Api Targets
// GetApiTargets get api targets.
func (c *ApiProxy) GetApiTargets(opts ...filters.Option) (*response.ResultSet[ApiTarget], error) {
	apiTargets := &response.ResultSet[ApiTarget]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/api-proxy/api/v1/api-targets").
		Query(params).
		Get(&apiTargets)

	return apiTargets, err
}

// SearchApiTargets search api targets.
func (c *ApiProxy) SearchApiTargets(search *ApiTargetSearchRequest, opts ...filters.Option) (*response.ResultSet[ApiTarget], error) {
	apiTargets := &response.ResultSet[ApiTarget]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/api-proxy/api/v1/api-targets/search").
		Query(params).
		Post(&search, &apiTargets)

	return apiTargets, err
}

// GetApiTargetTags get api target tags.
func (c *ApiProxy) GetApiTargetTags(opts ...filters.Option) (*response.ResultSet[string], error) {
	tags := &response.ResultSet[string]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/api-proxy/api/v1/api-targets/tags").
		Query(params).
		Get(&tags)

	return tags, err
}

// CreateApiTarget create a api target.
func (c *ApiProxy) CreateApiTarget(apiTarget *ApiTarget) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/api-proxy/api/v1/api-targets").
		Post(&apiTarget, &identifier)

	return identifier, err
}

// GetApiTarget get api target by id.
func (c *ApiProxy) GetApiTarget(apiTargetID string) (*ApiTarget, error) {
	apiTarget := &ApiTarget{}

	_, err := c.api.
		URL("/api-proxy/api/v1/api-targets/%s", apiTargetID).
		Get(&apiTarget)

	return apiTarget, err
}

// UpdateApiTarget update api target by id.
func (c *ApiProxy) UpdateApiTarget(apiTargetID string, apiTarget *ApiTarget) error {
	_, err := c.api.
		URL("/api-proxy/api/v1/api-targets/%s", apiTargetID).
		Put(&apiTarget)

	return err
}

// DeleteApiTarget delete api target by id.
func (c *ApiProxy) DeleteApiTarget(apiTargetID string) error {
	_, err := c.api.
		URL("/api-proxy/api/v1/api-targets/%s", apiTargetID).
		Delete()

	return err
}

// MARK: Current Users Client Credentials
// GetCurrentUserClientCredentials get current users client credentials.
func (c *ApiProxy) GetCurrentUserClientCredentials(opts ...filters.Option) (*response.ResultSet[ApiTarget], error) {
	creds := &response.ResultSet[ApiTarget]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/api-proxy/api/v1/users/current/client-credentials").
		Query(params).
		Get(&creds)

	return creds, err
}

// CreateCurrentUserClientCredential create client crendetial for current user.
func (c *ApiProxy) CreateCurrentUserClientCredential(creds *ClientCredential) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/api-proxy/api/v1/users/current/client-credentials").
		Post(&creds, &identifier)

	return identifier, err
}

// GetCurrentUserClientCredential get current users client credential by credential id.
func (c *ApiProxy) GetCurrentUserClientCredential(credID string) (*ClientCredential, error) {
	cred := &ClientCredential{}

	_, err := c.api.
		URL("/api-proxy/api/v1/users/current/client-credentials/%s", credID).
		Get(&cred)

	return cred, err
}

// GetCurrentUserClientCredentialSecret get current users client credential secret by credential id.
func (c *ApiProxy) GetCurrentUserClientCredentialSecret(credID string) ([]byte, error) {
	secret, err := c.api.
		URL("/api-proxy/api/v1/users/current/client-credentials/%s/secret", credID).
		Fetch()

	return secret, err
}

// UpdateCurrentUserClientCredential update current user client credential by credential id.
func (c *ApiProxy) UpdateCurrentUserClientCredential(credID string, cred *ClientCredential) error {
	_, err := c.api.
		URL("/api-proxy/api/v1/users/current/client-credentials/%s", credID).
		Put(&cred)

	return err
}

// DeleteCurrentUserClientCredential delete current user client credential by credential id.
func (c *ApiProxy) DeleteCurrentUserClientCredential(credID string) error {
	_, err := c.api.
		URL("/api-proxy/api/v1/users/current/client-credentials/%s", credID).
		Delete()

	return err
}
