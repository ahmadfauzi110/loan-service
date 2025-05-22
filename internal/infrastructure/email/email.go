package email

import (
	"fmt"
	"log"

	sib "github.com/sendinblue/APIv3-go-library/v2/lib"

	"github.com/ahmadfauzi110/loan-service/config"
	port "github.com/ahmadfauzi110/loan-service/internal/port/email"
)

type BrevoEmailService struct {
	Client *sib.APIClient
	From   sib.SendSmtpEmailSender
}

func NewBrevoEmailService(brevoConfig *config.BREVO) port.EmailSender {
	cfg := sib.NewConfiguration()
	cfg.AddDefaultHeader("api-key", brevoConfig.API_KEY)

	return &BrevoEmailService{
		Client: sib.NewAPIClient(cfg),
		From: sib.SendSmtpEmailSender{
			Name:  brevoConfig.SENDER_NAME,
			Email: brevoConfig.SENDER_EMAIL,
		},
	}
}

func (s *BrevoEmailService) Send(to, subject, body string) error {
	toList := []sib.SendSmtpEmailTo{{Email: to}}

	fmt.Printf("toList: %v", toList)

	fmt.Printf("subject: %v", subject)
	fmt.Printf("body: %v", body)

	req := sib.SendSmtpEmail{
		Sender:      &s.From,
		To:          toList,
		Subject:     subject,
		TextContent: body,
	}

	_, _, err := s.Client.TransactionalEmailsApi.SendTransacEmail(nil, req)
	if err != nil {
		log.Printf("Brevo send error: %v", err.Error())
		return err
	}

	return nil
}
