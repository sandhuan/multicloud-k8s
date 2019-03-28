/*
Copyright 2018 Intel Corporation.
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

package app

import (
	"os"
	"plugin"
	"reflect"
	"testing"

	utils "k8splugin/internal"

	pkgerrors "github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
)

func LoadMockPlugins(krdLoadedPlugins map[string]*plugin.Plugin) error {
	if _, err := os.Stat("../../mock_files/mock_plugins/mockplugin.so"); os.IsNotExist(err) {
		return pkgerrors.New("mockplugin.so does not exist. Please compile mockplugin.go to generate")
	}

	mockPlugin, err := plugin.Open("../../mock_files/mock_plugins/mockplugin.so")
	if err != nil {
		return pkgerrors.Wrap(err, "Opening mock plugins")
	}

	krdLoadedPlugins["namespace"] = mockPlugin
	krdLoadedPlugins["deployment"] = mockPlugin
	krdLoadedPlugins["service"] = mockPlugin

	return nil
}

func TestInit(t *testing.T) {
	t.Run("Successfully create Kube Client", func(t *testing.T) {

		kubeClient := kubernetesClient{}
		err := kubeClient.init("../../mock_files/mock_configs/mock_config")
		if err != nil {
			t.Fatalf("TestGetKubeClient returned an error (%s)", err)
		}

		name := reflect.TypeOf(kubeClient.clientSet).Elem().Name()
		if name != "Clientset" {
			t.Fatalf("TestGetKubeClient returned :\n result=%v\n expected=%v", name, "Clientset")
		}

	})
}

func TestCreateResources(t *testing.T) {
	oldkrdPluginData := utils.LoadedPlugins

	defer func() {
		utils.LoadedPlugins = oldkrdPluginData
	}()

	err := LoadMockPlugins(utils.LoadedPlugins)
	if err != nil {
		t.Fatalf("LoadMockPlugins returned an error (%s)", err)
	}

	k8 := kubernetesClient{
		clientSet: &kubernetes.Clientset{},
	}

	t.Run("Successfully delete resources", func(t *testing.T) {
		data := map[string][]string{
			"Deployment": []string{"../../mock_files/mock_yamls/deployment.yaml"},
			"Service":    []string{"../../mock_files/mock_yamls/service.yaml"},
		}

		_, err := k8.createResources(data, "testnamespace")
		if err != nil {
			t.Fatalf("TestCreateResources returned an error (%s)", err)
		}
	})
}

func TestDeleteResources(t *testing.T) {
	oldkrdPluginData := utils.LoadedPlugins

	defer func() {
		utils.LoadedPlugins = oldkrdPluginData
	}()

	err := LoadMockPlugins(utils.LoadedPlugins)
	if err != nil {
		t.Fatalf("LoadMockPlugins returned an error (%s)", err)
	}

	k8 := kubernetesClient{
		clientSet: &kubernetes.Clientset{},
	}

	t.Run("Successfully delete resources", func(t *testing.T) {
		data := map[string][]string{
			"Deployment": []string{"deployment-1", "deployment-2"},
			"Service":    []string{"service-1", "service-2"},
		}

		err := k8.deleteResources(data, "test")
		if err != nil {
			t.Fatalf("TestCreateVNF returned an error (%s)", err)
		}
	})
}
