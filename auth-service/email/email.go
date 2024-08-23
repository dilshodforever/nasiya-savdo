package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"

	"gitlab.com/lingualeap/auth/config"
)

type SendEmailRequest struct {
	To      []string
	Type    string
	Body    map[string]string
	Subject string
	Code    string
}

const (
	VerificationEmail   = "Verification email"
	ForgotPasswordEmail = "forgot_password_email"
)

func SendEmail(cfg *config.Config, req *SendEmailRequest) error {
	from := cfg.SmtpSender
	to := req.To

	password := cfg.SmtpPassword

	var body bytes.Buffer

	templatePath := getTemplatePath(req.Type)
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = t.Execute(&body, req)
	if err != nil {
		log.Fatal(err)
		return err
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := fmt.Sprintf("Subject: %s\n", req.Subject)
	msg := []byte(subject + mime + body.String())

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
	err = smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		return err
	}
	return nil
}

func getTemplatePath(emailType string) string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	switch emailType {
	case VerificationEmail:
		return dir + "/email/verification_email.html"
	case ForgotPasswordEmail:
		return dir + "/email/forgot_password_email.html"
	}

	return ""
}
