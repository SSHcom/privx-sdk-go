package secretsmanager

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api/filters"
	"github.com/SSHcom/privx-sdk-go/api/response"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// SecretsManager is a secrets-manager client instance.
type SecretsManager struct {
	api restapi.Connector
}

// New secrets manager client constructor.
func New(api restapi.Connector) *SecretsManager {
	return &SecretsManager{api: api}
}

// MARK: STATUS
// Status get secrets manager microservice status.
func (c *SecretsManager) Status() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/secrets-manager/api/v1/status").
		Get(status)

	return status, err
}

// MARK: Password Policies
// GetPasswordPolicies get password policies.
func (c *SecretsManager) GetPasswordPolicies() (*response.ResultSet[PasswordPolicy], error) {
	policies := &response.ResultSet[PasswordPolicy]{}

	_, err := c.api.
		URL("/secrets-manager/api/v1/password-policies").
		Get(&policies)

	return policies, err
}

// CreatePasswordPolicy create password policy.
func (c *SecretsManager) CreatePasswordPolicy(policy *PasswordPolicy) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/secrets-manager/api/v1/password-policy").
		Post(&policy, &identifier)

	return identifier, err
}

// GetPasswordPolicy get password policy by id.
func (c *SecretsManager) GetPasswordPolicy(policyID string) (*PasswordPolicy, error) {
	policy := &PasswordPolicy{}

	_, err := c.api.
		URL("/secrets-manager/api/v1/password-policy/%s", policyID).
		Get(&policy)

	return policy, err
}

// UpdatePasswordPolicy update password policy.
func (c *SecretsManager) UpdatePasswordPolicy(policyID string, policy *PasswordPolicy) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/password-policy/%s", policyID).
		Put(&policy)

	return err
}

// DeletePasswordPolicy delete password policy.
func (c *SecretsManager) DeletePasswordPolicy(policyID string) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/password-policy/%s", policyID).
		Delete()

	return err
}

// MARK: Manage Passwords
// RotatePassword initiate password rotation.
func (c *SecretsManager) RotatePassword(hostID, account string) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/rotate/%s/%s", hostID, account).
		Post(nil)

	return err
}

// MARK: Manage Rotation Scripts
// GetScriptTemplates get script templates.
func (c *SecretsManager) GetScriptTemplates() (*response.ResultSet[ScriptTemplate], error) {
	templates := &response.ResultSet[ScriptTemplate]{}

	_, err := c.api.
		URL("/secrets-manager/api/v1/script-templates").
		Get(&templates)

	return templates, err
}

// CreateScriptTemplate create script template.
func (c *SecretsManager) CreateScriptTemplate(template *ScriptTemplate) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/secrets-manager/api/v1/script-template").
		Post(&template, &identifier)

	return identifier, err
}

// GetScriptTemplate get script template by id.
func (c *SecretsManager) GetScriptTemplate(templateID string) (*ScriptTemplate, error) {
	p := &ScriptTemplate{}

	_, err := c.api.
		URL("/secrets-manager/api/v1/script-template/%s", templateID).
		Get(&p)

	return p, err
}

// UpdateScriptTemplate update script template.
func (c *SecretsManager) UpdateScriptTemplate(templateID string, template *ScriptTemplate) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/script-template/%s", templateID).
		Put(&template)

	return err
}

// DeleteScriptTemplate delete script template.
func (c *SecretsManager) DeleteScriptTemplate(templateID string) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/password-policy/%s", templateID).
		Delete()

	return err
}

// CompileScript compile script with test data.
func (c *SecretsManager) CompileScript(compile CompileScript) (CompileScriptResponse, error) {
	compiled := CompileScriptResponse{}

	_, err := c.api.
		URL("/secrets-manager/api/v1/script-template/compile").
		Post(&compile, &compiled)

	return compiled, err
}

// MARK: Manage Secrets
// GetHostSecretMetadata get host secret metadata for all accounts.
func (c *SecretsManager) GetHostSecretMetadata(hostID string) (*PostHostSecret, error) {
	secret := &PostHostSecret{}

	_, err := c.api.
		URL("/secrets-manager/api/v1/host-secret/%s", hostID).
		Get(&secret)

	return secret, err
}

// CreateHostSecret create host secret.
func (c *SecretsManager) CreateHostSecret(hostID string, secret *PostHostSecret) (*PostHostSecret, error) {
	hostSecret := &PostHostSecret{}

	_, err := c.api.
		URL("/secrets-manager/api/v1/host-secret/%s", hostID).
		Post(&secret, &hostSecret)

	return hostSecret, err
}

// DeleteHostSecret delete host secret.
func (c *SecretsManager) DeleteHostSecret(hostID string) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/host-secret/%s", hostID).
		Delete()

	return err
}

// MARK: Target Domains
// GetTargetDomains get target domains.
func (c *SecretsManager) GetTargetDomains(opts ...filters.Option) (*response.ResultSet[TargetDomain], error) {
	tds := &response.ResultSet[TargetDomain]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains").
		Query(params).
		Get(&tds)

	return tds, err
}

// CreateTargetDomain create target domain.
func (c *SecretsManager) CreateTargetDomain(td *TargetDomain) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains").
		Post(&td, &identifier)

	return identifier, err
}

// SearchTargetDomain search target domains.
func (c *SecretsManager) SearchTargetDomain(search TargetDomainsSearch, opts ...filters.Option) (*response.ResultSet[TargetDomain], error) {
	tds := &response.ResultSet[TargetDomain]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/search").
		Query(params).
		Post(&search, &tds)

	return tds, err
}

