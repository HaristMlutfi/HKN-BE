package objects

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type BaseResponse struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

//func SetResponse(echo echo.Context, data interface{}, err error){
//	code := http.StatusOK
//	if err != nil {
//		code = http.StatusBadRequest
//	}
//	return echo.JSON(
//		code, BaseResponse{
//			Code:  code,
//			Data:  data,
//			Error: err,
//		},
//	)
//}

func SetResponse(c echo.Context, err error, result interface{}) error {
	code := http.StatusOK
	errString := ""
	if err != nil {
		code = http.StatusBadRequest
		errString = err.Error()
	}
	return c.JSON(
		code, BaseResponse{
			Code:  code,
			Data:  result,
			Error: errString,
		},
	)
}
