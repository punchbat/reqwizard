package email

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Mailer struct {
	smtpConfig *SMTP
	auth       smtp.Auth
}

type SMTP struct {
	HOST, PORT, FROM, PASSWORD string
}

func (s *SMTP) Address() string {
	return s.HOST + ":" + s.PORT
}

// NewMailer returns a configured SMTP Mailer.
func NewMailer() *Mailer {
	log.Println("New mailer")

	var smtpConfig = &SMTP{
		viper.GetString("services.email.SMTP_HOST"),
		viper.GetString("services.email.SMTP_PORT"),
		viper.GetString("services.email.SMTP_FROM"),
		viper.GetString("services.email.SMTP_PASSWORD"),
	}

	// Authentication
	var auth = smtp.PlainAuth("", smtpConfig.FROM, smtpConfig.PASSWORD, smtpConfig.HOST)

	return &Mailer{
		smtpConfig, auth,
	}
}

type Message struct {
	Subject      string
	To           []string
	Content      interface{}
	TemplateName string
}

// Send sends the mail via smtp.
func (m *Mailer) Send(email *Message) error {
	body, err := ParseTemplate("./internal/templates/email", email.TemplateName, email.Content)

	if err != nil {
		log.Println("Error parsing template:", err)

		return err
	}

	var to = email.To[0]

	message := []byte(
		"To: " + to + "\r\n" +
			"Subject: " + email.Subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"utf-8\"\r\n" + "\r\n" +
			body,
	)

	err = smtp.SendMail(m.smtpConfig.Address(), m.auth, m.smtpConfig.FROM, email.To, message)

	if err != nil {
		log.Println("Error sengind message", err)

		return err
	}

	return nil
}

func ParseTemplate(templateDir string, TemplateName string, data interface{}) (string, error) {
	templates := template.New("")

	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			_, err = templates.ParseFiles(path)
		}
		return err
	})

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = templates.ExecuteTemplate(buf, TemplateName, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}