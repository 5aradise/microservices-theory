package model

type Mail struct {
	Domain     string
	Host       string
	Port       int
	Username   string
	Password   string
	Encryption string
	FromAddr   string
	FromName   string
}

type Message struct {
	From    string
	To      []string
	Subject string
	Body    string
}

func (m Message) Raw() []byte {
	return []byte(m.Subject + "\n" + m.Body)
}
