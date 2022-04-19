package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v3"

	"galasa.dev/scheduler/pkg/model"
	schyaml "galasa.dev/scheduler/pkg/scheduleryaml"
)


func Login(url string) {

    configuration := GetApiConfiguration()

	newUrl := url + "/graphql"

	var qlrequest model.GraphQlRequest
	qlrequest.Query = "query { serverStatus { apiReport } }"

	jsonBytes, err := json.Marshal(qlrequest)
    if err != nil {
        panic(err)
    }

	resp, err := configuration.Client.Post(newUrl, "application/json", bytes.NewBuffer(jsonBytes))
    if err != nil {
        panic(err)
    }
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

	var statusResponse model.ServerStatusResponse

	json.Unmarshal(body, &statusResponse)

	fmt.Printf("resp=%v\n", statusResponse.Data.ServerStatus.ApiReport)

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


	configYaml.Url = newUrl

    data, err := yaml.Marshal(&configYaml)
	if err != nil {
        panic( err )
    }

	err = ioutil.WriteFile(fileName, data, 0600)
	if err != nil {
        panic( err )
    }

}