---
subcategory: "Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_monitoring_alarms"
sidebar_current: "docs-oci-datasource-monitoring-alarms"
description: |-
  Provides the list of Alarms in Oracle Cloud Infrastructure Monitoring service
---

# Data Source: oci_monitoring_alarms
This data source provides the list of Alarms in Oracle Cloud Infrastructure Monitoring service.

Lists the alarms for the specified compartment.
For more information, see
[Listing Alarms](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/list-alarm.htm).
For important limits information, see
[Limits on Monitoring](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).

This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
or transactions, per second (TPS) for a given tenancy.


## Example Usage

```hcl
data "oci_monitoring_alarms" "test_alarms" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	compartment_id_in_subtree = var.alarm_compartment_id_in_subtree
	display_name = var.alarm_display_name
	state = var.alarm_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the resources monitored by the metric that you are searching for. Use tenancyId to search in the root compartment.  Example: `ocid1.compartment.oc1..exampleuniqueID` 
* `compartment_id_in_subtree` - (Optional) When true, returns resources from all compartments and subcompartments. The parameter can only be set to true when compartmentId is the tenancy OCID (the tenancy is the root compartment). A true value requires the user to have tenancy-level permissions. If this requirement is not met, then the call is rejected. When false, returns resources from only the compartment specified in compartmentId. Default is false. 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. Use this filter to list an alarm by name. Alternatively, when you know the alarm OCID, use the GetAlarm operation. 
* `state` - (Optional) A filter to return only alarms that match the given lifecycle state exactly. When not specified, only alarms in the ACTIVE lifecycle state are listed. 


## Attributes Reference

The following attributes are exported:

* `alarms` - The list of alarms.

### Alarm Reference

The following attributes are exported:

* `body` - The human-readable content of the delivered alarm notification. Oracle recommends providing guidance to operators for resolving the alarm condition. Consider adding links to standard runbook practices. Avoid entering confidential information.  Example: `High CPU usage alert. Follow runbook instructions for resolution.` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the alarm. 
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"Operations.CostCenter": "42"}` 
* `destinations` - A list of destinations for alarm notifications. Each destination is represented by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a related resource, such as a [topic](https://docs.cloud.oracle.com/iaas/api/#/en/notification/latest/NotificationTopic). Supported destination services: Notifications, Streaming. Limit: One destination per supported destination service. 
* `display_name` - A user-friendly name for the alarm. It does not have to be unique, and it's changeable.

	This value determines the title of each alarm notification.

	Example: `High CPU Utilization` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm. 
* `is_enabled` - Whether the alarm is enabled.  Example: `true` 
* `is_notifications_per_metric_dimension_enabled` - Whether the alarm sends a separate message for each metric stream. See [Creating an Alarm That Splits Messages by Metric Stream](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/create-alarm-split.htm). Example: `true` 
* `message_format` - The format to use for alarm notifications. The formats are:
	* `RAW` - Raw JSON blob. Default value. When the `destinations` attribute specifies `Streaming`, all alarm notifications use this format.
	* `PRETTY_JSON`: JSON with new lines and indents. Available when the `destinations` attribute specifies `Notifications` only.
	* `ONS_OPTIMIZED`: Simplified, user-friendly layout. Available when the `destinations` attribute specifies `Notifications` only. Applies to Email subscription types only. 
* `metric_compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the metric being evaluated by the alarm. 
* `metric_compartment_id_in_subtree` - When true, the alarm evaluates metrics from all compartments and subcompartments. The parameter can only be set to true when metricCompartmentId is the tenancy OCID (the tenancy is the root compartment). A true value requires the user to have tenancy-level permissions. If this requirement is not met, then the call is rejected. When false, the alarm evaluates metrics from only the compartment specified in metricCompartmentId. Default is false.  Example: `true` 
* `namespace` - The source service or application emitting the metric that is evaluated by the alarm.  Example: `oci_computeagent` 
* `notification_version` - The version of the alarm notification to be delivered. Allowed value: `1.X` The value must start with a number (up to four digits), followed by a period and an uppercase X. 
* `overrides` - A set of overrides that control evaluations of the alarm. 

	Each override can specify values for query, severity, body, and pending duration. When an alarm contains overrides, the Monitoring service evaluates each override in order, beginning with the first override in the array (index position `0`), and then evaluates the alarm's base values (`ruleName` value of `BASE`). 
	* `body` - The human-readable content of the delivered alarm notification. Oracle recommends providing guidance to operators for resolving the alarm condition. Consider adding links to standard runbook practices. Avoid entering confidential information.  Example: `High CPU usage alert. Follow runbook instructions for resolution.` 
	* `pending_duration` - The period of time that the condition defined in the alarm must persist before the alarm state changes from "OK" to "FIRING". For example, a value of 5 minutes means that the alarm must persist in breaching the condition for five minutes before the alarm updates its state to "FIRING".

		The duration is specified as a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H` for one hour). Minimum: PT1M. Maximum: PT1H. Default: PT1M.

		Under the default value of PT1M, the first evaluation that breaches the alarm updates the state to "FIRING".

		The alarm updates its status to "OK" when the breaching condition has been clear for the most recent minute.

		Example: `PT5M` 
	* `query` - The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of the Monitoring service interprets results for each returned time series as Boolean values, where zero represents false and a non-zero value represents true. A true value means that the trigger rule condition has been met. The query must specify a metric, statistic, interval, and trigger rule (threshold or absence). Supported values for interval depend on the specified time range. More interval values are supported for smaller time ranges. You can optionally specify dimensions and grouping functions. Supported grouping functions: `grouping()`, `groupBy()`. For information about writing MQL expressions, see [Editing the MQL Expression for a Query](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/query-metric-mql.htm). For details about MQL, see [Monitoring Query Language (MQL) Reference](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Reference/mql.htm). For available dimensions, review the metric definition for the supported service. See [Supported Services](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#SupportedServices).

		Example of threshold alarm:

		-----

		CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.groupBy(availabilityDomain).percentile(0.9) > 85

		-----

		Example of absence alarm:

		-----

		CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.absent()

		----- 
	* `rule_name` - A user-friendly description for this alarm override. Must be unique across all `ruleName` values for the alarm. 
	* `severity` - The perceived severity of the alarm with regard to the affected system.  Example: `CRITICAL` 
* `pending_duration` - The period of time that the condition defined in the alarm must persist before the alarm state changes from "OK" to "FIRING". For example, a value of 5 minutes means that the alarm must persist in breaching the condition for five minutes before the alarm updates its state to "FIRING".

	The duration is specified as a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H` for one hour). Minimum: PT1M. Maximum: PT1H. Default: PT1M.

	Under the default value of PT1M, the first evaluation that breaches the alarm updates the state to "FIRING".

	The alarm updates its status to "OK" when the breaching condition has been clear for the most recent minute.

	Example: `PT5M` 
* `query` - The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of the Monitoring service interprets results for each returned time series as Boolean values, where zero represents false and a non-zero value represents true. A true value means that the trigger rule condition has been met. The query must specify a metric, statistic, interval, and trigger rule (threshold or absence). Supported values for interval depend on the specified time range. More interval values are supported for smaller time ranges. You can optionally specify dimensions and grouping functions. Supported grouping functions: `grouping()`, `groupBy()`. For information about writing MQL expressions, see [Editing the MQL Expression for a Query](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/query-metric-mql.htm). For details about MQL, see [Monitoring Query Language (MQL) Reference](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Reference/mql.htm). For available dimensions, review the metric definition for the supported service. See [Supported Services](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#SupportedServices).

	Example of threshold alarm:

	-----

	CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.groupBy(availabilityDomain).percentile(0.9) > 85

	-----

	Example of absence alarm:

	-----

	CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.absent()

	----- 
* `repeat_notification_duration` - The frequency for re-submitting alarm notifications, if the alarm keeps firing without interruption. Format defined by ISO 8601. For example, `PT4H` indicates four hours. Minimum: PT1M. Maximum: P30D.

	Default value: null (notifications are not re-submitted).

	Example: `PT2H` 
* `resolution` - The time between calculated aggregation windows for the alarm. Supported value: `1m` 
* `resource_group` - Resource group to match for metric data retrieved by the alarm. A resource group is a custom string that you can match when retrieving custom metrics. Only one resource group can be applied per metric. A valid resourceGroup value starts with an alphabetical character and includes only alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($).  Example: `frontend-fleet` 
* `rule_name` - Identifier of the alarm's base values for alarm evaluation, for use when the alarm contains overrides.  A valid ruleName value starts with an alphabetic character and includes only alphanumeric characters, underscores and square brackets.  Minimum number of characters: 3. Default value is `BASE`. For information about alarm overrides, see [AlarmOverride](https://docs.cloud.oracle.com/iaas/api/#/en/monitoring/latest/datatypes/AlarmOverride). 
* `severity` - The perceived type of response required when the alarm is in the "FIRING" state.  Example: `CRITICAL` 
* `state` - The current lifecycle state of the alarm.  Example: `DELETED` 
* `suppression` - The configuration details for suppressing an alarm. 
	* `description` - Human-readable reason for suppressing alarm notifications. It does not have to be unique, and it's changeable. Avoid entering confidential information.

		Oracle recommends including tracking information for the event or associated work, such as a ticket number.

		Example: `Planned outage due to change IT-1234.` 
	* `time_suppress_from` - The start date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2023-02-01T01:02:29.600Z` 
	* `time_suppress_until` - The end date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2023-02-01T02:02:29.600Z` 
* `time_created` - The date and time the alarm was created. Format defined by RFC3339.  Example: `2023-02-01T01:02:29.600Z` 
* `time_updated` - The date and time the alarm was last updated. Format defined by RFC3339.  Example: `2023-02-03T01:02:29.600Z` 

