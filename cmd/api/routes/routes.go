package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"uptime-api/m/v2/cmd/api/deps"
	"uptime-api/m/v2/cmd/api/handler"
	"uptime-api/m/v2/internal/core/usecase"
)

func Routes(deps *deps.Dependencies) http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           3600,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	getMonitorById := usecase.NewGetMonitorByIdUseCase(deps.MonitorRepository, deps.NotificationRepository)
	getAllMonitorByIds := usecase.NewGetAllMonitorsByIdsUseCase(deps.MonitorRepository)
	getMonitorRunsTimeSeries := usecase.NewGetMonitorRunsTimeSeriesUseCase(deps.MonitorRunRepository)
	assignMonitorToNotification := usecase.NewAssignMonitorToNotificationUseCase(getMonitorById, deps.NotificationRepository)
	getStatusPageById := usecase.NewGetStatusPageByIdUseCase(
		deps.StatusPageRepository,
		deps.MonitorRepository)
	getStatusPageViewData := usecase.NewGetStatusPageViewDataUseCase(getStatusPageById, getMonitorRunsTimeSeries)
	assignMonitorToStatusPage := usecase.NewAssignMonitorToStatusPageUseCase(getAllMonitorByIds, getStatusPageById, deps.StatusPageRepository)
	unassignMonitorFromStatusPage := usecase.NewUnassignMonitorFromStatusPageUseCase(getMonitorById, getStatusPageById, deps.StatusPageRepository)

	mux.Post("/monitors", handler.NewCreateMonitorHandler(usecase.NewCreateMonitorUseCase(deps.MonitorRepository)))
	mux.Get("/monitors/{id}", handler.NewGetMonitorByIdHandler(getMonitorById))
	mux.Put("/monitors/{id}", handler.NewUpdateMonitorHandler(usecase.NewUpdateMonitorUseCase(getMonitorById, deps.MonitorRepository)))
	mux.Delete("/monitors/{id}", handler.NewDeleteMonitorHandler(usecase.NewDeleteMonitorUseCase(getMonitorById, deps.MonitorRepository)))

	mux.Post("/notifications", handler.NewCreateNotificationHandler(usecase.NewCreateNotificationUseCase(
		deps.NotificationRepository,
		assignMonitorToNotification,
	)))
	mux.Get("/notifications/{id}", handler.NewGetNotificationByIdHandler(usecase.NewGetNotificationByIdUseCase(deps.NotificationRepository)))
	mux.Post("/notifications/{id}/assignments", handler.NewAssignMonitorToNotificationHandler(assignMonitorToNotification))
	mux.Delete("/notifications/{id}/assignments/{monitor_id}", handler.NewUnassignMonitorFromNotificationHandler(
		usecase.NewUnassignMonitorFromNotificationUseCase(deps.NotificationRepository)))

	mux.Post("/pages", handler.NewCreateStatusPageHandler(usecase.NewCreateStatusPageUseCase(
		deps.StatusPageRepository,
		deps.MonitorRepository)))
	mux.Get("/pages/{id}", handler.NewGetPageByIdHandler(getStatusPageById))
	mux.Get("/pages/{id}/view", handler.NewGetPageViewByIdHandler(getStatusPageViewData))
	mux.Post("/pages/{id}/assignments", handler.NewAssignMonitorToStatusPageHandler(assignMonitorToStatusPage))
	mux.Delete("/pages/{id}/assignments/{monitor_id}", handler.NewUnassignMonitorFromStatusPageHandler(unassignMonitorFromStatusPage))

	return mux
}
