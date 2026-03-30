package model

type SubmissionParams struct {
	Auth  *AuthParams
	Log   *LogParams
	Mail  *MailParams
	Queue *QueueParams
}

type AuthParams struct {
	Email    string
	Password string
}

type LogParams struct {
	Name string
	Data string
}

type MailParams struct {
	From    string
	To      string
	Subject string
	Message string
}

type QueueParams struct {
	Key  string
	Data LogParams
}
