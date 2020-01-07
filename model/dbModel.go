package model

import (
	"time"

	"github.com/Kamva/mgm"
	kubeapi "k8s.io/client-go/tools/clientcmd/api/v1"
)

type Clusterprovisions struct {
	mgm.DefaultModel `bson:",inline"`
	MetaData         struct {
		ClusterName   string            `json:"clusterName"`
		CreatedBy     string            `json:"createdBy"`
		CreatedTime   time.Time         `json:"createdTime"`
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
	ProvisionHistory      *[]ProvisonHistoryObj `json:"provisionHistory"`
	ProvisionResult       *ProvisionResultObj   `json:"provisionResult"`
	ProvisionConfig       *ProvisionConfigObj   `json:"provisionConfig"`
	K8sAPIGroupVersionObj map[string]K8sAPIObj  `json:"k8sApiGroupVersionObj"`
}

type ProvisonHistoryObj struct {
	JobID        int       `json:"jobId"`
	Status       string    `json:"status"`
	CreatedTime  time.Time `json:"createdTime"`
	ModifiedTime time.Time `json:"modifiedTime"`
}

type ProvisionConfigObj struct {
	CallbackAPI struct {
		APIURL    string             `json:"apiUrl" `
		BodyParam ProvisionResultObj `json:"bodyParam"`
		Host      string             `json:"host"`
		Method    string             `json:"method"`
	} `json:"callbackAPI"`
	Logging struct {
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
	Terraform map[string]interface{} `json:"terraform"`
}

type ProvisionResultObj struct {
	ClusterName interface{}     `json:"clusterName"`
	Errors      interface{}     `json:"errors"`
	Kubeconfig  *kubeapi.Config `json:"kubeconfig"`
	Message     interface{}     `json:"message"`
	Status      interface{}     `json:"status"`
	Success     interface{}     `json:"success"`
}

type K8sAPIObj struct {
	GroupVersion string `json:"groupVersion"`
	Namespaced   bool   `json:"namespaced"`
	ResourceName string `json:"resourceName"`
	ResourceURL  string `json:"resourceUrl"`
}
