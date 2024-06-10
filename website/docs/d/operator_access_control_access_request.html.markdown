---
subcategory: "Operator Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_operator_access_control_access_request"
sidebar_current: "docs-oci-datasource-operator_access_control-access_request"
description: |-
  Provides details about a specific Access Request in Oracle Cloud Infrastructure Operator Access Control service
---

# Data Source: oci_operator_access_control_access_request
This data source provides details about a specific Access Request resource in Oracle Cloud Infrastructure Operator Access Control service.

Gets details of an access request.

## Example Usage

```hcl
data "oci_operator_access_control_access_request" "test_access_request" {
	#Required
	access_request_id = oci_operator_access_control_access_request.test_access_request.id
}
```

## Argument Reference

The following arguments are supported:

* `access_request_id` - (Required) unique AccessRequest identifier


## Attributes Reference

The following attributes are exported:

* `access_reason_summary` - Summary comment by the operator creating the access request.
* `action_requests_list` - List of operator actions for which approval is sought by the operator user.
* `approver_comment` - The last recent Comment entered by the approver of the request.
* `approver_details` - Contains the user ids who have approved the accessRequest for extension.
	* `approval_action` - The action done by the approver.
	* `approval_additional_message` - Additional message specified by the approver of the request.
	* `approval_comment` - Comment specified by the approver of the request.
	* `approver_id` - The userId of the approver.
	* `time_approved_for_access` - Time for when the access request should start that is authorized by the customer in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z' 
	* `time_of_authorization` - Time when the access request was authorized by the customer in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z' 
* `audit_type` - Specifies the type of auditing to be enabled. There are two levels of auditing: command-level and keystroke-level.  By default, auditing is enabled at the command level i.e., each command issued by the operator is audited. When keystroke-level is chosen,  in addition to command level logging, key strokes are also logged. 
* `closure_comment` - The comment entered by the operator while closing the request.
* `compartment_id` - The OCID of the compartment that contains the access request.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. 
* `duration` - Duration in hours for which access is sought on the target resource.
* `extend_duration` - Duration in hours for which extension access is sought on the target resource.
* `extension_approver_details` - Contains the user ids who have approved the accessRequest for extension.
	* `approval_action` - The action done by the approver.
	* `approval_additional_message` - Additional message specified by the approver of the request.
	* `approval_comment` - Comment specified by the approver of the request.
	* `approver_id` - The userId of the approver.
	* `time_approved_for_access` - Time for when the access request should start that is authorized by the customer in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z' 
	* `time_of_authorization` - Time when the access request was authorized by the customer in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z' 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. 
* `id` - The OCID of the access request.
* `is_auto_approved` - Whether the access request was automatically approved.
* `is_validate_assignment` - Whether the access request was requested for Validate Assignment.
* `lifecycle_details` - more in detail about the lifeCycleState.
* `number_of_approvers` - Number of approvers who have authorized an access request.
* `number_of_approvers_required` - Number of approvers required to approve an access request.
* `number_of_extension_approvers` - Number of approvers who have authorized an access request for extension.
* `opctl_additional_message` - Additional message specific to the access request that can be specified by the approver at the time of approval.
* `opctl_id` - The OCID of the operator control governing the target resource.
* `opctl_name` - Name of the Operator control governing the target resource.
* `operator_id` - A unique identifier associated with the operator who raised the request. This identifier can not be used directly to identify the operator. You need to provide this identifier if you would like Oracle to provide additional information about the operator action within Oracle tenancy. 
* `reason` - Summary reason for which the operator is requesting access on the target resource.
* `request_id` - This is an automatic identifier generated by the system which is easier for human comprehension.
* `resource_id` - The OCID of the target resource associated with the access request. The operator raises an access request to get approval to  access the target resource. 
* `resource_name` - The name of the target resource.
* `resource_type` - resourceType for which the AccessRequest is applicable
* `severity` - Priority assigned to the access request by the operator
* `state` - The current state of the AccessRequest.
* `sub_resource_list` - The subresources requested for approval.
* `system_message` - System message that will be displayed to the operator at login to the target resource.
* `time_of_creation` - Time when the access request was created in [RFC 3339](https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z' 
* `time_of_modification` - Time when the access request was last modified in [RFC 3339](https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z' 
* `time_of_user_creation` - The time when access request is scheduled to be approved in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z' 
* `time_requested_for_future_access` - Time in future when the user for the access request needs to be created in [RFC 3339](https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z' 
* `user_id` - The OCID of the user that last modified the access request.
* `workflow_id` - The OCID of the workflow associated with the access request. This is needed if you want to contact Oracle Support for a stuck access request or for an access request that encounters an internal error. 

