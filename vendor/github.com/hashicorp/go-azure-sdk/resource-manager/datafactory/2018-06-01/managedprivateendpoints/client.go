package managedprivateendpoints

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedPrivateEndpointsClient struct {
	Client *resourcemanager.Client
}

func NewManagedPrivateEndpointsClientWithBaseURI(api environments.Api) (*ManagedPrivateEndpointsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "managedprivateendpoints", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedPrivateEndpointsClient: %+v", err)
	}

	return &ManagedPrivateEndpointsClient{
		Client: client,
	}, nil
}
