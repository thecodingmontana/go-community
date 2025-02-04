package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"sync"

	"github.com/thecodingmontana/go-community/pkg/types"
)

// AppSendMail sends emails to multiple recipients
func AppSendMail(emails []string, mail types.MailBody) error {
	APP_GMAIL_USERNAME := os.Getenv("APP_GMAIL_USERNAME")
	APP_GMAIL_PASSWORD := os.Getenv("APP_GMAIL_PASSWORD")
	SMTP_HOST := "smtp.gmail.com"
	SMTP_PORT := "587"

	if APP_GMAIL_USERNAME == "" || APP_GMAIL_PASSWORD == "" {
		return fmt.Errorf("missing SMTP credentials in environment variables")
	}

	mail.From = APP_GMAIL_USERNAME

	mailChannel := sendMail(emails, mail, types.SMTPAuth{
		Username: APP_GMAIL_USERNAME,
		Password: APP_GMAIL_PASSWORD,
		Host:     SMTP_HOST,
		Port:     SMTP_PORT,
	})

	var failedEmails []string
	for result := range mailChannel {
		if result.Err != nil {
			log.Printf("Failed to send email to %s: %v", result.Email, result.Err)
			failedEmails = append(failedEmails, result.Email)
		} else {
			log.Printf("Email sent to %s successfully!", result.Email)
		}
	}

	if len(failedEmails) > 0 {
		return fmt.Errorf("failed to send emails to: %v", failedEmails)
	}

	return nil
}

type MailResult struct {
	Email string
	Err   error
}

func sendMail(emails []string, mail types.MailBody, credentials types.SMTPAuth) chan MailResult {
	var mailWG sync.WaitGroup
	mailChannel := make(chan MailResult, len(emails))

	for _, email := range emails {
		mailWG.Add(1)
		go send(email, credentials, mail, mailChannel, &mailWG)
	}

	go func() {
		mailWG.Wait()
		close(mailChannel)
	}()
	return mailChannel
}

func send(email string, credentials types.SMTPAuth, mail types.MailBody, ch chan MailResult, wg *sync.WaitGroup) {
	defer wg.Done()

	auth := smtp.PlainAuth("", credentials.Username, credentials.Password, credentials.Host)
	addr := fmt.Sprintf("%s:%s", credentials.Host, credentials.Port)

	err := smtp.SendMail(addr, auth, mail.From, []string{email}, mail.Message)
	ch <- MailResult{Email: email, Err: err}
}
