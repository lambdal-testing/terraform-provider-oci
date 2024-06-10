// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "subscribed_service_id" {}
variable "commitment_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_onesubscription_commitments" "test_commitments" {
  #Required
  compartment_id        = var.compartment_id
  subscribed_service_id = var.subscribed_service_id
}

data "oci_onesubscription_commitment" "test_commitment" {
  #Required
  commitment_id = var.commitment_id
}