---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_problem_entity"
sidebar_current: "docs-oci-datasource-cloud_guard-problem_entity"
description: |-
  Provides details about a specific Problem Entity in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_problem_entity
This data source provides details about a specific Problem Entity resource in Oracle Cloud Infrastructure Cloud Guard service.

Returns a list of entities for a CloudGuard Problem


## Example Usage

```hcl
data "oci_cloud_guard_problem_entity" "test_problem_entity" {
	#Required
	problem_id = oci_cloud_guard_problem.test_problem.id
}
```

## Argument Reference

The following arguments are supported:

* `problem_id` - (Required) OCId of the problem.


## Attributes Reference

The following attributes are exported:

* `items` - List of problem entities summaries related to a data source.
	* `entity_details` - List of event related to a DataSource
		* `display_name` - The display name of entity
		* `type` - Type of entity
		* `value` - The entity value
	* `problem_id` - Attached problem id
	* `regions` - Data source problem entities region
	* `result_url` - Log result query url for a data source query
	* `time_first_detected` - Data source problem entities first detected time
	* `time_last_detected` - Data source problem entities last detected time

