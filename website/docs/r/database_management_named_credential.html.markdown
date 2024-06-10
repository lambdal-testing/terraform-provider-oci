---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_named_credential"
sidebar_current: "docs-oci-resource-database_management-named_credential"
description: |-
  Provides the Named Credential resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_named_credential
This resource provides the Named Credential resource in Oracle Cloud Infrastructure Database Management service.

Creates a named credential.


## Example Usage

```hcl
resource "oci_database_management_named_credential" "test_named_credential" {
	#Required
	compartment_id = var.compartment_id
	content {
		#Required
		credential_type = var.named_credential_content_credential_type
		password_secret_access_mode = var.named_credential_content_password_secret_access_mode
		password_secret_id = oci_vault_secret.test_secret.id
		role = var.named_credential_content_role
		user_name = oci_identity_user.test_user.name
	}
	name = var.named_credential_name
	scope = var.named_credential_scope
	type = var.named_credential_type

	#Optional
	associated_resource = var.named_credential_associated_resource
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.named_credential_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `associated_resource` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource that  is associated to the named credential. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the named credential resides. 
* `content` - (Required) (Updatable) The details of the named credential.
	* `credential_type` - (Required) (Updatable) The type of named credential. Only 'BASIC' is supported currently.
	* `password_secret_access_mode` - (Required) (Updatable) The mechanism used to access the password plain text value.
	* `password_secret_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Vault service secret that contains the database user password.
	* `role` - (Required) (Updatable) The role of the database user.
	* `user_name` - (Required) (Updatable) The user name used to connect to the database.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The information specified by the user about the named credential.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `name` - (Required) The name of the named credential. Valid characters are uppercase or lowercase letters, numbers, and "_". The name of the named credential cannot be modified. It must be unique in the compartment and must begin with an alphabetic character. 
* `scope` - (Required) (Updatable) The scope of the named credential.
* `type` - (Required) The type of resource associated with the named credential.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Named Credential
	* `update` - (Defaults to 20 minutes), when updating the Named Credential
	* `delete` - (Defaults to 20 minutes), when destroying the Named Credential


## Import

NamedCredentials can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_named_credential.test_named_credential "id"
```

