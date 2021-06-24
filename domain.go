package sw

const (
	Endpoint         	string = "https://scrutinizer-ci.com/"
	LoginPageURL     	string = Endpoint + "login#"
	ReposPageURL     	string = Endpoint + "dashboard/repositories"
	BBOrgName        	string = "silintl"
	GHOrgName        	string = "silinternational"
	APISecretEnv     	string = "SW_GA_API_SECRET"
	MeasurementIDEnv 	string = "SW_GA_MEASUREMENT_ID"
	ScrutinizerUsrnm	string = "SCRUTINIZER_USRNM"
	ScrutinizerPsswd	string = "SCRUTINIZER_PSSWD"
	ClientID       		string = "silinternational/ga-event-tracker"
	GitHub          	string = "github"
	BitBucket			string = "bitbucket"
)
