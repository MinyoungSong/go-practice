package resources

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"k8s.io/client-go/dynamic"
	rest "k8s.io/client-go/rest"
	clientcmd "k8s.io/client-go/tools/clientcmd"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"skcloud.io/cloudzcp/zcpctl-backend/db"
	"skcloud.io/cloudzcp/zcpctl-backend/model"
	"skcloud.io/cloudzcp/zcpctl-backend/util"
)

// GetClsuterList - Cluster 정보 조회
func GetK8sResources() echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		type resultItem struct {
			ClusterName string      `"json:clusterName"`
			Values      interface{} `"json:values"`
		}

		resultArr := []model.ZCPResources{}

		param := new(model.ReqK8sGetResources)

		if err = c.Bind(param); err != nil {
			return err
		}

		clusters := param.Condition.Clusters
		namespace := param.Condition.Namespace
		kind := param.Condition.Kind
		name := param.Condition.Name
		labels := param.Condition.Labels
		asTable := param.Condition.AsTable

		clusterConfigList, err := GetK8sClientConfig(clusters, c)
		if err != nil {
			return err
		}

		log.Println(clusterConfigList)

		for _, config := range clusterConfigList {

			absPath := ""

			k8sObj, err := config.K8sAPIGroupVersionObj[kind]
			if !err {
				return errors.New("k8sObj is null")
			}

			absPath += k8sObj.ResourceURL

			if !util.IsEmptyString(namespace) && k8sObj.Namespaced {
				absPath += "/namespaces/" + namespace
			}

			absPath += "/" + k8sObj.ResourceName

			if !util.IsEmptyString(name) {
				absPath += "/" + name
			}

			resClient := config.ClientObj.Verb("GET").AbsPath(absPath)

			if asTable {
				resClient = resClient.SetHeader("Accecpt", "application/json;as=Table;g=meta.k8s.io;v=v1beta1, application/json")
			}

			if !util.IsEmptyString(labels) {
				resClient = resClient.Param("labelSelector", labels)
			}

			elem := model.ZCPResources{}
			elem.ClusterName = config.ClusterName

			resultStr, errs := resClient.DoRaw()
			if errs != nil {
				elem.Values = errs
				resultArr = append(resultArr, elem)

			} else {
				json.Unmarshal(resultStr, &elem.Values)
				resultArr = append(resultArr, elem)

			}

		}

		// if err != nil {
		// 	// log.Fatal(err)
		// 	return c.JSON(http.StatusInternalServerError, util.SetSuccessFalse(err.Error(), err))
		// }

		return c.JSON(http.StatusOK, util.SetSuccessTrue(resultArr))

	}

}

type ClusterConfigMap struct {
	ClusterName           string                     `json:"clusterName"`
	ClientObj             *rest.RESTClient           `json:"clientObj"`
	K8sAPIGroupVersionObj map[string]model.K8sAPIObj `json:"k8sApiGroupVersionObj"`
}

func GetK8sClientConfig(clusterArr []string, c echo.Context) ([]ClusterConfigMap, error) {

	clusterConfigList := []ClusterConfigMap{}

	filter := bson.D{{"metaData.clusterName", bson.D{{"$in", clusterArr}}}}

	cur, err := db.Select(filter)
	if err != nil {
		return clusterConfigList, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem model.Clusterprovisions
		var configmap ClusterConfigMap
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		// var kubeConfig kubeapi.Config
		kubeConfigStr, err := json.Marshal(elem.ProvisionResult.Kubeconfig)
		if err != nil {
			return clusterConfigList, err
		}

		clientcmd.BuildConfigFromFlags()

		kubeRestConfig, err := clientcmd.RESTConfigFromKubeConfig(kubeConfigStr)
		if err != nil {
			return clusterConfigList, err
		}

		dynamicConfig, _ := dynamic.NewForConfig(kubeRestConfig)

		kubeRestConfig.NegotiatedSerializer = dynamic.basicNegotiatedSerializer{}

		resClient, err := rest.UnversionedRESTClientFor(kubeRestConfig)
		if err != nil {
			return clusterConfigList, err
		}

		configmap.ClientObj = resClient
		configmap.ClusterName = elem.MetaData.ClusterName
		configmap.K8sAPIGroupVersionObj = elem.K8sAPIGroupVersionObj

		clusterConfigList = append(clusterConfigList, configmap)

	}
	// Close the cursor once finished
	cur.Close(context.TODO())

	return clusterConfigList, err

}

func GenerateReqPath(config ClusterConfigMap, param model.ReqK8sGetResources) (rest.Request, error) {

	request := rest.Request{}

	return request, nil
}

func GetK8sClientConfig2(clusterArr []string, c echo.Context) []model.ZCPResources {

	// var resultConfigMap map[string]interface{}
	// resultConfigMap = make(map[string]interface{})

	resultMapArr := []model.ZCPResources{}

	// log.Println(kubeConfigArr)

	// resultArr := []model.Clusterprovisions{}

	filter := bson.D{{"metaData.clusterName", bson.D{{"$in", clusterArr}}}}

	cur, err := db.Select(filter)

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem model.Clusterprovisions
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		// var kubeConfig kubeapi.Config
		kubeConfigStr, _ := json.Marshal(elem.ProvisionResult.Kubeconfig)
		// json.Unmarshal(kubeConfigStr, &kubeConfig)

		kubeRestConfig, _ := clientcmd.RESTConfigFromKubeConfig(kubeConfigStr)
		resClient, errRestClient := rest.RESTClientFor(kubeRestConfig)
		if errRestClient != nil {
			log.Println(errRestClient)
		}

		resultStrArr, _ := resClient.Get().AbsPath("api/v1/nodes").DoRaw()

		aaa := model.ZCPResources{
			ClusterName: elem.MetaData.ClusterName,
			Values:      nil,
		}

		json.Unmarshal(resultStrArr, &aaa.Values)

		// dClient, errClient := dynamic.NewForConfig(kubeRestConfig)

		// if errClient != nil {
		// 	log.Fatal(errClient)
		// }

		// testGVR := schema.GroupVersionResource{
		// 	Group:    "",
		// 	Version:  "v1",
		// 	Resource: "nodes",
		// }

		// k8sClient := dClient.Resource(testGVR)

		// opt := metav1.ListOptions{}

		// resultList, errList := k8sClient.List(opt)
		// if errList != nil {
		// 	log.Println(errList)
		// }
		// log.Printf("Got CRD: %v", resultList)

		resultMapArr = append(resultMapArr, aaa)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	if err != nil {
		// log.Fatal(err)
		// return c.JSON(http.StatusInternalServerError, util.SetSuccessFalse(err.Error(), err))
	}

	return resultMapArr

}
