package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
)

func main() {
	from := mail.NewEmail("Example User", "zendragon93@gmail.com")
	subject := "Sending with Twilio SendGrid is Fun"
	to := mail.NewEmail("Example User", "khai.luong@tpptechnology.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("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")
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
