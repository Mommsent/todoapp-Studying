package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_errors "github.com/Mommsent/todoapp-Studying.git/internal/core/errors"
	"github.com/go-playground/validator/v10"
)

type validatable interface {
	Validate() error
}

var requestValidator = validator.New()

func DecodeAndValidateRequest(request *http.Request, dest any) error {
	if err := json.NewDecoder(request.Body).Decode(dest); err != nil {
		return fmt.Errorf("decode json: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	var (
		err error
	)

	validatable, ok := dest.(validatable)
	if ok {
		err = validatable.Validate()
	} else {
		err = requestValidator.Struct(dest)
	}

	if err != nil {
		return fmt.Errorf(
			"request validator: %v: %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
