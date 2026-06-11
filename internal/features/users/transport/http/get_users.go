package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Mommsent/todoapp-Studying.git/internal/core/logger"
	core_http_request "github.com/Mommsent/todoapp-Studying.git/internal/core/transport/http/request"
	core_http_response "github.com/Mommsent/todoapp-Studying.git/internal/core/transport/http/response"
)

type GetUsersResponse []UserDTOResponse

// GetUsers 	godoc
// @Summary 	Список пользователей
// @Description Просмотр списка пользователей с опциональной пагинацией
// @Tags 		users
// @Produce 	json
// @Param       limit query int false "Размер страницы с пользователями"
// @Param       offset quey int false "Смещение страницы с ползователями"
// @Success 	200 {object} GetUsersResponse "Успешно вернул пользователя"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failuer     404 {object} core_http_response.ErrorResponse "User not found"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/users [get]
func (handler *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(request)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'limit'/'offset' query param",
		)

		return
	}

	userDomains, err := handler.userService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")
		return
	}

	response := GetUsersResponse(usersDTOFromDomains(userDomains))
	responseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(request *http.Request) (*int, *int, error) {

	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)

	limit, err := core_http_request.GetIntQueryParam(request, limitQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := core_http_request.GetIntQueryParam(request, offsetQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}

	return limit, offset, nil
}
