package model

type ServerStatusResponse struct {
	Data            ServerStatusData `json:"data"`
}

type ServerStatusData struct {
	ServerStatus    ServerStatus `json:"serverStatus"`
}

type ServerStatus struct {
	SchedulerReport string `json:"schedulerReport"`
	ApiReport       string `json:"apiReport"`
}