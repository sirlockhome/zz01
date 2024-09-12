package req

import (
	"foxomni/pkg/errs"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetIntFromVars(r *http.Request, name string) (int, error) {
	n, err := strconv.Atoi(mux.Vars(r)[name])
	if err != nil {
		return -1, errs.New(errs.InvalidRequest, "not number")
	}

	return n, nil
}
