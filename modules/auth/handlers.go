package auth

import (
	"net/http"

	"github.com/MauroRaya/bike-rental-api/httputil"
)

type handler struct {
	service Service
}

type Handler interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
}

func NewHandler(service Service) Handler {
	return &handler{service}
}

type SignUpParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload SignUpParams

	if err := httputil.DecodeJSON(r, &payload); err != nil {
		http.Error(w, "invalid request", http.StatusUnprocessableEntity)
		return
	}

	user, err := h.service.SignUp(ctx, payload.Email, payload.Password)
	if err != nil {
		http.Error(w, "unexpected error", http.StatusInternalServerError)
		return
	}

	httputil.EncodeJSON(w, http.StatusCreated, user)
}

func (h *handler) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload SignInParams

	if err := httputil.DecodeJSON(r, &payload); err != nil {
		http.Error(w, "invalid request", http.StatusUnprocessableEntity)
		return
	}

	token, err := h.service.SignIn(ctx, payload.Email, payload.Password)
	if err != nil {
		http.Error(w, "unexpected error", http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"token": token,
	}

	httputil.EncodeJSON(w, http.StatusOK, data)
}
