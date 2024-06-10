---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_named_credential"
sidebar_current: "docs-oci-datasource-database_management-named_credential"
description: |-
  Provides details about a specific Named Credential in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_named_credential
This data source provides details about a specific Named Credential resource in Oracle Cloud Infrastructure Database Management service.

Gets the details for the named credential specified by namedCredentialId.


## Example Usage

```hcl
data "oci_database_management_named_credential" "test_named_credential" {
	#Required
	named_credential_id = oci_database_management_named_credential.test_named_credential.id
}
```

## Argument Reference

The following arguments are supported:

* `named_credential_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the named credential.


## Attributes Reference

The following attributes are exported:

* `associated_resource` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource that  is associated to the named credential. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `content` - The details of the named credential.
	* `credential_type` - The type of named credential. Only 'BASIC' is supported currently.
	* `password_secret_access_mode` - The mechanism used to access the password plain text value.
	* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Vault service secret that contains the database user password.
	* `role` - The role of the database user.
	* `user_name` - The user name used to connect to the database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The information specified by the user about the named credential.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the named credential.
* `lifecycle_details` - The details of the lifecycle state.
* `name` - The name of the named credential.
* `scope` - The scope of the named credential.
* `state` - The current lifecycle state of the named credential.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the named credential was created.
* `time_updated` - The date and time the named credential was last updated.
* `type` - The type of resource associated with the named credential.

