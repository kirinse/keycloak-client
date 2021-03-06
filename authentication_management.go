package keycloak

import (
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	authenticationManagementPath = "/auth/admin/realms/:realm/authentication"
)

// GetAuthenticatorProviders returns a list of authenticator providers.
func (c *Client) GetAuthenticatorProviders(accessToken string, realmName string) ([]map[string]interface{}, error) {
	var resp = []map[string]interface{}{}
	var err = c.get(accessToken, &resp, url.Path(authenticationManagementPath+"/authenticator-providers"), url.Param("realm", realmName))
	return resp, err
}

// GetClientAuthenticatorProviders returns a list of client authenticator providers.
func (c *Client) GetClientAuthenticatorProviders(accessToken string, realmName string) ([]map[string]interface{}, error) {
	var resp = []map[string]interface{}{}
	var err = c.get(accessToken, &resp, url.Path(authenticationManagementPath+"/client-authenticator-providers"), url.Param("realm", realmName))
	return resp, err
}

// GetAuthenticatorProviderConfig returns the authenticator provider’s configuration description.
func (c *Client) GetAuthenticatorProviderConfig(accessToken string, realmName, providerID string) (AuthenticatorConfigInfoRepresentation, error) {
	var resp = AuthenticatorConfigInfoRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(authenticationManagementPath+"/config-description/:providerID"), url.Param("realm", realmName), url.Param("providerID", providerID))
	return resp, err
}

// GetAuthenticatorConfig returns the authenticator configuration.
func (c *Client) GetAuthenticatorConfig(accessToken string, realmName, configID string) (AuthenticatorConfigRepresentation, error) {
	var resp = AuthenticatorConfigRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(authenticationManagementPath+"/config/:id"), url.Param("realm", realmName), url.Param("id", configID))
	return resp, err
}

// UpdateAuthenticatorConfig updates the authenticator configuration.
func (c *Client) UpdateAuthenticatorConfig(accessToken string, realmName, configID string, config AuthenticatorConfigRepresentation) error {
	return c.put(accessToken, url.Path(authenticationManagementPath+"/config/:id"), url.Param("realm", realmName), url.Param("id", configID), body.JSON(config))
}

// DeleteAuthenticatorConfig deletes the authenticator configuration.
func (c *Client) DeleteAuthenticatorConfig(accessToken string, realmName, configID string) error {
	return c.delete(accessToken, url.Path(authenticationManagementPath+"/config/:id"), url.Param("realm", realmName), url.Param("id", configID))
}

// CreateAuthenticationExecution add new authentication execution
func (c *Client) CreateAuthenticationExecution(accessToken string, realmName string, authExec AuthenticationExecutionRepresentation) (string, error) {
	return c.post(accessToken, nil, url.Path(authenticationManagementPath+"/executions"), url.Param("realm", realmName), body.JSON(authExec))
}

// DeleteAuthenticationExecution deletes the execution.
func (c *Client) DeleteAuthenticationExecution(accessToken string, realmName, executionID string) error {
	return c.delete(accessToken, url.Path(authenticationManagementPath+"/executions/:id"), url.Param("realm", realmName), url.Param("id", executionID))
}

// UpdateAuthenticationExecution update execution with new configuration.
func (c *Client) UpdateAuthenticationExecution(accessToken string, realmName, executionID string, authConfig AuthenticatorConfigRepresentation) error {
	_, err := c.post(accessToken, nil, url.Path(authenticationManagementPath+"/executions/:id/config"), url.Param("realm", realmName), url.Param("id", executionID), body.JSON(authConfig))
	return err
}

// LowerExecutionPriority lowers the execution’s priority.
func (c *Client) LowerExecutionPriority(accessToken string, realmName, executionID string) error {
	_, err := c.post(accessToken, nil, url.Path(authenticationManagementPath+"/executions/:id/lower-priority"), url.Param("realm", realmName), url.Param("id", executionID))
	return err
}

// RaiseExecutionPriority raise the execution’s priority.
func (c *Client) RaiseExecutionPriority(accessToken string, realmName, executionID string) error {
	_, err := c.post(accessToken, nil, url.Path(authenticationManagementPath+"/executions/:id/raise-priority"), url.Param("realm", realmName), url.Param("id", executionID))
	return err
}

// CreateAuthenticationFlow creates a new authentication flow.
func (c *Client) CreateAuthenticationFlow(accessToken string, realmName string, authFlow AuthenticationFlowRepresentation) error {
	_, err := c.post(accessToken, nil, url.Path(authenticationManagementPath+"/flows"), url.Param("realm", realmName), body.JSON(authFlow))
	return err
}

// GetAuthenticationFlows returns a list of authentication flows.
func (c *Client) GetAuthenticationFlows(accessToken string, realmName string) ([]AuthenticationFlowRepresentation, error) {
	var resp = []AuthenticationFlowRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(authenticationManagementPath+"/flows"), url.Param("realm", realmName))
	return resp, err
}

// CopyExistingAuthenticationFlow copy the existing authentication flow under a new name.
// 'flowAlias' is the name of the existing authentication flow,
// 'newName' is the new name of the authentication flow.
func (c *Client) CopyExistingAuthenticationFlow(accessToken string, realmName, flowAlias, newName string) error {
	var m = map[string]string{"newName": newName}
	_, err := c.post(accessToken, nil, url.Path(authenticationManagementPath+"/flows/:flowAlias/copy"), url.Param("realm", realmName), url.Param("flowAlias", flowAlias), body.JSON(m))
	return err
}

// GetAuthenticationExecutionForFlow returns the authentication executions for a flow.
func (c *Client) GetAuthenticationExecutionForFlow(accessToken string, realmName, flowAlias string) (AuthenticationExecutionInfoRepresentation, error) {
	var resp = AuthenticationExecutionInfoRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(authenticationManagementPath+"/flows/:flowAlias/executions"), url.Param("realm", realmName), url.Param("flowAlias", flowAlias))
	return resp, err
}

