package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Mommsent/todoapp-Studying.git/internal/core/domain"
	core_logger "github.com/Mommsent/todoapp-Studying.git/internal/core/logger"
	core_http_request "github.com/Mommsent/todoapp-Studying.git/internal/core/transport/http/request"
	core_http_response "github.com/Mommsent/todoapp-Studying.git/internal/core/transport/http/response"
	core_http_types "github.com/Mommsent/todoapp-Studying.git/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

type PatchUserResponse UserDTOResponse

func (userPatchRequest *PatchUserRequest) Validate() error {
	if userPatchRequest.FullName.Set {
		if userPatchRequest.FullName.Value == nil {
			return fmt.Errorf("'FullName' cant be NULL")
		}

		fullNameLen := len([]rune(*userPatchRequest.FullName.Value))
		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf("'FullName' must be between 3 and 100 symbols")
		}
	}

	if userPatchRequest.PhoneNumber.Set {
		if userPatchRequest.PhoneNumber.Value != nil {
			phoneNumberLen := len([]rune(*userPatchRequest.PhoneNumber.Value))
			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf("'PhoneNumber' must be between 10 and 15 symbols")
			}

			if !strings.HasPrefix(*userPatchRequest.PhoneNumber.Value, "+") {
				return fmt.Errorf("'PhoneNumber' must begin from '+' symbol")
			}
		}
	}

	return nil
}

func (handler *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)

		return
	}

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate http request",
		)

		return
	}

	userPatch := userPatchFromRequest(request)
	userDomain, err := handler.userService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)

		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		request.FullName.ToDomain(),
		request.PhoneNumber.ToDomain(),
	)
}
