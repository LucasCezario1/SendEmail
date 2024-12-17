package mail

import (
	"SendEmail/internal/campaign"
	"fmt"
	"gopkg.in/gomail.v2"
	"time"
)

func SendMail(campaign *campaign.Campaign) error {
	fmt.Println("Sending Email")

	start := time.Now()
	d := gomail.NewDialer("smtp.gmail.com", 465, "teste12@gmail.com", "51512")

	var emails []string
	for _, contact := range campaign.Contacts {
		emails = append(emails, contact.Email)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "teste12@gmail.com") // de ondem vai o email
	m.SetHeader("To", emails...)             // pra quem eu estou enviado
	m.SetHeader("Subject", campaign.Name)
	m.SetBody("text/html", campaign.Content)

	err := d.DialAndSend(m)

	duration := time.Since(start)
	fmt.Println(duration)

	return err
}
