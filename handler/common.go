package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

var appEnv string

func init() {
	appEnv = os.Getenv("APP_ENV")
}

func queryParamInt(r *http.Request, key string, defaultValue int64) (int64, error) {
	v := r.FormValue(key)

	if v == "" {
		return defaultValue, nil
	}
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil && appEnv == "prod" {
		return i, errors.New("Parameter " + key + " must be an integer")
	}
	return strconv.ParseInt(v, 10, 64)
}
