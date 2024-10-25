package rabbit

import (
	"bytes"
	"text/template"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Connect returns a RabbitMQ connection pointer.
func Connect(protocol, username, password, hostname, port string) (*amqp.Connection, error) {
	tpl := template.New("connectionParams")
	tplSrc := "{{ .Protocol }}://{{ .Username }}:{{ .Password }}@{{ .Hostname }}:{{ .Port }}"
	tpl, err := tpl.Parse(tplSrc)
	if err != nil {
		return nil, err
	}
	w := &bytes.Buffer{}
	params := connectionParams{
		Protocol: protocol,
		Username: username,
		Password: password,
		Hostname: hostname,
		Port:     port,
	}
	if err := tpl.ExecuteTemplate(w, "connectionParams", params); err != nil {
		return nil, err
	}
	return amqp.Dial(w.String())
}

type connectionParams struct {
	Protocol string
	Username string
	Password string
	Hostname string
	Port     string
}
