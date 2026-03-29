package model

type SubmissionParams struct {
	Auth *AuthParams
	Log  *LogParams
	Mail *MailParams
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
