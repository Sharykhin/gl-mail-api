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
	n, err := strconv.Atoi(v)
	return int64(n), err
}
