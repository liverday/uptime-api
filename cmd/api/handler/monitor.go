package handler

import (
	"errors"
	"net/http"
	"uptime-api/m/v2/internal/core/inputs"
	"uptime-api/m/v2/internal/core/usecase"
)

type CreateMonitorRequest struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Periodicity   string `json:"periodicity"`
	Url           string `json:"url"`
	Method        string `json:"method"`
	Headers       string `json:"headers"`
	Body          string `json:"body"`
	DegradedAfter int64  `json:"degraded_after"`
}

type UpdateMonitorRequest struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Periodicity   string `json:"periodicity"`
	Url           string `json:"url"`
	Method        string `json:"method"`
	Headers       string `json:"headers"`
	Body          string `json:"body"`
	DegradedAfter int64  `json:"degraded_after"`
}

var (
	ErrIDRequired = errors.New("id is required")
)

func NewCreateMonitorHandler(useCase *usecase.CreateMonitorUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateMonitorRequest

		err := ReadJSON(w, r, &req)

		if err != nil {
			ErrorJSON(w, err)
			return
		}

		m, err := useCase.Execute(r.Context(), &inputs.CreateMonitorInput{
			Title:         req.Title,
			Description:   req.Description,
			Periodicity:   req.Periodicity,
			Url:           req.Url,
			Method:        req.Method,
			Headers:       req.Headers,
			Body:          req.Body,
			DegradedAfter: req.DegradedAfter,
		})

		if err != nil {
			ErrorJSON(w, err)
			return
		}

		err = WriteJSON(w, http.StatusCreated, m)

		if err != nil {
			ErrorJSON(w, err)
		}
	}
}

func NewUpdateMonitorHandler(useCase *usecase.UpdateMonitorUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := GetID(r)

		if err != nil {
			ErrorJSON(w, err)
			return
		}

		var req UpdateMonitorRequest

		err = ReadJSON(w, r, &req)

		if err != nil {
			ErrorJSON(w, err)
			return
		}

		m, err := useCase.Execute(r.Context(), &inputs.UpdateMonitorInput{
			MonitorId:     id,
			Title:         req.Title,
			Description:   req.Description,
			Periodicity:   req.Periodicity,
			Url:           req.Url,
			Method:        req.Method,
			Headers:       req.Headers,
			Body:          req.Body,
			DegradedAfter: req.DegradedAfter,
		})

		if err != nil {
			ErrorJSON(w, err)
			return
		}

		err = WriteJSON(w, http.StatusOK, m)

		if err != nil {
			ErrorJSON(w, err)
		}
	}
}

func NewGetMonitorByIdHandler(useCase *usecase.GetMonitorByIdUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := GetID(r)

		if err != nil {
			ErrorJSON(w, err)
			return
		}

		m, err := useCase.Execute(r.Context(), id)

		if err != nil {
			ErrorJSON(w, err)
			return
		}

		err = WriteJSON(w, http.StatusOK, m)

		if err != nil {
			ErrorJSON(w, err)
		}
	}
}

func NewDeleteMonitorHandler(useCase *usecase.DeleteMonitorUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := GetID(r)

		if err != nil {
			ErrorJSON(w, err)
			return
		}

		err = useCase.Execute(r.Context(), id)

		if err != nil {
			ErrorJSON(w, err, http.StatusNotFound)
			return
		}

		err = WriteJSON(w, http.StatusNoContent, nil)

		if err != nil {
			ErrorJSON(w, err)
		}
	}
}
