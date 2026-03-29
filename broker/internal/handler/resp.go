package handler

type SubmissionResp struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"` // map[string]any or struct{...}
}
