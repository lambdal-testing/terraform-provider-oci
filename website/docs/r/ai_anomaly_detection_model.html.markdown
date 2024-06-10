---
subcategory: "Ai Anomaly Detection"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_anomaly_detection_model"
sidebar_current: "docs-oci-resource-ai_anomaly_detection-model"
description: |-
  Provides the Model resource in Oracle Cloud Infrastructure Ai Anomaly Detection service
---

# oci_ai_anomaly_detection_model
This resource provides the Model resource in Oracle Cloud Infrastructure Ai Anomaly Detection service.

Creates a new Model.


## Example Usage

```hcl
resource "oci_ai_anomaly_detection_model" "test_model" {
	#Required
	compartment_id = var.compartment_id
	model_training_details {
		#Required
		data_asset_ids = var.model_model_training_details_data_asset_ids

		#Optional
		algorithm_hint = var.model_model_training_details_algorithm_hint
		target_fap = var.model_model_training_details_target_fap
		training_fraction = var.model_model_training_details_training_fraction
		window_size = var.model_model_training_details_window_size
	}
	project_id = oci_ai_anomaly_detection_project.test_project.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.model_description
	display_name = var.model_display_name
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID for the ai model's compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A short description of the ai model.
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `model_training_details` - (Required) Specifies the details of the MSET model during the create call.
	* `algorithm_hint` - (Optional) User can choose specific algorithm for training.
	* `data_asset_ids` - (Required) The list of OCIDs of the data assets to train the model. The dataAssets have to be in the same project where the ai model would reside.
	* `target_fap` - (Optional) A target model accuracy metric user provides as their requirement
	* `training_fraction` - (Optional) Fraction of total data that is used for training the model. The remaining is used for validation of the model.
	* `window_size` - (Optional) This value would determine the window size of the training algorithm.
* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the model.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID for the model's compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A short description of the Model.
* `display_name` - A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the model that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `model_training_details` - Specifies the details of the MSET model during the create call.
	* `algorithm_hint` - User can choose specific algorithm for training.
	* `data_asset_ids` - The list of OCIDs of the data assets to train the model. The dataAssets have to be in the same project where the ai model would reside.
	* `target_fap` - A target model accuracy metric user provides as their requirement
	* `training_fraction` - Fraction of total data that is used for training the model. The remaining is used for validation of the model.
	* `window_size` - This value would determine the window size of the training algorithm.
* `model_training_results` - Specifies the details for an Anomaly Detection model trained with MSET.
	* `algorithm` - Actual algorithm used to train the model
	* `fap` - The final-achieved model accuracy metric on individual value level
	* `is_training_goal_achieved` - A boolean value to indicate if train goal/targetFap is achieved for trained model
	* `multivariate_fap` - The model accuracy metric on timestamp level.
	* `row_reduction_details` - Information regarding how/what row reduction methods will be applied. If this property is not present or is null, then it means row reduction is not applied.
		* `is_reduction_enabled` - A boolean value to indicate if row reduction is applied
		* `reduction_method` - Method for row reduction:
			* DELETE_ROW - delete rows with equal intervals
			* AVERAGE_ROW - average multiple rows to one row 
		* `reduction_percentage` - A percentage to reduce data size down to on top of original data
	* `signal_details` - The list of signal details.
		* `details` - detailed information for a signal.
		* `fap` - Accuracy metric for a signal.
		* `is_quantized` - A boolean value to indicate if a signal is quantized or not.
		* `max` - Max value within a signal.
		* `min` - Min value within a signal.
		* `mvi_ratio` - The ratio of missing values in a signal filled/imputed by the IDP algorithm.
		* `signal_name` - The name of a signal.
		* `status` - Status of the signal:
			* ACCEPTED - the signal is used for training the model
			* DROPPED - the signal does not meet requirement, and is dropped before training the model.
			* OTHER - placeholder for other status 
		* `std` - Standard deviation of values within a signal.
	* `warning` - A warning message to explain the reason when targetFap cannot be achieved for trained model
	* `window_size` - Window size defined during training or deduced by the algorithm.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the model.
* `state` - The state of the model.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the Model was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the Model was updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Model
	* `update` - (Defaults to 20 minutes), when updating the Model
	* `delete` - (Defaults to 20 minutes), when destroying the Model


## Import

Models can be imported using the `id`, e.g.

```
$ terraform import oci_ai_anomaly_detection_model.test_model "id"
```

