package cluster

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"skcloud.io/cloudzcp/zcpctl-backend/db"
	"skcloud.io/cloudzcp/zcpctl-backend/model"
	"skcloud.io/cloudzcp/zcpctl-backend/util"
)

func GetClsuterList() echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		resultArr := []model.Clusterprovisions{}

		filter := bson.D{{}}

		if param := c.Param("cluster_name"); param != "" {
			filter = bson.D{{"metaData.clusterName", param}}
		}

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

			resultArr = append(resultArr, elem)
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		// Close the cursor once finished
		cur.Close(context.TODO())

		if err != nil {
			// log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, util.SetSuccessFalse(err.Error(), err))
		}

		return c.JSON(http.StatusOK, util.SetSuccessTrue(resultArr))

	}

}

func GetClsuterCredential() echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		clusterModel := model.Clusterprovisions{}
		var result map[string]interface{}

		filter := bson.D{{}}

		if param := c.Param("cluster_name"); param != "" {
			filter = bson.D{{"metaData.clusterName", param}}
		}

		resultSet := db.SelectOne(filter)

		if err1 := resultSet.Decode(&clusterModel); err1 != nil {
			// log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, util.SetSuccessFalse(err1.Error(), err1))
		}

		result = clusterModel.ProvisionResult.Kubeconfig

		return c.JSON(http.StatusOK, util.SetSuccessTrue(result))

	}

}

// func GetClsuterListMD() echo.HandlerFunc {

// 	return func(c echo.Context) (err error) {

// 		resultArr := []model.Clusterprovisions{}
// 		// obj := &model.Clusterprovisions{}

// 		filter := bson.D{{}}

// 		if param := c.Param("cluster_name"); param != "" {
// 			filter = bson.D{{"metaData.clusterName", param}}
// 		}

// 		resultArr = db.Select(filter)

// 		if err != nil {
// 			// log.Fatal(err)
// 			return c.JSON(http.StatusInternalServerError, util.SetSuccessFalse(err.Error(), err))
// 		}
// 		// db.Select(collenction, obj, filter)
// 		// db.SelectGen(&reaultArr, paramModel, filter)

// 		return c.JSON(http.StatusOK, util.SetSuccessTrue(resultArr))

// 	}

// }
