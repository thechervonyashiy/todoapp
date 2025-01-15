package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thechervonyashiy/todoapp/pkg/dto"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	var req dto.SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid input: %s", err), http.StatusBadRequest)
		return
	}

	res, err := h.services.Authorization.CreateUser(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Failed to encode responce", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {

}
