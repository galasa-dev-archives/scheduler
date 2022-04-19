package util

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	yaml "gopkg.in/yaml.v3"

	schyaml "galasa.dev/scheduler/pkg/scheduleryaml"
)

type ApiConfiguration struct {
	Client    http.Client;
    BaseUrl   string;
}



func GetApiConfiguration() ApiConfiguration {

	var configuration ApiConfiguration

	configuration.Client = http.Client{}

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
			configuration.BaseUrl = configYaml.Url
		}
    } else if errors.Is(err, os.ErrNotExist) {
	} else {
	  panic(err)
	}

	return configuration
}