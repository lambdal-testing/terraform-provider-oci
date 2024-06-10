---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_pluggabledatabase_pluggable_database_dbm_features_management"
sidebar_current: "docs-oci-resource-database_management-pluggabledatabase_pluggable_database_dbm_features_management"
description: |-
  Provides the Pluggabledatabase Pluggable Database Dbm Features Management resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_pluggabledatabase_pluggable_database_dbm_features_management
This resource provides the Pluggabledatabase Pluggable Database Dbm Features Management resource in Oracle Cloud Infrastructure Database Management service.

Enables a Database Management feature for the specified Oracle cloud pluggable database.


## Example Usage

```hcl
resource "oci_database_management_pluggabledatabase_pluggable_database_dbm_features_management" "test_pluggabledatabase_pluggable_database_dbm_features_management" {
	#Required
	pluggable_database_id = oci_database_pluggable_database.test_pluggable_database.id
	enable_pluggable_database_dbm_feature = var.enable_pluggable_database_dbm_feature

	#Optional
	feature_details {
		#Required
		feature = var.pluggabledatabase_pluggable_database_dbm_features_management_feature_details_feature
		enable_pluggable_database_dbm_feature = var.enable_pluggable_database_dbm_feature

		#Optional
		connector_details {

			#Optional
			connector_type = var.pluggabledatabase_pluggable_database_dbm_features_management_feature_details_connector_details_connector_type
			database_connector_id = oci_database_management_database_connector.test_database_connector.id
			management_agent_id = oci_management_agent_management_agent.test_management_agent.id
			private_end_point_id = oci_database_management_private_end_point.test_private_end_point.id
		}
		database_connection_details {

			#Optional
			connection_credentials {

				#Optional
				credential_name = var.pluggabledatabase_pluggable_database_dbm_features_management_feature_details_database_connection_details_connection_credentials_credential_name
				credential_type = var.pluggabledatabase_pluggable_database_dbm_features_management_feature_details_database_connection_details_connection_credentials_credential_type
				password_secret_id = oci_vault_secret.test_secret.id
				role = var.pluggabledatabase_pluggable_database_dbm_features_management_feature_details_database_connection_details_connection_credentials_role
				ssl_secret_id = oci_vault_secret.test_secret.id
				user_name = oci_identity_user.test_user.name
			}
			connection_string {

				#Optional
				connection_type = var.pluggabledatabase_pluggable_database_dbm_features_management_feature_details_database_connection_details_connection_string_connection_type
				port = var.pluggabledatabase_pluggable_database_dbm_features_management_feature_details_database_connection_details_connection_string_port
				protocol = var.pluggabledatabase_pluggable_database_dbm_features_management_feature_details_database_connection_details_connection_string_protocol
				service = var.pluggabledatabase_pluggable_database_dbm_features_management_feature_details_database_connection_details_connection_string_service
			}
		}
		management_type = var.pluggabledatabase_pluggable_database_dbm_features_management_feature_details_management_type
		is_auto_enable_pluggable_database = var.pluggabledatabase_pluggable_database_dbm_features_management_feature_details_is_auto_enable_pluggable_database
	}
}
```

## Argument Reference

The following arguments are supported:

* `feature_details` - (Optional) The details required to enable the specified Database Management feature.
	* `connector_details` - (Optional) The connector details required to connect to an Oracle cloud database.
		* `connector_type` - (Optional) The list of supported connection types:
			* PE: Private endpoint
			* MACS: Management agent
			* EXTERNAL: External database connector 
		* `database_connector_id` - (Applicable when connector_type=EXTERNAL) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database connector.
		* `management_agent_id` - (Applicable when connector_type=MACS) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent.
		* `private_end_point_id` - (Applicable when connector_type=PE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint.
	* `database_connection_details` - (Optional) The connection details required to connect to the database.
		* `connection_credentials` - (Optional) The credentials used to connect to the database. Currently only the `DETAILS` type is supported for creating MACS connector credentials. 
			* `credential_name` - (Optional) The name of the credential information that used to connect to the DB system resource. The name should be in "x.y" format, where the length of "x" has a maximum of 64 characters, and length of "y" has a maximum of 199 characters. The name strings can contain letters, numbers and the underscore character only. Other characters are not valid, except for the "." character that separates the "x" and "y" portions of the name. *IMPORTANT* - The name must be unique within the Oracle Cloud Infrastructure region the credential is being created in. If you specify a name that duplicates the name of another credential within the same Oracle Cloud Infrastructure region, you may overwrite or corrupt the credential that is already using the name.

				For example: inventorydb.abc112233445566778899 
			* `credential_type` - (Optional) The type of credential used to connect to the database.
			* `password_secret_id` - (Applicable when credential_type=DETAILS | SSL_DETAILS) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
			* `role` - (Applicable when credential_type=DETAILS | SSL_DETAILS) The role of the user connecting to the database.
			* `ssl_secret_id` - (Applicable when credential_type=SSL_DETAILS) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the SSL keystore and truststore details.
			* `user_name` - (Applicable when credential_type=DETAILS | SSL_DETAILS) The user name used to connect to the database.
		* `connection_string` - (Optional) The details of the Oracle Database connection string. 
			* `connection_type` - (Optional) The list of supported connection types:
				* BASIC: Basic connection details 
			* `port` - (Optional) The port number used to connect to the database.
			* `protocol` - (Optional) The protocol used to connect to the database.
			* `service` - (Optional) The service name of the database.
	* `feature` - (Required) The name of the Database Management feature.
	* `is_auto_enable_pluggable_database` - (Optional) Indicates whether the pluggable database can be enabled automatically.
	* `management_type` - (Optional) The management type for the database.
* `pluggable_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle cloud pluggable database.
* `enable_pluggable_database_dbm_feature` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Pluggabledatabase Pluggable Database Dbm Features Management
	* `update` - (Defaults to 20 minutes), when updating the Pluggabledatabase Pluggable Database Dbm Features Management
	* `delete` - (Defaults to 20 minutes), when destroying the Pluggabledatabase Pluggable Database Dbm Features Management
