/*
Copyright 2023 The Radius Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package status

import (
	"context"
	"testing"

	"go.uber.org/mock/gomock"
	"github.com/radius-project/radius/pkg/cli/clients"
	"github.com/radius-project/radius/pkg/cli/clients_new/generated"
	"github.com/radius-project/radius/pkg/cli/clierrors"
	"github.com/radius-project/radius/pkg/cli/config"
	"github.com/radius-project/radius/pkg/cli/connections"
	"github.com/radius-project/radius/pkg/cli/framework"
	"github.com/radius-project/radius/pkg/cli/output"
	"github.com/radius-project/radius/pkg/cli/workspaces"
	"github.com/radius-project/radius/pkg/corerp/api/v20231001preview"
	"github.com/radius-project/radius/pkg/to"
	"github.com/radius-project/radius/pkg/ucp/resources"
	"github.com/radius-project/radius/test/radcli"
	"github.com/stretchr/testify/require"
)

func Test_CommandValidation(t *testing.T) {
	radcli.SharedCommandValidation(t, NewCommand)
}

func Test_Validate(t *testing.T) {
	testcases := []radcli.ValidateInput{
		{
			Name:          "Status Command with default application",
			Input:         []string{},
			ExpectedValid: true,
			ConfigHolder: framework.ConfigHolder{
				ConfigFilePath: "",
				Config:         radcli.LoadConfigWithWorkspace(t),
				DirectoryConfig: &config.DirectoryConfig{
					Workspace: config.DirectoryWorkspaceConfig{
						Application: "test-application",
					},
				},
			},
		},
		{
			Name:          "Status Command with flag",
			Input:         []string{"-a", "test-app"},
			ExpectedValid: true,
			ConfigHolder: framework.ConfigHolder{
				ConfigFilePath: "",
				Config:         radcli.LoadConfigWithWorkspace(t),
			},
		},
		{
			Name:          "Status Command with positional arg",
			Input:         []string{"test-app"},
			ExpectedValid: true,
			ConfigHolder: framework.ConfigHolder{
				ConfigFilePath: "",
				Config:         radcli.LoadConfigWithWorkspace(t),
			},
		},
		{
			Name:          "Status Command with fallback workspace",
			Input:         []string{"--application", "test-app", "--group", "test-group"},
			ExpectedValid: true,
			ConfigHolder: framework.ConfigHolder{
				ConfigFilePath: "",
				Config:         radcli.LoadEmptyConfig(t),
			},
		},
		{
			Name:          "Status Command with incorrect args",
			Input:         []string{"foo", "bar"},
			ExpectedValid: false,
			ConfigHolder: framework.ConfigHolder{
				ConfigFilePath: "",
				Config:         radcli.LoadConfigWithWorkspace(t),
			},
		},
	}
	radcli.SharedValidateValidation(t, NewCommand, testcases)
}

func Test_Run(t *testing.T) {
	t.Run("Success: Application Found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		application := v20231001preview.ApplicationResource{
			Name: to.Ptr("test-app"),
		}

		appManagementClient := clients.NewMockApplicationsManagementClient(ctrl)
		appManagementClient.EXPECT().
			ShowApplication(gomock.Any(), "test-app").
			Return(application, nil).
			Times(1)

		resourceList := []generated.GenericResource{
			{
				Name: to.Ptr("test-container"),
				ID:   to.Ptr("/planes/radius/local/resourceGroups/test-group/providers/Applications.Core/containers/test-container"),
			},
			{
				Name: to.Ptr("test-gateway"),
				ID:   to.Ptr("/planes/radius/local/resourceGroups/test-group/providers/Applications.Core/gateways/test-gateway"),
			},
		}

		appManagementClient.EXPECT().
			ListAllResourcesByApplication(gomock.Any(), "test-app").
			Return(resourceList, nil).
			Times(1)

		diagnosticsClient := clients.NewMockDiagnosticsClient(ctrl)
		diagnosticsClient.EXPECT().
			GetPublicEndpoint(gomock.Any(), clients.EndpointOptions{ResourceID: mustParse(t, "/planes/radius/local/resourceGroups/test-group/providers/Applications.Core/containers/test-container")}).
			Return(nil, nil).
			Times(1)

		diagnosticsClient.EXPECT().
			GetPublicEndpoint(gomock.Any(), clients.EndpointOptions{ResourceID: mustParse(t, "/planes/radius/local/resourceGroups/test-group/providers/Applications.Core/gateways/test-gateway")}).
			Return(to.Ptr("http://some-url.example.com"), nil).
			Times(1)

		workspace := &workspaces.Workspace{
			Connection: map[string]any{
				"kind":    "kubernetes",
				"context": "kind-kind",
			},
			Name:  "kind-kind",
			Scope: "/planes/radius/local/resourceGroups/test-group",
		}
		outputSink := &output.MockOutput{}
		runner := &Runner{
			ConnectionFactory: &connections.MockFactory{
				ApplicationsManagementClient: appManagementClient,
				DiagnosticsClient:            diagnosticsClient,
			},
			Workspace:       workspace,
			Format:          "table",
			Output:          outputSink,
			ApplicationName: "test-app",
		}

		err := runner.Run(context.Background())
		require.NoError(t, err)

		applicationStatus := clients.ApplicationStatus{
			Name:          "test-app",
			ResourceCount: 2,
			Gateways: []clients.GatewayStatus{
				{
					Name:     "test-gateway",
					Endpoint: "http://some-url.example.com",
				},
			},
		}

		expected := []any{
			output.FormattedOutput{
				Format:  "table",
				Obj:     applicationStatus,
				Options: statusFormat(),
			},
			output.LogOutput{
				Format: "",
			},
			output.FormattedOutput{
				Format:  "table",
				Obj:     applicationStatus.Gateways,
				Options: gatewayFormat(),
			},
		}

		require.Equal(t, expected, outputSink.Writes)
	})

	t.Run("Error: Application Not Found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appManagementClient := clients.NewMockApplicationsManagementClient(ctrl)
		appManagementClient.EXPECT().
			ShowApplication(gomock.Any(), "test-app").
			Return(v20231001preview.ApplicationResource{}, radcli.Create404Error()).
			Times(1)

		workspace := &workspaces.Workspace{
			Connection: map[string]any{
				"kind":    "kubernetes",
				"context": "kind-kind",
			},
			Name:  "kind-kind",
			Scope: "/planes/radius/local/resourceGroups/test-group",
		}
		outputSink := &output.MockOutput{}
		runner := &Runner{
			ConnectionFactory: &connections.MockFactory{ApplicationsManagementClient: appManagementClient},
			Workspace:         workspace,
			Format:            "table",
			Output:            outputSink,
			ApplicationName:   "test-app",
		}

		err := runner.Run(context.Background())
		require.Error(t, err)
		require.Equal(t, clierrors.Message("The application \"test-app\" was not found or has been deleted."), err)

		require.Empty(t, outputSink.Writes)
	})
}

func mustParse(t *testing.T, s string) resources.ID {
	t.Helper()
	id, err := resources.Parse(s)
	require.NoError(t, err)
	return id
}
