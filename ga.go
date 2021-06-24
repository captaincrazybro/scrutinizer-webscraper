package sw

import (
	"fmt"
	"time"

	lu "github.com/captaincrazybro/literalutil"
	"github.com/silinternational/ga-event-tracker"
)

// SendDataToGA sends the average scrutinizer repository score to google analytics
func SendDataToGA(averageScore float64) error {
	envVars := []string{APISecretEnv, MeasurementIDEnv}
	GetEnvVariables(envVars)

	err := ga.SendEvent(ga.Meta{
		APISecret:     envVars[0],
		ClientID:      ClientID,
		MeasurementID: envVars[1],
	}, []ga.Event{
		{
			Name: "weekly_score_average",
			Params: ga.Params{
				"score": averageScore,
				"date":  timeToDateString(time.Now()),
			},
		},
	})

	return err
}

func timeToDateString(time time.Time) lu.String {
	y, m, d := time.Date()
	return lu.String(fmt.Sprintf("%s/%d/%d", m, d, y))
}
