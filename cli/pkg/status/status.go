/*
 * Copyright contributors to the Galasa project
 */
package status

import (
	"context"
	"fmt"

	openapiclient "galasa.dev/scheduler/pkg/openapi"
	"galasa.dev/scheduler/pkg/util"
)

func StatusReport() {
    configuration := util.ContextConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatusApi.GetStatus(context.Background()).Execute()
    if err != nil {
        fmt.Printf("Error when calling `StatusApi.GetStatus``: %v\n", err)
        fmt.Printf("Full HTTP response: %v\n", r)
    }
    // response from `GetStatus`: Status
	fmt.Printf("API Status      : %v\n", resp.GetApiReport());
	fmt.Printf("Scheduler Status: %v\n", resp.GetSchedulerReport());
}
