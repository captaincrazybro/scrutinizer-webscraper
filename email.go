package sw

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	lu "github.com/captaincrazybro/literalutil"
	c "github.com/captaincrazybro/literalutil/console"
)

const (
	region lu.String = "us"
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
	to := lu.String(os.Getenv(ToEmails))
	if from == "" || to == "" {
		c.Flnf("environment variable %s has not been set", lu.STernary(from == "", EmailUsrnm, ToEmails))
	}
	emails := to.Trim(" ").Split(",")
	toEmails := toPtrString(emails.Tosa())

	input := &ses.SendTemplatedEmailInput{
		Source:   &from,
		Template: &template,
		Destination: &ses.Destination{
			ToAddresses: toEmails,
		},
		TemplateData: &data,
	}

	_, err := svc.SendTemplatedEmail(input)
	if err != nil {
		return
	}
}

func toPtrString(a []string) []*string {
	var pa []*string
	for _, v := range a {
		pa = append(pa, &v)
	}
	return pa
}
