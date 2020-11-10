package oauth

import (
	"strings"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

/*

With implements fallback auth strategy, it choose the desired
strategy depends on provided configuration options.

1. OAuth2 Resource Owner Password Grant is default and recommended strategy.
This strategy requires definition of
* Client Access Key, see oauth.Access(...)
* Client Secret Key, see oauth.Secret(...)
* Client Secret Digest, also known as OAuth Client Secrets, see oauth.Digest(...)

2. OAuth2 Authorization Code Grant strategy is used when only single pair
of access/secret key is provided. It allows usage of username and password
* Client Access Key, see oauth.Access(...)
* Client Secret Key, see oauth.Secret(...)

3. Finally, it falls back explicit token token definition

*/
func With(client restapi.Connector, opts ...Option) restapi.Authorizer {
	auth := newAuth(client, opts...)

	if strings.HasPrefix(auth.secret, "Bearer") {
		return &tAuthExplicit{auth.secret}
	}

	if auth.access != "" && auth.secret != "" && auth.digest != "" {
		return &tAuthPassword{tAuth: auth}
	}

	if auth.access != "" && auth.secret != "" {
		return &tAuthCode{tAuth: auth}
	}

	return &tAuthExplicit{auth.secret}
}
