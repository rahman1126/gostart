package common

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ResponseBody struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(w http.ResponseWriter, data interface{}, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	code := strconv.Itoa(GetStatusCode(err))

	res := &ResponseBody{
		Code:    code,
		Message: GetErrorMessage(err),
		Data:    data,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	// TODO : list more error status
	switch err {
	default:
		return http.StatusNotFound
	}
}

func GetErrorMessage(err error) string {
	if err == nil {
		return "Success"
	} else {
		return err.Error()
	}
}