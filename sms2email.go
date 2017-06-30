package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func receiveMessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sendEmail(r.FormValue("From"), r.FormValue("Body"))
}

func sendEmail(number string, body string) {
	from := mail.NewEmail(number, "federico.builes@gmail.com")
	subject := "New SMS"
	to := mail.NewEmail("Federico Builes", "federico.builes@gmail.com")
	content := mail.NewContent("text/plain", body)

	message := mail.NewV3MailInit(from, subject, to, content)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func main() {
	fmt.Println(os.Getenv("SENDGRID_API_KEY"))
	fmt.Println("Starting server...")
	http.HandleFunc("/send_message", receiveMessage)
	http.ListenAndServe(":8080", nil)
}
