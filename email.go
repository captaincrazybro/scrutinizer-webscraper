package sw

import (
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	lu "github.com/captaincrazybro/literalutil"
	c "github.com/captaincrazybro/literalutil/console"
)

const (
	region lu.String = "us-west-2"
)

// SendEmail sends an email with a certain template and data
func SendEmail(template string, data string) {
	// creates session of ses
	config := &aws.Config{
		Region:      aws.String(region.Tos()),
		Credentials: credentials.NewEnvCredentials(),
	}

	sess := session.Must(session.NewSession(config))
	svc := ses.New(sess)
	from := os.Getenv(EmailUsrnm)
	to := os.Getenv(ToEmails)
	if from == "" || to == "" {
		c.Flnf("environment variable %s has not been set", lu.STernary(from == "", EmailUsrnm, ToEmails))
	}
	to = strings.Trim(to, " ")
	emails := strings.Split(to, ",")
	toEmails := toPtrString(emails)

	input := &ses.SendTemplatedEmailInput{
		Source:   &from,
		Template: &template,
		Destination: &ses.Destination{
			// TODO: change this to env variable
			ToAddresses: toEmails,
		},
		TemplateData: &data,
	}

	_, err := svc.SendTemplatedEmail(input)
	if err != nil {
		return
	}
}

/*
// SendEmail sends an email of the repositories and their average quality score
func SendRepoMail() {
	envVars := []string{ScrutinizerUsrnm, ScrutinizerPsswd}
	GetEnvVariables(envVars)

	//Copied from the first gomail example
	m := gomail.NewMessage()
	m.SetHeader("From", envVars[0])
	m.SetHeader("To", "camerongordon111@gmail.com")
	m.SetHeader("Subject", "S")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer("smtp.gmail.com", 587, envVars[0], envVars[1])

	// Send the email to EmailUsername.
	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("could not send email:\n%s", err)
	}
}*/

func SendScoreMail() {

}

func toPtrString(a []string) []*string {
	var pa []*string
	for _, v := range a {
		pa = append(pa, &v)
	}
	return pa
}
