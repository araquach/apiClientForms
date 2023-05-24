package helpers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v3"
	"html/template"
	"log"
	"os"
	"salonHub/cmd/models"
	"time"
)

type EmailInfo struct {
	Name  string
	Salon string
}

func SendSkinTestEmail(data models.SkinTest) {
	var salonName, salonEmail string

	switch data.Salon {
	case 1:
		salonName = "Jakata Salon"
		salonEmail = "info@jakatasalon.co.uk"
	case 2:
		salonName = "Paul Kemp Hairdressing"
		salonEmail = "info@paulkemphairdressing.com"
	case 3:
		salonName = "Base Hairdressing"
		salonEmail = "info@basehairdressing.com"
	}

	info := EmailInfo{
		Name:  data.FirstName,
		Salon: salonName,
	}

	htmlContent, err := ParseEmailTemplate("templates/skintest.gohtml", info)
	if err != nil {
		log.Fatalln(err)
	}

	textContent, err := ParseEmailTemplate("templates/skintest.txt", info)
	if err != nil {
		log.Fatalln(err)
	}

	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_KEY"))

	sender := salonEmail
	subject := "Allergy Alert Test for " + salonName
	body := textContent
	recipient := data.Email

	m := mg.NewMessage(sender, subject, body, recipient)

	m.SetHtml(htmlContent)
	m.AddAttachment("output/skinTest/skintest.pdf")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message	with a 10 second timeout
	resp, id, err := mg.Send(ctx, m)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}

func SendExtensionsEmail(data models.Extensions) {
	var salonName, salonEmail string

	switch data.Salon {
	case 1:
		salonName = "Jakata Salon"
		salonEmail = "info@jakatasalon.co.uk"
	case 2:
		salonName = "Paul Kemp Hairdressing"
		salonEmail = "info@paulkemphairdressing.com"
	case 3:
		salonName = "Base Hairdressing"
		salonEmail = "info@basehairdressing.com"
	}

	info := EmailInfo{
		Name:  data.FirstName,
		Salon: salonName,
	}

	htmlContent, err := ParseEmailTemplate("templates/extensions.gohtml", info)
	if err != nil {
		log.Fatalln(err)
	}

	textContent, err := ParseEmailTemplate("templates/extensions.txt", info)
	if err != nil {
		log.Fatalln(err)
	}

	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_KEY"))

	sender := salonEmail
	subject := "Extensions Info for " + salonName
	body := textContent
	recipient := data.Email

	m := mg.NewMessage(sender, subject, body, recipient)

	m.SetHtml(htmlContent)
	m.AddAttachment("output/extensions/extensions.pdf")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message	with a 10 second timeout
	resp, id, err := mg.Send(ctx, m)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}

func ParseEmailTemplate(templateFileName string, data interface{}) (content string, err error) {

	tmpl, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err := tmpl.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
