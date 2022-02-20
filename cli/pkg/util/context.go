package util

import (
	"errors"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v3"

	openapiclient "galasa.dev/scheduler/pkg/openapi"
	schyaml "galasa.dev/scheduler/pkg/scheduleryaml"
)


func ContextConfiguration() *openapiclient.Configuration {

	configuration := openapiclient.NewConfiguration()

	dirname, err := os.UserHomeDir()
    if err != nil {
        panic( err )
    }

	fileName := dirname + "/.galasa/scheduler"
	if _, err := os.Stat(fileName); err == nil {
		var configYaml schyaml.Config
		yamlFile, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic( err )
		}
		err = yaml.Unmarshal(yamlFile, &configYaml)
		if err != nil {
			panic( err )
		}

		if configYaml.Url != "" {
			configuration.Servers = openapiclient.ServerConfigurations{
				{
					URL: configYaml.Url,
					Description: "No description provided",
				},
			}
		}
    } else if errors.Is(err, os.ErrNotExist) {
	} else {
	  panic(err)
	}

	return configuration
}