// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSshDetails Details of the SSH key that will be used. Required for source database Manual and UserManagerOci connection types.
// Not required for source container database connections.
type CreateSshDetails struct {

	// Name of the host the SSH key is valid for.
	Host *string `mandatory:"true" json:"host"`

	// Private SSH key string.
	Sshkey *string `mandatory:"true" json:"sshkey"`

	// SSH user
	User *string `mandatory:"true" json:"user"`

	// Sudo location
	SudoLocation *string `mandatory:"false" json:"sudoLocation"`
}

func (m CreateSshDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSshDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
