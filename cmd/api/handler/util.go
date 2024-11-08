package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type JsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ReadJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	maxBytes := int64(1048576) // 1 MB

	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(v)

	if err != nil {
		return err
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}, headers ...http.Header) error {
	out, err := json.Marshal(v)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)

	if err != nil {
		return err
	}

	return nil
}

func ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	response := &JsonResponse{
		Error:   true,
		Message: err.Error(),
	}

	return WriteJSON(w, statusCode, response)
}

func GetID(r *http.Request) (int64, error) {
	id := r.PathValue("id")

	if id == "" {
		return 0, ErrIDRequired
	}

	return strconv.ParseInt(id, 10, 64)
}

func ErrParamRequired(param string) error {
	return errors.New(fmt.Sprintf("param %s is required", param))
}

func GetParam(r *http.Request, param string) (string, error) {
	value := r.PathValue(param)

	if value == "" {
		return "", ErrParamRequired(param)
	}

	return value, nil
}

func GetIntParam(r *http.Request, param string) (int64, error) {
	value, err := GetParam(r, param)

	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(value, 10, 64)
}
