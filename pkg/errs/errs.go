package errs

import (
	"fmt"
	"runtime"
)

type Error struct {
	Err     error
	Kind    Kind
	Code    Code
	Level   int
	Message string
}

type Kind uint8

type Code string

func (e *Error) Error() string {
	return e.Err.Error()
}

const (
	Other Kind = iota
	Invalid
	Exist
	NotExist
	NotFound
	Internal
	Validation
	InvalidRequest

	Unauthenticated
	Unauthorized
)

func (k Kind) String() string {
	switch k {
	case Other:
		return "other_error"
	case Invalid:
		return "invalid_operation"
	case Exist:
		return "already_exists"
	case NotExist:
		return "not_exist"
	case NotFound:
		return "not_found"
	case Internal:
		return "internal"
	case Validation:
		return "input_validation"
	case InvalidRequest:
		return "invalid_request"
	case Unauthenticated:
		return "unauthenticated_request"
	case Unauthorized:
		return "unauthorized_request"
	default:
		return "unknown_error_kind"
	}
}

func New(args ...interface{}) error {
	if len(args) == 0 {
		panic("call to errs.New with no arguments")
	}

	var err = &Error{}
	err.Message = "no message"

	for _, arg := range args {
		switch arg := arg.(type) {
		case string:
			err.Message = arg
		case Kind:
			err.Kind = arg
		case error:
			err.Err = arg
		default:
			_, file, line, _ := runtime.Caller(1)
			return fmt.Errorf("errs.New: bad call from %s:%d: %v, unknown type %T, value %v in error call", file, line, args, arg, arg)
		}
	}

	return err
}
