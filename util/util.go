package util

import (
	"skcloud.io/cloudzcp/zcpctl-backend/model"
)

func IsEmptyString(value string) bool {

	if value != "" {
		return false
	}

	return true
}

func SetSuccessTrue(data interface{}) *model.ResponseBody {

	result := new(model.ResponseBody)
	result.Success = true
	result.Message = ""
	result.Errors = ""
	result.Data = data

	return result

}
