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

func SetSuccessFalse(msg string, err interface{}) *model.ResponseBody {

	result := new(model.ResponseBody)
	result.Success = false
	result.Message = msg
	result.Errors = err
	result.Data = new(interface{})

	return result

}
