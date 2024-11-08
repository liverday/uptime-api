package handler

import (
	"net/http"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/inputs"
	"uptime-api/m/v2/internal/core/usecase"
)

type CreateStatusPageRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	MonitorIds  []int64 `json:"monitor_ids"`
}

type AssignMonitorToStatusPageRequest struct {
	MonitorIds []int64 `json:"monitor_ids"`
}

func NewCreateStatusPageHandler(useCase *usecase.CreateStatusPageUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateStatusPageRequest

		err := ReadJSON(w, r, &req)

		if err != nil {
			_ = ErrorJSON(w, err)
			return
		}

		sp, err := useCase.Execute(r.Context(), &inputs.CreateStatusPageInput{
			Title:       req.Title,
			Description: req.Description,
			MonitorIds:  req.MonitorIds,
		})

		if err != nil {
			_ = ErrorJSON(w, err)
			return
		}

		_ = WriteJSON(w, http.StatusCreated, sp)
	}
}

type MonitorUptime struct {
	Monitor *domain.Monitor
	Uptime  float64
	Entries []*domain.UptimeEntry
}

type Response struct {
	Title       string
	Description string
	Monitors    []*MonitorUptime
}

func NewGetPageByIdHandler(getStatusPageById *usecase.GetStatusPageByIdUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := GetID(r)

		if err != nil {
			_ = ErrorJSON(w, err)
		}

		sp, err := getStatusPageById.Execute(r.Context(), id)

		if err != nil {
			_ = ErrorJSON(w, err)
		}

		_ = WriteJSON(w, http.StatusOK, sp)
	}
}

func NewGetPageViewByIdHandler(
	useCase *usecase.GetStatusPageViewDataUseCase,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := GetID(r)

		if err != nil {
			render(w, "notfound.page.gohtml", nil)
			return
		}

		data, err := useCase.Execute(r.Context(), id)

		if err != nil {
			render(w, "notfound.page.gohtml", nil)
			return
		}

		render(w, "status.page.gohtml", data)

		if err != nil {
			render(w, "notfound.page.gohtml", nil)
		}
	}
}

func NewAssignMonitorToStatusPageHandler(useCase *usecase.AssignMonitorToStatusPageUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notificationId, err := GetID(r)

		if err != nil {
			_ = ErrorJSON(w, err)
			return
		}

		var req AssignMonitorToStatusPageRequest

		err = ReadJSON(w, r, &req)

		if err != nil {
			_ = ErrorJSON(w, err)
			return
		}

		err = useCase.Execute(r.Context(), &inputs.AssignMonitorToStatusPageInput{
			StatusPageId: notificationId,
			MonitorIds:   req.MonitorIds,
		})

		if err != nil {
			_ = ErrorJSON(w, err)
			return
		}

		_ = WriteJSON(w, http.StatusNoContent, nil)
	}
}

func NewUnassignMonitorFromStatusPageHandler(useCase *usecase.UnassignMonitorFromStatusPageUseCase) http.HandlerFunc {
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

		err = useCase.Execute(r.Context(), &inputs.UnassignMonitorFromStatusPageInput{
			StatusPageId: notificationId,
			MonitorId:    monitorId,
		})

		if err != nil {
			_ = ErrorJSON(w, err)
			return
		}

		_ = WriteJSON(w, http.StatusNoContent, nil)
	}
}
