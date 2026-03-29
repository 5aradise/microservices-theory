package external

import (
	"micro/authentication/internal/model"
	contreq "micro/common/contracts/http/req"
)

func LogToReq(auth model.LogParams) contreq.WriteLog {
	return contreq.WriteLog{
		Name: auth.Name,
		Data: auth.Data,
	}
}
