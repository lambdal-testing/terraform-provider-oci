---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_user_consumer_group_privilege"
sidebar_current: "docs-oci-datasource-database_management-managed_database_user_consumer_group_privilege"
description: |-
  Provides details about a specific Managed Database User Consumer Group Privilege in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_user_consumer_group_privilege
This data source provides details about a specific Managed Database User Consumer Group Privilege resource in Oracle Cloud Infrastructure Database Management service.

Gets the list of consumer group privileges granted to a specific user.

## Example Usage

```hcl
data "oci_database_management_managed_database_user_consumer_group_privilege" "test_managed_database_user_consumer_group_privilege" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	user_name = oci_identity_user.test_user.name

	#Optional
	name = var.managed_database_user_consumer_group_privilege_name
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `name` - (Optional) A filter to return only resources that match the entire name.
* `user_name` - (Required) The name of the user whose details are to be viewed.


## Attributes Reference

The following attributes are exported:

* `items` - An array of consumer group privileges.
	* `grant_option` - Indicates whether the privilege is granted with the GRANT option (YES) or not (NO).
	* `initial_group` - Indicates whether the consumer group is designated as the default for this user or role (YES) or not (NO).
	* `name` - The name of the granted consumer group privilege.

