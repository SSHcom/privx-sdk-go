package secretsmanager

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/common"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// Client is a secrets-manager client instance.
type Client struct {
	api restapi.Connector
}

// New creates a new secrets-manager client instance
func New(api restapi.Connector) *Client {
	return &Client{api: api}
}

// SecretsManagerStatus get microservice status
func (s *Client) SecretsManagerStatus() (*common.ServiceStatus, error) {
	status := &common.ServiceStatus{}

	_, err := s.api.
		URL("/secrets-manager/api/v1/status").
		Get(status)

	return status, err
}

// MARK: Password Policies
// PasswordPolicies lists all password policies
func (s *Client) PasswordPolicies() ([]PasswordPolicy, error) {
	result := pwPolicyResult{}

	_, err := s.api.
		URL("/secrets-manager/api/v1/password-policies").
		Get(&result)

	return result.Items, err
}

// CreatePasswordPolicy create a password policy
func (s *Client) CreatePasswordPolicy(p PasswordPolicy) (string, error) {
	var object struct {
		ID string `json:"id"`
	}

	_, err := s.api.
		URL("/secrets-manager/api/v1/password-policy").
		Post(&p, &object)

	return object.ID, err
}

// PasswordPolicy get password policy by id
func (s *Client) PasswordPolicy(policyId string) (*PasswordPolicy, error) {
	p := &PasswordPolicy{}

	_, err := s.api.
		URL("/secrets-manager/api/v1/password-policy/%s", url.PathEscape(policyId)).
		Get(&p)

	return p, err
}

// UpdatePasswordPolicy update existing password policy
func (s *Client) UpdatePasswordPolicy(policyId string, p PasswordPolicy) error {
	_, err := s.api.
		URL("/secrets-manager/api/v1/password-policy/%s", url.PathEscape(policyId)).
		Put(p)

	return err
}

// DeletePasswordPolicy delete a password policy
func (s *Client) DeletePasswordPolicy(policyId string) error {
	_, err := s.api.
		URL("/secrets-manager/api/v1/password-policy/%s", url.PathEscape(policyId)).
		Delete()

	return err
}

// MARK: Manage passwords
// RotatePassword initiate password rotation
func (s *Client) RotatePassword(hostId, account string) error {
	_, err := s.api.
		URL("/secrets-manager/api/v1/rotate/%s/%s", url.PathEscape(hostId), url.PathEscape(account)).
		Post(nil)

	return err
}

// MARK: Manage rotation scripts
// ScriptTemplates lists all script templates
func (s *Client) ScriptTemplates() ([]ScriptTemplate, error) {
	result := scriptTemplateResult{}

	_, err := s.api.
		URL("/secrets-manager/api/v1/script-templates").
		Get(&result)

	return result.Items, err
}

// CreateScriptTemplate create a script template
func (s *Client) CreateScriptTemplate(t ScriptTemplate) (string, error) {
	var object struct {
		ID string `json:"id"`
	}

	_, err := s.api.
		URL("/secrets-manager/api/v1/script-template").
		Post(&t, &object)

	return object.ID, err
}

// ScriptTemplate get script template by id
func (s *Client) ScriptTemplate(templateId string) (*ScriptTemplate, error) {
	p := &ScriptTemplate{}

	_, err := s.api.
		URL("/secrets-manager/api/v1/script-template/%s", url.PathEscape(templateId)).
		Get(&p)

	return p, err
}

// UpdateScriptTemplate update existing script template
func (s *Client) UpdateScriptTemplate(templateId string, t ScriptTemplate) error {
	_, err := s.api.
		URL("/secrets-manager/api/v1/script-template/%s", url.PathEscape(templateId)).
		Put(t)

	return err
}

// DeleteScriptTemplate delete a script template
func (s *Client) DeleteScriptTemplate(templateId string) error {
	_, err := s.api.
		URL("/secrets-manager/api/v1/password-policy/%s", url.PathEscape(templateId)).
		Delete()

	return err
}

// CompileScript compile script with test data
func (s *Client) CompileScript(r CompileScriptRequest) (string, error) {
	var object struct {
		Script string `json:"script"`
	}

	_, err := s.api.
		URL("/secrets-manager/api/v1/script-template/compile").
		Post(&r, &object)

	return object.Script, err
}

// MARK: Target domains
// TargetDomains lists all target domains
func (s *Client) TargetDomains(offset, limit int, sortkey, sortdir string) ([]TargetDomain, error) {
	result := tdResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}

	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains").
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// CreateTargetDomain create a target domain
func (s *Client) CreateTargetDomain(td TargetDomain) (string, error) {
	var object struct {
		ID string `json:"id"`
	}

	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains").
		Post(&td, &object)

	return object.ID, err
}

// SearchTargetDomain search for existing target domain
func (s *Client) SearchTargetDomain(sortkey, sortdir string, offset, limit int, searchObject TargetDomainsSearch) ([]TargetDomain, error) {
	result := tdResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}

	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/search").
		Query(&filters).
		Post(&searchObject, &result)

	return result.Items, err
}

// TargetDomain get target domain by id
func (s *Client) TargetDomain(tdId string) (*TargetDomain, error) {
	td := &TargetDomain{}

	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s", url.PathEscape(tdId)).
		Get(&td)

	return td, err
}

// UpdateTargetDomain update existing target domain
func (s *Client) UpdateTargetDomain(tdId string, td TargetDomain) error {
	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s", url.PathEscape(tdId)).
		Put(td)

	return err
}

