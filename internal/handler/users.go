package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Time-Tracker/internal"
	"github.com/Time-Tracker/internal/storage"
)

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	user := internal.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Password == "" || user.Name == "" {
		http.Error(w, "empty name or password fields", http.StatusBadRequest)
		return
	}

	passwordHash := h.Auth.GeneratePasswordHash(user.Password)

	id, err := h.Service.Users.Create(r.Context(), user.Name, passwordHash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := h.Auth.GenerateToken(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := Response{Token: token}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	user := internal.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Password == "" || user.Name == "" {
		http.Error(w, "empty name or password fields", http.StatusBadRequest)
		return
	}

	passwordHash := h.Auth.GeneratePasswordHash(user.Password)

	id, err := h.Service.Users.GetId(r.Context(), user.Name, passwordHash)
	if err != nil {
		if err != storage.ErrNotExist {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Error(w, "user not exist", http.StatusBadRequest)
		return
	}

	token, err := h.Auth.GenerateToken(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := Response{Token: token}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) UpdateUserName(w http.ResponseWriter, r *http.Request, id int) {
	name := r.PathValue("name")

	if name == "" {
		http.Error(w, "empty name field", http.StatusBadRequest)
		return
	}

	err := h.Service.Users.Update(r.Context(), id, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateUserPassword(w http.ResponseWriter, r *http.Request, id int) {
	password := r.PathValue("password")

	if password == "" {
		http.Error(w, "empty password field", http.StatusBadRequest)
		return
	}

	err := h.Service.Users.UpdatePassword(r.Context(), id, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	err := h.Service.Users.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
