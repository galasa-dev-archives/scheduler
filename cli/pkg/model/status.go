package model

type ServerStatusResponse struct {
	Data            ServerStatusData `json:"data"`
	Errors        []GraphQlError     `json:"errors"`
}

type ServerStatusData struct {
	ServerStatus    ServerStatus `json:"serverStatus"`
}

type ServerStatus struct {
	SchedulerReport string `json:"schedulerReport"`
	ApiReport       string `json:"apiReport"`
}
