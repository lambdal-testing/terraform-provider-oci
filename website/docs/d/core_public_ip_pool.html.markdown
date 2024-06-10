---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_public_ip_pool"
sidebar_current: "docs-oci-datasource-core-public_ip_pool"
description: |-
  Provides details about a specific Public Ip Pool in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_public_ip_pool
This data source provides details about a specific Public Ip Pool resource in Oracle Cloud Infrastructure Core service.

Gets the specified `PublicIpPool` object. You must specify the object's [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
data "oci_core_public_ip_pool" "test_public_ip_pool" {
	#Required
	public_ip_pool_id = oci_core_public_ip_pool.test_public_ip_pool.id
}
```

## Argument Reference

The following arguments are supported:

* `public_ip_pool_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP pool.


## Attributes Reference

The following attributes are exported:

* `cidr_blocks` - The CIDR blocks added to this pool. This could be all or a portion of a BYOIP CIDR block. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing this pool. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP pool.
* `state` - The public IP pool's current state.
* `time_created` - The date and time the public IP pool was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

