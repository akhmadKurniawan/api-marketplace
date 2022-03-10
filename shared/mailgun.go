package shared

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type Mailgun struct {
	Sender    string
	Subject   string
	Body      string
	Recipient string
}

func SendMailgun(mail Mailgun) error {
	privateAPIKey := os.Getenv("MAILGUN_API_KEY")
	var Domain string = os.Getenv("MAILGUN_DOMAIN")
	mg := mailgun.NewMailgun(Domain, privateAPIKey)

	sender := mail.Sender
	subject := mail.Subject
	body := mail.Body
	recipient := mail.Recipient

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	_, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Println("MAILGUN - error :", err)
		return err
	}

	fmt.Printf("ID: %s\n", id)
	return nil
}
