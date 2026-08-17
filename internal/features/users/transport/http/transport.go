package users_transport_http

import (
	"net/http"
	corre_http_server "todolist/internal/core/transport/http/server"
)

type UserHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
}

func NewUsersHTTPHandler(usersService UsersService) *UserHTTPHandler {
	return &UserHTTPHandler{usersService: usersService}
}

func (h *UserHTTPHandler) Routes() []corre_http_server.Route {
	return []corre_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}
