package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

func main() {

	// Sender data.
	from := "cakmakforbusiness@gmail.com"
	password := "*****"

	// Receiver email address.
	to := []string{
		"mertcakmak2@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("template/template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Describe   string
		ProductUrl string
		ImageUrl   string
		Price      string
	}{
		Describe:   "This is a test message in a HTML template",
		ProductUrl: "https://www.trendyol.com/nike/turkiye-milli-takim-ic-saha-erkek-futbol-formasi-euro-2020-euro2020-p-114318607?v=s",
		ImageUrl:   "https://cdn.dsmcdn.com/ty22/product/media/images/20201109/10/24125717/102876292/1/1_org_zoom.jpeg",
		Price:      "309",
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
