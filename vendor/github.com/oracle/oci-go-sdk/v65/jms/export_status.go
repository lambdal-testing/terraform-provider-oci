// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExportStatus Attributes of fleet's export status.
type ExportStatus struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the fleet.
	FleetId *string `mandatory:"true" json:"fleetId"`

	// The date and time of the last export run.
	TimeLastRun *common.SDKTime `mandatory:"true" json:"timeLastRun"`

	// The date and time of the next export run.
	TimeNextRun *common.SDKTime `mandatory:"true" json:"timeNextRun"`

	// The status of the latest export run.
	LatestRunStatus ExportRunStatusEnum `mandatory:"true" json:"latestRunStatus"`
}

func (m ExportStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExportRunStatusEnum(string(m.LatestRunStatus)); !ok && m.LatestRunStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LatestRunStatus: %s. Supported values are: %s.", m.LatestRunStatus, strings.Join(GetExportRunStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
