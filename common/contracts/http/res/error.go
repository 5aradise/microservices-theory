package contres

import (
	util "micro/common/utils"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"` // map[string]any or struct{...}
}

func WriteError(w http.ResponseWriter, err error, status ...int) error {
	code := http.StatusBadRequest
	if len(status) > 0 {
		code = status[0]
	}
	return util.WriteJSON(w, code, Error{
		Message: err.Error(),
	})
}
