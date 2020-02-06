package model

import ()

type ResponseBody struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

//ZCPResources represents general multicluster resource response structure
type ZCPResources struct {
	ClusterName string      `json:"clusterName"`
	Values      interface{} `json:"values"`
}

//ZCPArrayResources represents general multicluster resource whose value is array
type ZCPArrayResources struct {
	ClusterName string                   `json:"clusterName"`
	Values      []map[string]interface{} `json:"values"`
}
