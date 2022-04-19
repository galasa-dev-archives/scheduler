package model

type GraphQlRequest struct {
	Query         string `json:"query"`
	Errors        []GraphQlError `json:"errors"`
}

type GraphQlError struct {
	Message       string `json:"message"`
}

