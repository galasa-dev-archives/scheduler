/*
 * Copyright contributors to the Galasa project
 */
package status

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"galasa.dev/scheduler/pkg/model"
	"galasa.dev/scheduler/pkg/util"
)

func StatusReport() {
    configuration := util.GetApiConfiguration()
    
	var qlrequest model.GraphQlRequest
	qlrequest.Query = "query { serverStatus { apiReport schedulerReport } }"

	jsonBytes, err := json.Marshal(qlrequest)
    if err != nil {
        panic(err)
    }

	resp, err := configuration.Client.Post(configuration.BaseUrl, "application/json", bytes.NewBuffer(jsonBytes))
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
    
	fmt.Printf("API Status      : %v\n", statusResponse.Data.ServerStatus.ApiReport);
	fmt.Printf("Scheduler Status: %v\n", statusResponse.Data.ServerStatus.SchedulerReport);
}
