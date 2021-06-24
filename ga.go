package sw

import (
	"fmt"
	"os"
	"time"

	lu "github.com/captaincrazybro/literalutil"

	"github.com/silinternational/ga-event-tracker"
)

// SendDataToGA sends the average scrutinizer repository score to google analytics
func SendDataToGA(averageScore float64) error {
	envVars, err := getEnvVariables()
	if err != nil {
		return err
	}

	err = ga.SendEvent(ga.Meta{
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

// GetEnvVariables gets the environment variables required for sending data to ga
func getEnvVariables() ([]string, error) {

	// makes sure variables exist
	fmt.Println(os.Getenv(APISecretEnv))
	apiSec := os.Getenv(APISecretEnv)
	measId := os.Getenv(MeasurementIDEnv)
	if apiSec == "" {
		return []string{}, fmt.Errorf("%q environment variable has not been set, vars are %q and %q", APISecretEnv, apiSec, APISecretEnv)
	} else if measId == "" {
		return []string{}, fmt.Errorf("%q environment variable has not been set", MeasurementIDEnv)
	}

	return []string{apiSec, measId}, nil

}
