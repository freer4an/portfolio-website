package util

import (
	"errors"
	"strconv"
)

var (
	ErrParam = errors.New("invalid param")
)

func UrlParamToInt(paramStr string) (uint, error) {
	param, err := strconv.Atoi(paramStr)

	if err != nil || param < 1 {
		return 0, ErrParam
	}

	return uint(param), err
}
