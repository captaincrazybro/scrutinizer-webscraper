package sw

import (
	"log"

	"github.com/go-gomail/gomail"
)

// SendEmail sends an email of the repositories and their average quality score
func SendEmail(/*repo []string, avg float64*/) {
	envVars := []string{ScrutinizerUsrnm, ScrutinizerPsswd}
	GetEnvVariables(envVars)

	//Copied from the first gomail example
	m := gomail.NewMessage()
	m.SetHeader("From", envVars[0])
	m.SetHeader("To", "camerongordon111@gmail.com")
	m.SetHeader("Subject", "S")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, envVars[0], envVars[1])

	// Send the email to EmailUsername.
	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("could not send email:\n%s", err)
	}
}