// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_identity_domains_account_mgmt_info", IdentityDomainsAccountMgmtInfoDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_account_mgmt_infos", IdentityDomainsAccountMgmtInfosDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_account_recovery_setting", IdentityDomainsAccountRecoverySettingDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_account_recovery_settings", IdentityDomainsAccountRecoverySettingsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_api_key", IdentityDomainsApiKeyDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_api_keys", IdentityDomainsApiKeysDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_app", IdentityDomainsAppDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_app_role", IdentityDomainsAppRoleDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_app_roles", IdentityDomainsAppRolesDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_approval_workflow", IdentityDomainsApprovalWorkflowDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_approval_workflow_assignment", IdentityDomainsApprovalWorkflowAssignmentDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_approval_workflow_assignments", IdentityDomainsApprovalWorkflowAssignmentsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_approval_workflow_step", IdentityDomainsApprovalWorkflowStepDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_approval_workflow_steps", IdentityDomainsApprovalWorkflowStepsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_approval_workflows", IdentityDomainsApprovalWorkflowsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_apps", IdentityDomainsAppsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_auth_token", IdentityDomainsAuthTokenDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_auth_tokens", IdentityDomainsAuthTokensDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_authentication_factor_setting", IdentityDomainsAuthenticationFactorSettingDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_authentication_factor_settings", IdentityDomainsAuthenticationFactorSettingsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_branding_setting", IdentityDomainsBrandingSettingDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_branding_settings", IdentityDomainsBrandingSettingsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_cloud_gate", IdentityDomainsCloudGateDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_cloud_gate_mapping", IdentityDomainsCloudGateMappingDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_cloud_gate_mappings", IdentityDomainsCloudGateMappingsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_cloud_gate_server", IdentityDomainsCloudGateServerDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_cloud_gate_servers", IdentityDomainsCloudGateServersDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_cloud_gates", IdentityDomainsCloudGatesDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_condition", IdentityDomainsConditionDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_conditions", IdentityDomainsConditionsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_customer_secret_key", IdentityDomainsCustomerSecretKeyDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_customer_secret_keys", IdentityDomainsCustomerSecretKeysDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_dynamic_resource_group", IdentityDomainsDynamicResourceGroupDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_dynamic_resource_groups", IdentityDomainsDynamicResourceGroupsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_grant", IdentityDomainsGrantDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_grants", IdentityDomainsGrantsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_group", IdentityDomainsGroupDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_groups", IdentityDomainsGroupsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_identity_propagation_trust", IdentityDomainsIdentityPropagationTrustDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_identity_propagation_trusts", IdentityDomainsIdentityPropagationTrustsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_identity_provider", IdentityDomainsIdentityProviderDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_identity_providers", IdentityDomainsIdentityProvidersDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_identity_setting", IdentityDomainsIdentitySettingDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_identity_settings", IdentityDomainsIdentitySettingsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_kmsi_setting", IdentityDomainsKmsiSettingDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_kmsi_settings", IdentityDomainsKmsiSettingsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_api_key", IdentityDomainsMyApiKeyDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_api_keys", IdentityDomainsMyApiKeysDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_apps", IdentityDomainsMyAppsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_auth_token", IdentityDomainsMyAuthTokenDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_auth_tokens", IdentityDomainsMyAuthTokensDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_completed_approval", IdentityDomainsMyCompletedApprovalDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_completed_approvals", IdentityDomainsMyCompletedApprovalsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_customer_secret_key", IdentityDomainsMyCustomerSecretKeyDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_customer_secret_keys", IdentityDomainsMyCustomerSecretKeysDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_device", IdentityDomainsMyDeviceDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_devices", IdentityDomainsMyDevicesDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_groups", IdentityDomainsMyGroupsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_oauth2client_credential", IdentityDomainsMyOAuth2ClientCredentialDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_oauth2client_credentials", IdentityDomainsMyOAuth2ClientCredentialsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_pending_approval", IdentityDomainsMyPendingApprovalDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_pending_approvals", IdentityDomainsMyPendingApprovalsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_requestable_groups", IdentityDomainsMyRequestableGroupsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_requests", IdentityDomainsMyRequestsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_smtp_credential", IdentityDomainsMySmtpCredentialDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_smtp_credentials", IdentityDomainsMySmtpCredentialsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_support_account", IdentityDomainsMySupportAccountDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_support_accounts", IdentityDomainsMySupportAccountsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_trusted_user_agent", IdentityDomainsMyTrustedUserAgentDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_trusted_user_agents", IdentityDomainsMyTrustedUserAgentsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_user_db_credential", IdentityDomainsMyUserDbCredentialDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_my_user_db_credentials", IdentityDomainsMyUserDbCredentialsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_network_perimeter", IdentityDomainsNetworkPerimeterDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_network_perimeters", IdentityDomainsNetworkPerimetersDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_notification_setting", IdentityDomainsNotificationSettingDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_notification_settings", IdentityDomainsNotificationSettingsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_oauth2client_credential", IdentityDomainsOAuth2ClientCredentialDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_oauth2client_credentials", IdentityDomainsOAuth2ClientCredentialsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_oauth_client_certificate", IdentityDomainsOAuthClientCertificateDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_oauth_client_certificates", IdentityDomainsOAuthClientCertificatesDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_oauth_partner_certificate", IdentityDomainsOAuthPartnerCertificateDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_oauth_partner_certificates", IdentityDomainsOAuthPartnerCertificatesDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_password_policies", IdentityDomainsPasswordPoliciesDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_password_policy", IdentityDomainsPasswordPolicyDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_policies", IdentityDomainsPoliciesDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_policy", IdentityDomainsPolicyDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_resource_type_schema_attributes", IdentityDomainsResourceTypeSchemaAttributesDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_rule", IdentityDomainsRuleDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_rules", IdentityDomainsRulesDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_security_question", IdentityDomainsSecurityQuestionDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_security_question_setting", IdentityDomainsSecurityQuestionSettingDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_security_question_settings", IdentityDomainsSecurityQuestionSettingsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_security_questions", IdentityDomainsSecurityQuestionsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_self_registration_profile", IdentityDomainsSelfRegistrationProfileDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_self_registration_profiles", IdentityDomainsSelfRegistrationProfilesDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_setting", IdentityDomainsSettingDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_settings", IdentityDomainsSettingsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_smtp_credential", IdentityDomainsSmtpCredentialDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_smtp_credentials", IdentityDomainsSmtpCredentialsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_user", IdentityDomainsUserDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_user_attributes_setting", IdentityDomainsUserAttributesSettingDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_user_attributes_settings", IdentityDomainsUserAttributesSettingsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_user_db_credential", IdentityDomainsUserDbCredentialDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_user_db_credentials", IdentityDomainsUserDbCredentialsDataSource())
	tfresource.RegisterDatasource("oci_identity_domains_users", IdentityDomainsUsersDataSource())
}