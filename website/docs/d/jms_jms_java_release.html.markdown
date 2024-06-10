---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_jms_java_release"
sidebar_current: "docs-oci-datasource-jms-jms_java_release"
description: |-
  Provides details about a specific Jms Java Release in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_jms_java_release
This data source provides details about a specific Jms Java Release resource in Oracle Cloud Infrastructure Jms service.

Returns detail of a Java release.

## Example Usage

```hcl
data "oci_jms_jms_java_release" "test_jms_java_release" {
	#Required
	release_version = var.jms_java_release_release_version
}
```

## Argument Reference

The following arguments are supported:

* `release_version` - (Required) Unique Java release version identifier


## Attributes Reference

The following attributes are exported:

* `artifacts` - List of Java artifacts.
	* `approximate_file_size_in_bytes` - Approximate compressed file size in bytes.
	* `artifact_content_type` - Product content type of this artifact.
	* `artifact_description` - Description of the binary artifact. Typically includes the OS, architecture, and installer type.
	* `artifact_id` - Unique identifier for the artifact.
	* `sha256` - SHA256 checksum of the artifact.
* `family_details` - Complete information of a specific Java release family. 
	* `display_name` - The display name of the release family.
	* `doc_url` - Link to access the documentation for the release.
	* `end_of_support_life_date` - The End of Support Life (EOSL) date of the Java release family (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 
	* `family_version` - The Java release family identifier.
	* `support_type` - This indicates the support category for the Java release family.
* `family_version` - Java release family identifier.
* `license_details` - Information about a license type for Java.
	* `display_name` - Commonly used name for the license type.
	* `license_type` - License Type
	* `license_url` - Publicly accessible license URL containing the detailed terms and conditions.
* `license_type` - License type for the Java version.
* `parent_release_version` - Parent Java release version identifier. This is applicable for BPR releases.
* `release_date` - The release date of the Java version (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `release_notes_url` - Release notes associated with the Java version.
* `release_type` - Release category of the Java version.
* `release_version` - Java release version identifier.
* `security_status` - The security status of the Java version.

