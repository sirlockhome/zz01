package errs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type ErrorResp struct {
	Error ServiceError `json:"error"`
}

type ServiceError struct {
	Kind    string `json:"kind,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func HTTPErrorResponse(ctx context.Context, w http.ResponseWriter, err error) {
	if err == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var e *Error
	if errors.As(err, &e) {
		switch e.Kind {
		case Unauthenticated:
			return
		default:
			typcialErrorResponse(ctx, w, e)
			return
		}
	}

	unknownErrorResponse(ctx, w, err)
}

func typcialErrorResponse(ctx context.Context, w http.ResponseWriter, e *Error) {
	const errMsg = "error response sent to client"
	httpStatusCode := httpErrorStatusCode(e.Kind)

	log.Ctx(ctx).Err(e.Err).Str("kind", e.Kind.String()).Int("status_code", httpStatusCode).Msg(errMsg)

	errResp := newErrResponse(e)

	errJSON, _ := json.Marshal(errResp)
	ej := string(errJSON)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	// Write HTTP Statuscode
	w.WriteHeader(httpStatusCode)

	// Write response body (json)
	fmt.Fprintln(w, ej)
}

func httpErrorStatusCode(k Kind) int {
	switch k {
	case Invalid, Exist, NotExist, Validation, InvalidRequest:
		return http.StatusBadRequest
	case NotFound:
		return http.StatusNotFound
	case Other, Internal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func unknownErrorResponse(ctx context.Context, w http.ResponseWriter, err error) {
	er := ErrorResp{
		Error: ServiceError{
			Kind:    "unknown",
			Message: "unexpected error - contact support",
		},
	}

	log.Ctx(ctx).Error().Err(err).Msg("unknown error")

	// Marshal errResponse struct to JSON for the response body
	errJSON, _ := json.Marshal(er)
	ej := string(errJSON)

	// Write Content-Type headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	// Write HTTP Statuscode
	w.WriteHeader(http.StatusInternalServerError)

	// Write response body (json)
	fmt.Fprintln(w, ej)
}

func newErrResponse(e *Error) ErrorResp {
	const msg string = "internal server error - please contact support"

	switch e.Kind {
	case Internal:
		return ErrorResp{
			Error: ServiceError{
				Kind:    Internal.String(),
				Message: msg,
			},
		}
	default:
		return ErrorResp{
			Error: ServiceError{
				Kind:    e.Kind.String(),
				Message: e.Message,
			},
		}
	}
}
