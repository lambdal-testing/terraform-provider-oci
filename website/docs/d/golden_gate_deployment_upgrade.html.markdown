---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_deployment_upgrade"
sidebar_current: "docs-oci-datasource-golden_gate-deployment_upgrade"
description: |-
  Provides details about a specific Deployment Upgrade in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_deployment_upgrade
This data source provides details about a specific Deployment Upgrade resource in Oracle Cloud Infrastructure Golden Gate service.

Retrieves a deployment upgrade.


## Example Usage

```hcl
data "oci_golden_gate_deployment_upgrade" "test_deployment_upgrade" {
	#Required
	deployment_upgrade_id = oci_golden_gate_deployment_upgrade.test_deployment_upgrade.id
}
```

## Argument Reference

The following arguments are supported:

* `deployment_upgrade_id` - (Required) A unique Deployment Upgrade identifier. 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `defined_tags` - Tags defined for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace.bar-key": "value"}` 
* `deployment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `deployment_upgrade_type` - The type of the deployment upgrade: MANUAL or AUTOMATIC 
* `description` - Metadata about this specific object. 
* `display_name` - An object's Display Name. 
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment upgrade being referenced. 
* `is_cancel_allowed` - Indicates if cancel is allowed. Scheduled upgrade can be cancelled only if target version is not forced by service,  otherwise only reschedule allowed. 
* `is_reschedule_allowed` - Indicates if reschedule is allowed. Upgrade can be rescheduled postponed until the end of the service defined auto-upgrade period. 
* `is_rollback_allowed` - Indicates if rollback is allowed. In practice only the last upgrade can be rolled back.
	* Manual upgrade is allowed to rollback only until the old version isn't deprecated yet.
	* Automatic upgrade by default is not allowed, unless a serious issue does not justify. 
* `is_security_fix` - Indicates if OGG release contains security fix. 
* `is_snoozed` - Indicates if upgrade notifications are snoozed or not. 
* `lifecycle_details` - Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state. 
* `lifecycle_sub_state` - Possible GGS lifecycle sub-states. 
* `ogg_version` - Version of OGG 
* `previous_ogg_version` - Version of OGG 
* `release_type` - The type of release. 
* `state` - Possible lifecycle states. 
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_finished` - The date and time the request was finished. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_ogg_version_supported_until` - The time until OGG version is supported. After this date has passed OGG version will not be available anymore. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_released` - The time the resource was released. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_schedule` - The time of upgrade schedule. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_schedule_max` - Indicates the latest time until the deployment upgrade could be rescheduled. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_snoozed_until` - The time the upgrade notifications are snoozed until. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_started` - The date and time the request was started. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 

