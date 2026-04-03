package model

type Action int

const (
	_ Action = iota
	AuthAction
	LogAction
	GRPCLogAction
	MailAction
)
