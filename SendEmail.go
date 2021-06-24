package sw

import (
	"github.com/go-gomail/gomail"
)

func SendEmail(/*repo []string, avg float64*/) {
	//Copied from the first gomail example
	m := gomail.NewMessage()
	m.SetHeader("From", "cameron_gordon@sil.org")
	m.SetHeader("To", "camerongordon111@gmail.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "S")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "user", "123456")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}