package mail

import (
	"fmt"
	"net/smtp"
)

// Mailer handles email sending
type Mailer struct {
	host     string
	port     int
	username string
	password string
	from     string
	fromName string
}

// New creates a new mailer
func New(host string, port int, username, password, from, fromName string) *Mailer {
	return &Mailer{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
		fromName: fromName,
	}
}

// Send sends an email
func (m *Mailer) Send(to, subject, body string) error {
	if m.host == "" {
		return fmt.Errorf("mail not configured")
	}

	auth := smtp.PlainAuth("", m.username, m.password, m.host)
	addr := fmt.Sprintf("%s:%d", m.host, m.port)

	msg := []byte(fmt.Sprintf("From: %s <%s>\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=utf-8\r\n"+
		"\r\n"+
		"%s", m.fromName, m.from, to, subject, body))

	return smtp.SendMail(addr, auth, m.from, []string{to}, msg)
}

// SendVerification sends a verification email
func (m *Mailer) SendVerification(to, token string) error {
	subject := "Verify your email"
	body := fmt.Sprintf("<p>Please verify your email by clicking <a href=\"%s/verify?token=%s\">here</a></p>",
		"", token) // TODO: Add app URL
	return m.Send(to, subject, body)
}

// SendPasswordReset sends a password reset email
func (m *Mailer) SendPasswordReset(to, token string) error {
	subject := "Reset your password"
	body := fmt.Sprintf("<p>Click <a href=\"%s/reset-password?token=%s\">here</a> to reset your password</p>",
		"", token) // TODO: Add app URL
	return m.Send(to, subject, body)
}
