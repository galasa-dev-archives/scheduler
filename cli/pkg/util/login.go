package util

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v3"

	openapiclient "galasa.dev/scheduler/pkg/openapi"
	schyaml "galasa.dev/scheduler/pkg/scheduleryaml"
)


func Login(url string) {

    configuration := openapiclient.NewConfiguration()
	configuration.Servers = openapiclient.ServerConfigurations{
		{
			URL: url,
			Description: "No description provided",
		},
	}

    apiClient := openapiclient.NewAPIClient(configuration)
    _, _, err := apiClient.StatusApi.GetStatus(context.Background()).Execute()
    if err != nil {
        panic(err)
    }

	fmt.Println("Successfully logged onto api server");

	dirname, err := os.UserHomeDir()
    if err != nil {
        panic( err )
    }

	dirname = dirname + "/.galasa"
	err = os.MkdirAll(dirname, os.ModePerm)
    if err != nil {
        panic( err )
    }

	fileName := dirname + "/scheduler"
	var configYaml schyaml.Config
	if _, err := os.Stat(fileName); err == nil {
		yamlFile, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic( err )
		}
		err = yaml.Unmarshal(yamlFile, &configYaml)
		if err != nil {
			panic( err )
		}
    } else if errors.Is(err, os.ErrNotExist) {
	} else {
	  panic(err)
	}


	configYaml.Url = url

    data, err := yaml.Marshal(&configYaml)
	if err != nil {
        panic( err )
    }

	err = ioutil.WriteFile(fileName, data, 0600)
	if err != nil {
        panic( err )
    }

}