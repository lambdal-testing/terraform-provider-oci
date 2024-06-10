---
subcategory: "Container Instances"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_container_instances_container_instance_shape"
sidebar_current: "docs-oci-datasource-container_instances-container_instance_shape"
description: |-
  Provides details about a specific Container Instance Shape in Oracle Cloud Infrastructure Container Instances service
---

# Data Source: oci_container_instances_container_instance_shape
This data source provides details about a specific Container Instance Shape resource in Oracle Cloud Infrastructure Container Instances service.

Get a list of shapes for creating Container Instances and their details.

## Example Usage

```hcl
data "oci_container_instances_container_instance_shape" "test_container_instance_shape" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.container_instance_shape_availability_domain
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The ID of the compartment in which to list resources.


## Attributes Reference

The following attributes are exported:

* `items` - List of shapes.
	* `memory_options` - For a flexible shape, the amount of memory available for instances that use this shape. 
		* `default_per_ocpu_in_gbs` - The default amount of memory per OCPU available for this shape, in gigabytes. 
		* `max_in_gbs` - The maximum amount of memory, in gigabytes. 
		* `max_per_ocpu_in_gbs` - The maximum amount of memory per OCPU available for this shape, in gigabytes. 
		* `min_in_gbs` - The minimum amount of memory, in gigabytes. 
		* `min_per_ocpu_in_gbs` - The minimum amount of memory per OCPU available for this shape, in gigabytes. 
	* `name` - The name identifying the shape. 
	* `networking_bandwidth_options` - For a flexible shape, the amount of networking bandwidth available for instances that use this shape. 
		* `default_per_ocpu_in_gbps` - The default amount of networking bandwidth per OCPU, in gigabits per second. 
		* `max_in_gbps` - The maximum amount of networking bandwidth, in gigabits per second. 
		* `min_in_gbps` - The minimum amount of networking bandwidth, in gigabits per second. 
	* `ocpu_options` - For a flexible shape, the number of OCPUs available for instances that use this shape. 
		* `max` - The maximum number of OCPUs. 
		* `min` - The minimum number of OCPUs. 
	* `processor_description` - A short description of the Instance's processor (CPU). 

