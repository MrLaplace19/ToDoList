package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	core_errors "github.com/MrLaplace19/ToDoList/internal/core/errors"
)

func GetIntQueryParams(r *http.Request, key string) (*int, error) {
	params := r.URL.Query().Get(key)
	if params == "" {
		return nil, nil
	}

	val, err := strconv.Atoi(params)
	if err != nil {
		return nil, fmt.Errorf(
			"params='%s' by key='%s' not a valid integer: %v: %w ",
			params,
			key,
			err,
			core_errors.ErrInvalidArgument,
		)
	}
	return &val, nil
}

func GetDateQueryParams(r *http.Request, key string) (*time.Time, error) {
	param := r.URL.Query().Get(key)
	if param == ""{
		return nil,nil
	}

	layout := "2006-01-02"
	qTime, err := time.Parse(layout, param)
	if err!=nil{
		return nil,fmt.Errorf(
			"param=%s by key=%s failed: %w",
			param,
			key,
			err,
		)
	}
	return &qTime, nil
}
