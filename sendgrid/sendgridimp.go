package sendgrid

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendgridstruct struct {
	Sendername    string
	Sendermail    string
	Sub           string
	Recipientname string
	Recipientmail string
	Contents      string
	templatehtml  string
}

//Sendgrid implements the MailService interface
type Sendgrid struct{}

//New return a new Sendgrid instance
func New() *Sendgrid {
	return &Sendgrid{}
}

//Send method from interface
func (sg *Sendgrid) Send(s *sendgridstruct) {
	from := mail.NewEmail(s.Sendername, s.Sendermail)
	subject := s.Sub
	to := mail.NewEmail(s.Recipientname, s.Recipientmail)
	plainTextContent := s.Contents
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, "")
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

//SendWithTemplate method
func (sg *Sendgrid) SendWithTemplate(s *sendgridstruct) {
	from := mail.NewEmail(s.Sendermail, s.Sendermail)
	subject := s.Sub
	to := mail.NewEmail(s.Recipientname, s.Recipientmail)
	plainTextContent := s.Contents
	htmlContent := s.templatehtml
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
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

func (sg *Sendgrid) SendMultiple() {}
