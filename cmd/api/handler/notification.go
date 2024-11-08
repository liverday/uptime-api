package handler

import (
	"errors"
	"net/http"
	"uptime-api/m/v2/internal/core/inputs"
	"uptime-api/m/v2/internal/core/usecase"
)

type CreateNotificationRequest struct {
	Title        string                 `json:"title"`
	Description  string                 `json:"description"`
	Provider     string                 `json:"provider"`
	ProviderData map[string]interface{} `json:"provider_data"`
	MonitorIds   []int64                `json:"monitor_ids"`
}

type AssignMonitorToNotificationRequest struct {
	MonitorId int64 `json:"monitor_id"`
}

var ErrMonitorIdIsRequired = errors.New("monitor_id is required")

func NewCreateNotificationHandler(useCase *usecase.CreateNotificationUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateNotificationRequest

		err := ReadJSON(w, r, &req)

		if err != nil {
			ErrorJSON(w, err)
			return
		}

		n, err := useCase.Execute(r.Context(), &inputs.CreateNotificationInput{
			Title:        req.Title,
			Description:  req.Description,
			Provider:     req.Provider,
			ProviderData: req.ProviderData,
			MonitorIds:   req.MonitorIds,
		})

		if err != nil {
			ErrorJSON(w, err)
			return
		}

		WriteJSON(w, http.StatusCreated, n)

		if err != nil {
			ErrorJSON(w, err)
		}
	}
}

func NewGetNotificationByIdHandler(useCase *usecase.GetNotificationByIdUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := GetID(r)

		if err != nil {
			_ = ErrorJSON(w, err)
			return
		}

		n, err := useCase.Execute(r.Context(), id)

		if err != nil {
			_ = ErrorJSON(w, err, http.StatusNotFound)
			return
		}

		_ = WriteJSON(w, http.StatusOK, n)
	}
}

func NewAssignMonitorToNotificationHandler(useCase *usecase.AssignMonitorToNotificationUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notificationId, err := GetID(r)

		if err != nil {
			ErrorJSON(w, err)
			return
		}

		var req AssignMonitorToNotificationRequest

		err = ReadJSON(w, r, &req)

		if req.MonitorId == 0 {
			_ = ErrorJSON(w, ErrMonitorIdIsRequired, http.StatusBadRequest)
			return
		}

		if err != nil {
			_ = ErrorJSON(w, err)
			return
		}

		err = useCase.Execute(r.Context(), &inputs.AssignMonitorToNotificationInput{
			NotificationId: notificationId,
			MonitorId:      req.MonitorId,
		})

		if err != nil {
			_ = ErrorJSON(w, err)
			return
		}

		WriteJSON(w, http.StatusNoContent, nil)
	}
}

func NewUnassignMonitorFromNotificationHandler(useCase *usecase.UnassignMonitorFromNotificationUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notificationId, err := GetID(r)

		if err != nil {
			_ = ErrorJSON(w, err)
			return
		}

		monitorId, err := GetIntParam(r, "monitor_id")

		if err != nil {
			_ = ErrorJSON(w, err)
			return
		}

		err = useCase.Execute(r.Context(), &inputs.UnassignMonitorFromNotificationInput{
			NotificationId: notificationId,
			MonitorId:      monitorId,
		})

		if err != nil {
			_ = ErrorJSON(w, err)
			return
		}

		_ = WriteJSON(w, http.StatusNoContent, nil)
	}
}
