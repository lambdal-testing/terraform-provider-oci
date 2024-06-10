// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_ons "github.com/oracle/oci-go-sdk/v65/ons"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_ons.NotificationControlPlaneClient", &OracleClient{InitClientFn: initOnsNotificationControlPlaneClient})
	RegisterOracleClient("oci_ons.NotificationDataPlaneClient", &OracleClient{InitClientFn: initOnsNotificationDataPlaneClient})
}

func initOnsNotificationControlPlaneClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_ons.NewNotificationControlPlaneClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) NotificationControlPlaneClient() *oci_ons.NotificationControlPlaneClient {
	return m.GetClient("oci_ons.NotificationControlPlaneClient").(*oci_ons.NotificationControlPlaneClient)
}

func initOnsNotificationDataPlaneClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_ons.NewNotificationDataPlaneClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) NotificationDataPlaneClient() *oci_ons.NotificationDataPlaneClient {
	return m.GetClient("oci_ons.NotificationDataPlaneClient").(*oci_ons.NotificationDataPlaneClient)
}
