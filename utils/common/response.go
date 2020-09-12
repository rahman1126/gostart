package common

import (
	"encoding/json"
	"gostart/dto"
	"net/http"
	"strconv"
)

func Response(w http.ResponseWriter, r *http.Request, data interface{}, err error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(GetStatusCode(err))

	if r.Method == http.MethodOptions {
		return
	}

	code := strconv.Itoa(GetStatusCode(err))

	res := &dto.ResponseBody{
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
	case ErrUnauthorized:
		return http.StatusUnauthorized
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