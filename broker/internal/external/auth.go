package external

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"micro/broker/internal/model"
	contres "micro/common/contracts/http/res"
	"net/http"
	"time"
)

type AuthService struct {
	url    string
	client http.Client
}

func NewAuthService(url string) *AuthService {
	return new(AuthService{
		client: http.Client{
			Timeout: time.Minute,
		},
		url: url,
	})
}

func (s *AuthService) Authenticate(ctx context.Context, params model.AuthParams) (data any, err error) {
	authBody, err := json.Marshal(AuthToReq(params))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, s.url, bytes.NewReader(authBody))
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body := json.NewDecoder(res.Body)

	if res.StatusCode != http.StatusAccepted {
		var v contres.Error
		err := body.Decode(&v)
		if err != nil {
			return nil, fmt.Errorf("invalid response format: %w", err)
		}

		return nil, errors.New(v.Message)
	}

	var v contres.Authenticate
	err = body.Decode(&v)
	if err != nil {
		return nil, fmt.Errorf("invalid response format: %w", err)
	}

	return v.Data, nil
}
