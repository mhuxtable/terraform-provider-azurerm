package labplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabPlanClient struct {
	Client *resourcemanager.Client
}

func NewLabPlanClientWithBaseURI(api environments.Api) (*LabPlanClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "labplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LabPlanClient: %+v", err)
	}

	return &LabPlanClient{
		Client: client,
	}, nil
}
