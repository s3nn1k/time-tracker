package handler

import (
	"net/http"

	"github.com/Time-Tracker/internal/auth"
	"github.com/Time-Tracker/internal/service"
)

type Handler struct {
	Service *service.Service
	Auth    auth.Auth
}

func New(s *service.Service, a auth.Auth) *Handler {
	return &Handler{
		Service: s,
		Auth:    a,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /user", h.signIn)
	router.HandleFunc("PUT /user", h.signUp)
	router.Handle("POST /user/name/{name}", h.Authorize(h.UpdateUserName))
	router.Handle("POST /user/password/{password}", h.Authorize(h.UpdateUserPassword))
	router.Handle("DELETE /user", h.Authorize(h.DeleteUser))

	router.Handle("POST /timers/create/{name}", h.Authorize(h.CreateTimer))
	router.Handle("POST /timers/start/{id}", h.Authorize(h.StartTimer))
	router.Handle("POST /timers/stop/{id}", h.Authorize(h.StopTimer))
	router.Handle("POST /timers/{id}/{name}", h.Authorize(h.UpdateTimerName))
	router.Handle("GET /timers", h.Authorize(h.GetTimers))
	router.Handle("GET /timers/{id}", h.Authorize(h.GetTimer))
	router.Handle("DELETE /timers/{id}", h.Authorize(h.DeleteTimer))

	return router
}
