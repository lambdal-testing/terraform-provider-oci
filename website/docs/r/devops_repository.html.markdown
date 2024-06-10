---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository"
sidebar_current: "docs-oci-resource-devops-repository"
description: |-
  Provides the Repository resource in Oracle Cloud Infrastructure Devops service
---

# oci_devops_repository
This resource provides the Repository resource in Oracle Cloud Infrastructure Devops service.

Creates a new repository.


## Example Usage

```hcl
resource "oci_devops_repository" "test_repository" {
	#Required
	name = var.repository_name
	project_id = oci_devops_project.test_project.id
	repository_type = var.repository_repository_type

	#Optional
	default_branch = var.repository_default_branch
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.repository_description
	freeform_tags = {"bar-key"= "value"}
	mirror_repository_config {

		#Optional
		connector_id = oci_devops_connector.test_connector.id
		repository_url = var.repository_mirror_repository_config_repository_url
		trigger_schedule {
			#Required
			schedule_type = var.repository_mirror_repository_config_trigger_schedule_schedule_type

			#Optional
			custom_schedule = var.repository_mirror_repository_config_trigger_schedule_custom_schedule
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `default_branch` - (Optional) (Updatable) The default branch of the repository.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `description` - (Optional) (Updatable) Details of the repository. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `mirror_repository_config` - (Optional) (Updatable) Configuration information for mirroring the repository.
	* `connector_id` - (Optional) (Updatable) Upstream git repository connection identifer.
	* `repository_url` - (Optional) (Updatable) URL of external repository you want to mirror.
	* `trigger_schedule` - (Optional) (Updatable) Specifies a trigger schedule. Timing information for when to initiate automated syncs.
		* `custom_schedule` - (Optional) (Updatable) Valid if type is CUSTOM. Following RFC 5545 recurrence rules, we can specify starting time, occurrence frequency, and interval size. Example for frequency could be DAILY/WEEKLY/HOURLY or any RFC 5545 supported frequency, which is followed by start time of this window. You can control the start time with BYHOUR, BYMINUTE and BYSECONDS. It is followed by the interval size. 
		* `schedule_type` - (Required) (Updatable) Different types of trigger schedule: NONE - No automated synchronization schedule. DEFAULT - Trigger schedule is every 30 minutes. CUSTOM - Custom triggering schedule. 
* `name` - (Required) (Updatable) Unique name of a repository.
* `project_id` - (Required) The OCID of the DevOps project containing the repository.
* `repository_type` - (Required) (Updatable) Type of repository. Allowed values:  `MIRRORED`  `HOSTED` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `branch_count` - The count of the branches present in the repository.
* `commit_count` - The count of the commits present in the repository.
* `compartment_id` - The OCID of the repository's compartment.
* `default_branch` - The default branch of the repository.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `description` - Details of the repository. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `http_url` - HTTP URL that you use to git clone, pull and push.
* `id` - The OCID of the repository. This value is unique and immutable.
* `lifecyle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `mirror_repository_config` - Configuration information for mirroring the repository.
	* `connector_id` - Upstream git repository connection identifer.
	* `repository_url` - URL of external repository you want to mirror.
	* `trigger_schedule` - Specifies a trigger schedule. Timing information for when to initiate automated syncs.
		* `custom_schedule` - Valid if type is CUSTOM. Following RFC 5545 recurrence rules, we can specify starting time, occurrence frequency, and interval size. Example for frequency could be DAILY/WEEKLY/HOURLY or any RFC 5545 supported frequency, which is followed by start time of this window. You can control the start time with BYHOUR, BYMINUTE and BYSECONDS. It is followed by the interval size. 
		* `schedule_type` - Different types of trigger schedule: NONE - No automated synchronization schedule. DEFAULT - Trigger schedule is every 30 minutes. CUSTOM - Custom triggering schedule. 
* `name` - Unique name of a repository. This value is mutable.
* `namespace` - Tenancy unique namespace.
* `project_id` - The OCID of the DevOps project containing the repository.
* `project_name` - Unique project name in a namespace.
* `repository_type` - Type of repository: MIRRORED - Repository created by mirroring an existing repository. HOSTED - Repository created and hosted using Oracle Cloud Infrastructure DevOps code repository. 
* `size_in_bytes` - The size of the repository in bytes.
* `ssh_url` - SSH URL that you use to git clone, pull and push.
* `state` - The current state of the repository.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time the repository was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `time_updated` - The time the repository was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `trigger_build_events` - Trigger build events supported for this repository: PUSH - Build is triggered when a push event occurs. COMMIT_UPDATES - Build is triggered when new commits are mirrored into a repository. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Repository
	* `update` - (Defaults to 20 minutes), when updating the Repository
	* `delete` - (Defaults to 20 minutes), when destroying the Repository


## Import

Repositories can be imported using the `id`, e.g.

```
$ terraform import oci_devops_repository.test_repository "id"
```

