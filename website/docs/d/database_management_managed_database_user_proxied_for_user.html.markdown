---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_user_proxied_for_user"
sidebar_current: "docs-oci-datasource-database_management-managed_database_user_proxied_for_user"
description: |-
  Provides details about a specific Managed Database User Proxied For User in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_user_proxied_for_user
This data source provides details about a specific Managed Database User Proxied For User resource in Oracle Cloud Infrastructure Database Management service.

Gets the list of users on whose behalf the current user acts as proxy.

## Example Usage

```hcl
data "oci_database_management_managed_database_user_proxied_for_user" "test_managed_database_user_proxied_for_user" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	user_name = oci_identity_user.test_user.name

	#Optional
	name = var.managed_database_user_proxied_for_user_name
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `name` - (Optional) A filter to return only resources that match the entire name.
* `user_name` - (Required) The name of the user whose details are to be viewed.


## Attributes Reference

The following attributes are exported:

* `items` - An array of user resources.
	* `authentication` - Indicates whether the proxy is required to supply the client credentials (YES) or not (NO).
	* `flags` - The flags associated with the proxy/client pair.
	* `name` - The name of a proxy user or the name of the client user.

