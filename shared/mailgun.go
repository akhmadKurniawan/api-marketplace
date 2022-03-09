package shared

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

func SendMailgun(email string) (err error) {
	privateAPIKey := os.Getenv("MAILGUN_API_KEY")
	var Domain string = os.Getenv("MAILGUN_DOMAIN")

	mg := mailgun.NewMailgun(Domain, privateAPIKey)

	sender := "sender@example.com"
	subject := "Fancy subject!"
	body := "Hello from Mailgun Go!"
	recipient := email

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
