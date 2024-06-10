---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_deployment_certificates"
sidebar_current: "docs-oci-datasource-golden_gate-deployment_certificates"
description: |-
  Provides the list of Deployment Certificates in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_deployment_certificates
This data source provides the list of Deployment Certificates in Oracle Cloud Infrastructure Golden Gate service.

Returns a list of certificates from truststore.

## Example Usage

```hcl
data "oci_golden_gate_deployment_certificates" "test_deployment_certificates" {
	#Required
	deployment_id = oci_golden_gate_deployment.test_deployment.id

	#Optional
	state = var.deployment_certificate_state
}
```

## Argument Reference

The following arguments are supported:

* `deployment_id` - (Required) A unique Deployment identifier. 
* `state` - (Optional) A filter to return only connections having the 'lifecycleState' given. 


## Attributes Reference

The following attributes are exported:

* `certificate_collection` - The list of certificate_collection.

### DeploymentCertificate Reference

The following attributes are exported:

* `authority_key_id` - The Certificate authority key id. 
* `certificate_content` - The base64 encoded content of the PEM file containing the SSL certificate. 
* `deployment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `is_ca` - Indicates if the certificate is ca. 
* `is_self_signed` - Indicates if the certificate is self signed. 
* `issuer` - The Certificate issuer. 
* `key` - The identifier key (unique name in the scope of the deployment) of the certificate being referenced.  It must be 1 to 32 characters long, must contain only alphanumeric characters and must start with a letter. 
* `md5hash` - The Certificate md5Hash. 
* `public_key` - The Certificate public key. 
* `public_key_algorithm` - The Certificate public key algorithm. 
* `public_key_size` - The Certificate public key size. 
* `serial` - The Certificate serial. 
* `sha1hash` - The Certificate sha1 hash. 
* `state` - Possible certificate lifecycle states. 
* `subject` - The Certificate subject. 
* `subject_key_id` - The Certificate subject key id. 
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_valid_from` - The time the certificate is valid from. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_valid_to` - The time the certificate is valid to. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `version` - The Certificate version. 

