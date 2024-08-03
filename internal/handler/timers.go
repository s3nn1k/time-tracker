package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Time-Tracker/internal"
)

func (h *Handler) CreateTimer(w http.ResponseWriter, r *http.Request, id int) {
	name := r.PathValue("name")

	if name == "" {
		http.Error(w, "wrong name field", http.StatusBadRequest)
		return
	}

	timerId, err := h.Service.Timers.Create(r.Context(), id, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := Response{TimerId: timerId}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) StartTimer(w http.ResponseWriter, r *http.Request, id int) {
	strTimerId := r.PathValue("id")

	timerId, err := strconv.Atoi(strTimerId)
	if err != nil {
		http.Error(w, "wrong id value", http.StatusBadRequest)
	}

	err = h.Service.Timers.Start(r.Context(), timerId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) StopTimer(w http.ResponseWriter, r *http.Request, id int) {
	strTimerId := r.PathValue("id")

	timerId, err := strconv.Atoi(strTimerId)
	if err != nil {
		http.Error(w, "wrong id value", http.StatusBadRequest)
	}

	err = h.Service.Timers.Stop(r.Context(), timerId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateTimerName(w http.ResponseWriter, r *http.Request, id int) {
	name := r.PathValue("name")
	strTimerId := r.PathValue("id")

	timerId, err := strconv.Atoi(strTimerId)
	if err != nil {
		http.Error(w, "wrong id value", http.StatusBadRequest)
	}

	if name == "" {
		http.Error(w, "empty name field", http.StatusBadRequest)
		return
	}

	err = h.Service.Timers.Update(r.Context(), timerId, id, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetTimers(w http.ResponseWriter, r *http.Request, id int) {
	timers, err := h.Service.Timers.GetByUserId(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cleanTimers := convertTimers(timers)

	resp := Response{Timers: cleanTimers}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetTimer(w http.ResponseWriter, r *http.Request, id int) {
	strTimerId := r.PathValue("id")

	timerId, err := strconv.Atoi(strTimerId)
	if err != nil {
		http.Error(w, "wrong id value", http.StatusBadRequest)
	}

	name, startTime, workTime, err := h.Service.Timers.GetById(r.Context(), timerId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	timer := internal.Timer{
		Id:        timerId,
		Name:      name,
		LastStart: startTime,
		WorkTime:  workTime,
	}

	timers := convertTimers([]internal.Timer{timer})

	resp := Response{Timers: timers}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteTimer(w http.ResponseWriter, r *http.Request, id int) {
	strTimerId := r.PathValue("id")

	timerId, err := strconv.Atoi(strTimerId)
	if err != nil {
		http.Error(w, "wrong id value", http.StatusBadRequest)
	}

	err = h.Service.Timers.Delete(r.Context(), timerId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
