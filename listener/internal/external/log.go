package external

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	contreq "micro/common/contracts/http/req"
	contres "micro/common/contracts/http/res"
	"net/http"
	"time"
)

type LogService struct {
	url    string
	client http.Client
}

func NewLogService(url string) *LogService {
	return new(LogService{
		client: http.Client{
			Timeout: time.Minute,
		},
		url: url,
	})
}

func (s *LogService) Log(ctx context.Context, params contreq.WriteLog) (data any, err error) {
	authBody, err := json.Marshal(params)
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

	var v contres.WriteLog
	err = body.Decode(&v)
	if err != nil {
		return nil, fmt.Errorf("invalid response format: %w", err)
	}

	return v.Message, nil
}