// DeleteTargetDomain delete a target domain
func (s *Client) DeleteTargetDomain(tdId string) error {
	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s", url.PathEscape(tdId)).
		Delete()

	return err
}

// RefreshTargetDomain trigger target domain account scan
func (s *Client) RefreshTargetDomain(tdId string) error {
	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/refresh", url.PathEscape(tdId)).
		Post(nil)

	return err
}

// MARK: Target domain accounts
// TargetDomainAccounts lists all accounts in target domain
func (s *Client) TargetDomainAccounts(offset, limit int, sortkey, sortdir, tdId string) ([]ScannedAccount, error) {
	result := scannedAccountResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}

	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/accounts", url.PathEscape(tdId)).
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// SearchTargetDomainAccounts search accounts in target domain
func (s *Client) SearchTargetDomainAccounts(sortkey, sortdir, tdId string, offset, limit int, searchObject ScannedAccountsSearch) ([]ScannedAccount, error) {
	result := scannedAccountResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}

	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/accounts/search", url.PathEscape(tdId)).
		Query(&filters).
		Post(&searchObject, &result)

	return result.Items, err
}

// TargetDomainAccount get target domain account
func (s *Client) TargetDomainAccount(tdId, accountId string) (ScannedAccount, error) {
	account := ScannedAccount{}

	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/accounts/%s", url.PathEscape(tdId), url.PathEscape(accountId)).
		Get(&account)

	return account, err
}

// UpdateTargetDomainAccount update target domain account
func (s *Client) UpdateTargetDomainAccount(tdId, accountId string, change ScannedAccountChangeSet) error {
	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/accounts/%s", url.PathEscape(tdId), url.PathEscape(accountId)).
		Put(change)

	return err
}

// BatchUpdateTargetDomain update target domain in batch
func (s *Client) BatchUpdateTargetDomain(tdId string, change ScannedAccountEditBatch) error {
	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/accounts/batch/edit", url.PathEscape(tdId)).
		Post(change)

	return err
}

// MARK: Managed accounts
// ManagedAccounts lists all managed accounts in a target domain
func (s *Client) ManagedAccounts(offset, limit int, sortkey, sortdir, tdId string) ([]ManagedAccount, error) {
	result := managedAccountResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}

	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts", url.PathEscape(tdId)).
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// CreateManagedAccount create a managed account
func (s *Client) CreateManagedAccount(tdId string, ma ManagedAccount) (string, error) {
	var object struct {
		ID string `json:"id"`
	}

	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts", url.PathEscape(tdId)).
		Post(&ma, &object)

	return object.ID, err
}

// SearchManagedAccounts search managed accounts in a target domain
func (s *Client) SearchManagedAccounts(sortkey, sortdir, tdId string, offset, limit int, searchObject ManagedAccountsSearch) ([]ManagedAccount, error) {
	result := managedAccountResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}

	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/search", url.PathEscape(tdId)).
		Query(&filters).
		Post(&searchObject, &result)

	return result.Items, err
}

// ManagedAccount get managed account
func (s *Client) ManagedAccount(tdId, maId string) (ManagedAccount, error) {
	account := ManagedAccount{}

	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/%s", url.PathEscape(tdId), url.PathEscape(maId)).
		Get(&account)

	return account, err
}

// UpdateTargetManagedAccount update managed account
func (s *Client) UpdateTargetManagedAccount(tdId, maId string, change ManagedAccount) error {
	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/%s", url.PathEscape(tdId), url.PathEscape(maId)).
		Put(change)

	return err
}

// DeleteManagedAccount delete managed account
func (s *Client) DeleteManagedAccount(tdId, maId string) error {
	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/%s", url.PathEscape(tdId), url.PathEscape(maId)).
		Delete()

	return err
}

// RotateManagedAccountPassword trigger managed account password rotation
func (s *Client) RotateManagedAccountPassword(tdId, maId string) error {
	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/%s/rotate", url.PathEscape(tdId), url.PathEscape(maId)).
		Post(nil)

	return err
}

// ManagedAccountPassword provide password for managed account
func (s *Client) ManagedAccountPassword(tdId, maId, password string) error {
	pwReq := ManagedAccountPasswordRequest{
		Password: password,
	}

	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/%s/password", url.PathEscape(tdId), url.PathEscape(maId)).
		Post(pwReq)

	return err
}

// BatchCreateManagedAccount create a batch of managed accounts
func (s *Client) BatchCreateManagedAccount(tdId string, ma ManagedAccountCreateBatch) ([]string, error) {
	var object struct {
		IDs []string `json:"ids"`
	}

	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/batch/create", url.PathEscape(tdId)).
		Post(&ma, &object)

	return object.IDs, err
}

// BatchUpdateManagedAccount update a batch of managed accounts
func (s *Client) BatchUpdateManagedAccount(tdId string, change ManagedAccountChangeSet) error {
	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/batch/edit", url.PathEscape(tdId)).
		Post(&change)

	return err
}

// BatchDeleteManagedAccount delete a batch of managed accounts
func (s *Client) BatchDeleteManagedAccount(tdId string, delete ManagedAccountBatch) error {
	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/batch/delete", url.PathEscape(tdId)).
		Post(&delete)

	return err
}

// BatchRotateManagedAccount rotate a batch of managed accounts
func (s *Client) BatchRotateManagedAccount(tdId string, rotate ManagedAccountBatch) error {
	_, err := s.api.
		URL("/secrets-manager/api/v1/targetdomains/%s/managedaccounts/batch/rotate", url.PathEscape(tdId)).
		Post(&rotate)

	return err
}
