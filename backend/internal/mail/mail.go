package mail

import (
	"backend/internal/config"
	"bytes"
	"embed"
	"github.com/resend/resend-go/v2"
	"html/template"
	"log"
)

//go:embed "templates"
var templateFS embed.FS

type MailClient struct {
	client *resend.Client
}

func NewMailClient() *MailClient {
	client := resend.NewClient(config.Env.ResendAPI)

	return &MailClient{
		client: client,
	}
}

func (m *MailClient) SendEmail(to string, templateFile string, data interface{}) error {

	tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)
	if err != nil {
		log.Fatal(err)
		return err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		log.Fatal(err)
		return err
	}

	plainBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)
	if err != nil {
		log.Fatal(err)
		return err
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		log.Fatal(err)
		return err
	}

	params := &resend.SendEmailRequest{
		From:    "Voltig <no-reply@voltig.dev>",
		To:      []string{to},
		Subject: subject.String(),
		Html:    htmlBody.String(),
		Text:    plainBody.String(),
	}

	_, err = m.client.Emails.Send(params)
	if err != nil {
		return err
	}

	return nil
}
