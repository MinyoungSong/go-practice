package model

import (

)

type ResponseBody struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Errors  string      `json:"errors"`
	Data    interface{} `json:"data"`
}