// GetTargetDomain get target domain by id.
func (c *SecretsManager) GetTargetDomain(tdID string) (*TargetDomain, error) {
	td := &TargetDomain{}

	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s", tdID).
		Get(&td)

	return td, err
}

// UpdateTargetDomain update target domain.
func (c *SecretsManager) UpdateTargetDomain(tdID string, td *TargetDomain) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s", tdID).
		Put(&td)

	return err
}

// DeleteTargetDomain delete target domain.
func (c *SecretsManager) DeleteTargetDomain(tdID string) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s", tdID).
		Delete()

	return err
}

// RefreshTargetDomain trigger target domain account scan.
func (c *SecretsManager) RefreshTargetDomain(tdID string) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/refresh", tdID).
		Post(nil)

	return err
}

// MARK: Target domain accounts
// GetTargetDomainAccounts get accounts in target domain.
func (c *SecretsManager) GetTargetDomainAccounts(tdID string, opts ...filters.Option) (*response.ResultSet[ScannedAccount], error) {
	accounts := &response.ResultSet[ScannedAccount]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/accounts", tdID).
		Query(params).
		Get(&accounts)

	return accounts, err
}

// SearchTargetDomainAccounts search accounts in target domain.
func (c *SecretsManager) SearchTargetDomainAccounts(tdID string, search ScannedAccountsSearch, opts ...filters.Option) (*response.ResultSet[ScannedAccount], error) {
	accounts := &response.ResultSet[ScannedAccount]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/accounts/search", tdID).
		Query(params).
		Post(&search, &accounts)

	return accounts, err
}

// GetTargetDomainAccount get target domain account by id.
func (c *SecretsManager) GetTargetDomainAccount(tdID, accountID string) (*ScannedAccount, error) {
	account := &ScannedAccount{}

	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/accounts/%s", tdID, accountID).
		Get(&account)

	return account, err
}

// UpdateTargetDomainAccount update target domain account.
func (c *SecretsManager) UpdateTargetDomainAccount(tdID, accountID string, change ScannedAccountChangeSet) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/accounts/%s", tdID, accountID).
		Put(&change)

	return err
}

// BatchUpdateTargetDomain update target domain in batch.
func (c *SecretsManager) BatchUpdateTargetDomain(tdID string, edit ScannedAccountEditBatch) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/accounts/batch/edit", tdID).
		Post(edit)

	return err
}

// MARK: Managed accounts
// GetManagedAccounts get managed accounts in a target domain.
func (c *SecretsManager) GetManagedAccounts(tdID string, opts ...filters.Option) (*response.ResultSet[ManagedAccount], error) {
	accounts := &response.ResultSet[ManagedAccount]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts", tdID).
		Query(params).
		Get(&accounts)

	return accounts, err
}

// CreateManagedAccount create a managed account.
func (c *SecretsManager) CreateManagedAccount(tdID string, account *ManagedAccount) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts", tdID).
		Post(&account, &identifier)

	return identifier, err
}

// SearchManagedAccounts search managed accounts in a target domain.
func (c *SecretsManager) SearchManagedAccounts(tdID string, search ManagedAccountsSearch, opts ...filters.Option) (*response.ResultSet[ManagedAccount], error) {
	accounts := &response.ResultSet[ManagedAccount]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/search", tdID).
		Query(params).
		Post(&search, &accounts)

	return accounts, err
}

// GetManagedAccount get managed account in target domain by id.
func (c *SecretsManager) GetManagedAccount(tdID, maID string) (*ManagedAccount, error) {
	account := &ManagedAccount{}

	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/%s", tdID, maID).
		Get(&account)

	return account, err
}

// UpdateTargetManagedAccount update managed account.
func (c *SecretsManager) UpdateTargetManagedAccount(tdID, maID string, account *ManagedAccount) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/%s", tdID, maID).
		Put(&account)

	return err
}

// DeleteManagedAccount delete managed account.
func (c *SecretsManager) DeleteManagedAccount(tdID, maID string) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/%s", tdID, maID).
		Delete()

	return err
}

// RotateManagedAccountPassword trigger managed account password rotation.
func (c *SecretsManager) RotateManagedAccountPassword(tdID, maID string) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/%s/rotate", tdID, maID).
		Post(nil)

	return err
}

// SetManagedAccountPassword set password for managed account.
func (c *SecretsManager) SetManagedAccountPassword(tdID, maID, password ManagedAccountPasswordSet) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/%s/password", tdID, maID).
		Post(&password)

	return err
}

// BatchCreateManagedAccount create a batch of managed accounts.
func (c *SecretsManager) BatchCreateManagedAccount(tdID string, create ManagedAccountCreateBatch) (IDList, error) {
	ids := IDList{}

	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/batch/create", tdID).
		Post(&create, &ids)

	return ids, err
}

// BatchUpdateManagedAccount update a batch of managed accounts.
func (c *SecretsManager) BatchUpdateManagedAccount(tdID string, change *ManagedAccountEditBatch) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/batch/edit", tdID).
		Post(&change)

	return err
}

// BatchDeleteManagedAccount delete a batch of managed accounts.
func (c *SecretsManager) BatchDeleteManagedAccount(tdID string, delete ManagedAccountDeleteBatch) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/batch/delete", tdID).
		Post(&delete)

	return err
}

// BatchRotateManagedAccount rotate a batch of managed accounts.
func (c *SecretsManager) BatchRotateManagedAccount(tdID string, rotate ManagedAccountRotateBatch) error {
	_, err := c.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/batch/rotate", tdID).
		Post(&rotate)

	return err
}
