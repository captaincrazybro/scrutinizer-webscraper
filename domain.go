package sw

const (
	Endpoint     string = "https://scrutinizer-ci.com/"
	LoginPageURL string = Endpoint + "login#"
	ReposPageURL string = Endpoint + "dashboard/repositories"
	BBOrgName    string = "silintl"
	GHOrgName    string = "silinternational"
	DivStyle     string = `text-align: center; color:#fff; font-size:18px; font-weight: bold;  background-repeat: repeat-x; background-image:linear-gradient(to bottom,
		#27ae60, #2c9244, #2c9244, #27ae60
);
background-color: #CCC; background-repeat: repeat-x; border:rgba(0, 0, 0, 0.1) rgba(0, 0, 0, 0.1) rgba(0, 0, 0, 0.25);
padding:11px 15px; `
	APISecretEnv     string = "SW_GA_API_SECRET"
	MeasurementIDEnv string = "SW_GA_MEASUREMENT_ID"
	EmailUsername	 string = "EMAIL_USERNAME"
	EmailPassword	 string = "EMAIL_PASSWORD"
	ClientID         string = "silinternational/ga-event-tracker"
)
