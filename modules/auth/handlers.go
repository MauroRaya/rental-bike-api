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

type SignInResponse struct {
	Token string `json:"token"`
}

// SignUp godoc
// @Summary Sign up
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body auth.SignUpParams true "Sign up payload"
// @Success 201 {object} repo.CreateUserRow
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Router /auth/sign-up [post]
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

// SignIn godoc
// @Summary Sign in
// @Description Authenticate user
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body auth.SignInParams true "Sign in payload"
// @Success 200 {object} auth.SignInResponse
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Router /auth/sign-in [post]
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

	data := SignInResponse{
		Token: token,
	}

	httputil.EncodeJSON(w, http.StatusOK, data)
}
