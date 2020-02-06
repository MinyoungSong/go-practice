package model

import ()

type ReqK8sGetResources struct {
	Condition struct {
		AsTable   bool     `json:"asTable"`
		Clusters  []string `json:"clusters"`
		Kind      string   `json:"kind"`
		Labels    string   `json:"labels"`
		Name      string   `json:"name"`
		Namespace string   `json:"namespace"`
	} `json:"condition"`
}