// UpdateAuthenticationExecutionForFlow updates the authentication executions of a flow.
func (c *Client) UpdateAuthenticationExecutionForFlow(accessToken string, realmName, flowAlias string, authExecInfo AuthenticationExecutionInfoRepresentation) error {
	return c.put(accessToken, url.Path(authenticationManagementPath+"/flows/:flowAlias/executions"), url.Param("realm", realmName), url.Param("flowAlias", flowAlias), body.JSON(authExecInfo))
}

// CreateAuthenticationExecutionForFlow add a new authentication execution to a flow.
// 'flowAlias' is the alias of the parent flow.
func (c *Client) CreateAuthenticationExecutionForFlow(accessToken string, realmName, flowAlias, provider string) (string, error) {
	var m = map[string]string{"provider": provider}
	return c.post(accessToken, nil, url.Path(authenticationManagementPath+"/flows/:flowAlias/executions/execution"), url.Param("realm", realmName), url.Param("flowAlias", flowAlias), body.JSON(m))
}

// CreateFlowWithExecutionForExistingFlow add a new flow with a new execution to an existing flow.
// 'flowAlias' is the alias of the parent authentication flow.
func (c *Client) CreateFlowWithExecutionForExistingFlow(accessToken string, realmName, flowAlias, alias, flowType, provider, description string) (string, error) {
	var m = map[string]string{"alias": alias, "type": flowType, "provider": provider, "description": description}
	return c.post(accessToken, nil, url.Path(authenticationManagementPath+"/flows/:flowAlias/executions/flow"), url.Param("realm", realmName), url.Param("flowAlias", flowAlias), body.JSON(m))
}

// GetAuthenticationFlow gets the authentication flow for id.
func (c *Client) GetAuthenticationFlow(accessToken string, realmName, flowID string) (AuthenticationFlowRepresentation, error) {
	var resp = AuthenticationFlowRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(authenticationManagementPath+"/flows/:id"), url.Param("realm", realmName), url.Param("id", flowID))
	return resp, err
}

// DeleteAuthenticationFlow deletes an authentication flow.
func (c *Client) DeleteAuthenticationFlow(accessToken string, realmName, flowID string) error {
	return c.delete(accessToken, url.Path(authenticationManagementPath+"/flows/:id"), url.Param("realm", realmName), url.Param("id", flowID))
}

// GetFormActionProviders returns a list of form action providers.
func (c *Client) GetFormActionProviders(accessToken string, realmName string) ([]map[string]interface{}, error) {
	var resp = []map[string]interface{}{}
	var err = c.get(accessToken, &resp, url.Path(authenticationManagementPath+"/form-action-providers"), url.Param("realm", realmName))
	return resp, err
}

// GetFormProviders returns a list of form providers.
func (c *Client) GetFormProviders(accessToken string, realmName string) ([]map[string]interface{}, error) {
	var resp = []map[string]interface{}{}
	var err = c.get(accessToken, &resp, url.Path(authenticationManagementPath+"/form-providers"), url.Param("realm", realmName))
	return resp, err
}

// GetConfigDescriptionForClients returns the configuration descriptions for all clients.
func (c *Client) GetConfigDescriptionForClients(accessToken string, realmName string) (map[string]interface{}, error) {
	var resp = map[string]interface{}{}
	var err = c.get(accessToken, &resp, url.Path(authenticationManagementPath+"/per-client-config-description"), url.Param("realm", realmName))
	return resp, err
}

// RegisterRequiredAction register a new required action.
func (c *Client) RegisterRequiredAction(accessToken string, realmName, providerID, name string) error {
	var m = map[string]string{"providerId": providerID, "name": name}
	_, err := c.post(accessToken, nil, url.Path(authenticationManagementPath+"/register-required-action"), url.Param("realm", realmName), body.JSON(m))
	return err
}

// GetRequiredActions returns a list of required actions.
func (c *Client) GetRequiredActions(accessToken string, realmName string) ([]RequiredActionProviderRepresentation, error) {
	var resp = []RequiredActionProviderRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(authenticationManagementPath+"/required-actions"), url.Param("realm", realmName))
	return resp, err
}

// GetRequiredAction returns the required action for the alias.
func (c *Client) GetRequiredAction(accessToken string, realmName, actionAlias string) (RequiredActionProviderRepresentation, error) {
	var resp = RequiredActionProviderRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(authenticationManagementPath+"/required-actions/:alias"), url.Param("realm", realmName), url.Param("alias", actionAlias))
	return resp, err
}

// UpdateRequiredAction updates the required action.
func (c *Client) UpdateRequiredAction(accessToken string, realmName, actionAlias string, action RequiredActionProviderRepresentation) error {
	return c.put(accessToken, url.Path(authenticationManagementPath+"/required-actions/:alias"), url.Param("realm", realmName), url.Param("alias", actionAlias), body.JSON(action))
}

// DeleteRequiredAction deletes the required action.
func (c *Client) DeleteRequiredAction(accessToken string, realmName, actionAlias string) error {
	return c.delete(accessToken, url.Path(authenticationManagementPath+"/required-actions/:alias"), url.Param("realm", realmName), url.Param("alias", actionAlias))
}

// GetUnregisteredRequiredActions returns a list of unregistered required actions.
func (c *Client) GetUnregisteredRequiredActions(accessToken string, realmName string) ([]map[string]interface{}, error) {
	var resp = []map[string]interface{}{}
	var err = c.get(accessToken, &resp, url.Path(authenticationManagementPath+"/unregistered-required-actions"), url.Param("realm", realmName))
	return resp, err
}
