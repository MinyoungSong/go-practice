package cluster

import (
	"net/http"

	"github.com/Kamva/mgm"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"skcloud.io/cloudzcp/zcpctl-backend/db"
	"skcloud.io/cloudzcp/zcpctl-backend/model"
	"skcloud.io/cloudzcp/zcpctl-backend/util"
)

//Mogodb ODM 사용 - Kamva
func GetClsuterList() echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		resultArr := []model.Clusterprovisions{}
		// obj := &model.Clusterprovisions{}

		collection := mgm.Coll(&model.Clusterprovisions{})

		filter := bson.D{{}}

		if param := c.Param("cluster_name"); param != "" {
			filter = bson.D{{"metaData.clusterName", param}}
		}

		err = collection.SimpleFind(&resultArr, filter)

		if err != nil {
			// log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, util.SetSuccessFalse(err.Error(), err))
		}
		// db.Select(collenction, obj, filter)
		// db.SelectGen(&reaultArr, paramModel, filter)

		return c.JSON(http.StatusOK, util.SetSuccessTrue(resultArr))

	}

}

func GetClsuterListMD() echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		resultArr := []model.Clusterprovisions{}
		// obj := &model.Clusterprovisions{}

		filter := bson.D{{}}

		if param := c.Param("cluster_name"); param != "" {
			filter = bson.D{{"metaData.clusterName", param}}
		}

		resultArr = db.Select(filter)

		if err != nil {
			// log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, util.SetSuccessFalse(err.Error(), err))
		}
		// db.Select(collenction, obj, filter)
		// db.SelectGen(&reaultArr, paramModel, filter)

		return c.JSON(http.StatusOK, util.SetSuccessTrue(resultArr))

	}

}
