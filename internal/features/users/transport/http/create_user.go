package users_transport_http

import (
	"net/http"

	"github.com/Mommsent/todoapp-Studying.git/internal/core/domain"
	core_logger "github.com/Mommsent/todoapp-Studying.git/internal/core/logger"
	core_http_request "github.com/Mommsent/todoapp-Studying.git/internal/core/transport/http/request"
	core_http_response "github.com/Mommsent/todoapp-Studying.git/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

type CreateUserResponse UserDTOResponse

func (handler *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	log.Debug("invoce CreateUser handler")

	var userRequest CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(request, &userRequest); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userDomain := domainFromDTO(userRequest)
	userDomain, err := handler.userService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to create user")
		return
	}

	response := CreateUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}
