package handlers

import (
	"backend/internal/services"
	"net/http"
)

type UserHandler struct {
	UserServices *services.UserServices
}

func (h *UserHandler) HandleCurrentGetUser(res http.ResponseWriter, req *http.Request) {
	// Get the user from the context
	// user, ok := req.Context().Value("user").(utils.JWTUser)
	// if !ok {
	// 	http.Error(res, "invalid user", http.StatusUnauthorized)
	// 	return
	// }

	// Get the user from the database
	// user, err := h.UserServices.GetUserByID(req.Context(), user.UserID)
	// if err != nil {
	// 	http.Error(res, err.Error(), http.StatusInternalServerError)
	// 	return
}

func (h *UserHandler) HandleCreateUser(res http.ResponseWriter, req *http.Request) {

	return
}

func (h *UserHandler) HandleUpdateUser(res http.ResponseWriter, req *http.Request) {
	return
}

func (h *UserHandler) HandleDeleteUser(res http.ResponseWriter, req *http.Request) {
	return
}
