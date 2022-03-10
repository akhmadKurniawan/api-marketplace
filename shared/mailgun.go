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
	mail = Mailgun{
		mail.Sender,
		mail.Subject,
		mail.Body,
		mail.Recipient,
	}

	privateAPIKey := os.Getenv("MAILGUN_API_KEY")
	var Domain string = os.Getenv("MAILGUN_DOMAIN")
	mg := mailgun.NewMailgun(Domain, privateAPIKey)

	sender := mail.Sender
	subject := mail.Subject
	body := mail.Body
	recipient := mail.Recipient

	fmt.Println(sender)
	fmt.Println(subject)
	fmt.Println(body)
	fmt.Println(recipient)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Println("MAILGUN - error :", err)
		return err
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}
