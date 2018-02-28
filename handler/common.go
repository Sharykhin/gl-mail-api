package handler

import (
	"net/http"
	"strconv"
)

func queryParamInt(r *http.Request, key string, defaultValue int64) (int64, error) {
	v := r.FormValue(key)

	if v == "" {
		return defaultValue, nil
	}
	return strconv.ParseInt(v, 10, 64)
}
