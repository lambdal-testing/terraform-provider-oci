---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_discovery_analytic"
sidebar_current: "docs-oci-datasource-data_safe-discovery_analytic"
description: |-
  Provides details about a specific Discovery Analytic in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_discovery_analytic
This data source provides details about a specific Discovery Analytic resource in Oracle Cloud Infrastructure Data Safe service.

Gets consolidated discovery analytics data based on the specified query parameters.
If CompartmentIdInSubtreeQueryParam is specified as true, the behaviour
is equivalent to accessLevel "ACCESSIBLE" by default.


## Example Usage

```hcl
data "oci_data_safe_discovery_analytic" "test_discovery_analytic" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	compartment_id_in_subtree = var.discovery_analytic_compartment_id_in_subtree
	group_by = var.discovery_analytic_group_by
	sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `group_by` - (Optional) Attribute by which the discovery analytics data should be grouped.
* `sensitive_data_model_id` - (Optional) A filter to return only the resources that match the specified sensitive data model OCID.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.


## Attributes Reference

The following attributes are exported:

* `items` - An array of discovery analytics summary objects.
	* `count` - The total count for the aggregation metric.
	* `dimensions` - The scope of analytics data.
		* `sensitive_data_model_id` - The OCID of the sensitive data model.
		* `target_id` - The OCID of the target database.
	* `metric_name` - The name of the aggregation metric.

