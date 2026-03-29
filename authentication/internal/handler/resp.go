package handler

import (
	"micro/authentication/internal/model"
	contres "micro/common/contracts/http/res"
)

func UserToResp(u model.User) contres.Authenticate {
	return contres.Authenticate{
		Data: contres.User{
			ID:        u.ID,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
		},
	}
}
