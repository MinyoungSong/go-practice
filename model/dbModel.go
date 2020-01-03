package model

import ()

type ClusterSchema struct {
	MetaData struct {
		ClusterName   string            `json:"clusterName"`
		CreatedBy     string            `json:"createdBy"`
		IsMainCluster bool              `json:"isMainCluster"`
		KubeVersion   string            `json:"kubeVersion"`
		Labels        map[string]string `json:"labels"`
		Provider      string            `json:"provider"`
	} `json:"metaData"`
	Status struct {
		Message string `json:"message"`
		Phase   string `json:"phase"`
		Reason  string `json:"reason"`
	} `json:"status"`
	ProvisionHistory []struct {
		JobID        int    `json:"jobId"`
		Status       string `json:"status"`
		CreatedTime  string `json:"createdTime"`
		ModifiedTime string `json:"modifiedTime"`
	} `json:"provisionHistory"`
	ProvisionResult *struct {
		ClusterName string      `json:"clusterName"`
		Errors      bool        `json:"errors"`
		Kubeconfig  interface{} `json:"kubeconfig"`
		Message     string      `json:"message"`
		Status      string      `json:"status"`
		Success     string      `json:"success"`
	} `json:"provisionResult"`
	ProvisionConfig *struct {
		CallbackAPI *struct {
			APIURL    string `json:"apiUrl" `
			BodyParam *struct {
				ClusterName string      `json:"clusterName"`
				Errors      bool        `json:"errors"`
				Kubeconfig  interface{} `json:"kubeconfig"`
				Message     string      `json:"message"`
				Status      string      `json:"status"`
				Success     string      `json:"success"`
			} `json:"bodyParam"`
			Host   string `json:"host"`
			Method string `json:"method"`
		} `json:"callbackAPI"`
		Logging *struct {
			CreateNamespace        bool   `json:"create_namespace"`
			DeploymentEnabled      bool   `json:"deployment_enabled"`
			DeploymentFile         string `json:"deployment_file"`
			DeploymentTemplateFile string `json:"deployment_template_file"`
			Fluentd                struct {
				Common struct {
					Image string `json:"image"`
					Tag   string `json:"tag"`
				} `json:"common"`
				Pub struct {
					TargetFluentdHost      string `json:"target_fluentd_host"`
					TargetFluentdPort      int    `json:"target_fluentd_port"`
					Varlibdockercontainers string `json:"varlibdockercontainers"`
				} `json:"pub"`
			} `json:"fluentd"`
			Install     bool   `json:"install"`
			Namespace   string `json:"namespace"`
			Releasename string `json:"releasename"`
		} `json:"logging"`
		Terraform *struct{} `json:"terraform"`
	} `json:"provisionConfig"`
	K8sAPIGroupVersionObj interface{} `json:"k8sApiGroupVersionObj"`
}
