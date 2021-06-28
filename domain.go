package sw

const (
	Endpoint             string = "https://scrutinizer-ci.com/"
	LoginPageURL         string = Endpoint + "login#"
	ReposPageURL         string = Endpoint + "dashboard/repositories"
	BBOrgName            string = "silintl"
	GHOrgName            string = "silinternational"
	APISecretEnv         string = "SW_GA_API_SECRET"
	MeasurementIDEnv     string = "SW_GA_MEASUREMENT_ID"
	ScrutinizerUsrnm     string = "SCRUTINIZER_USRNM"
	ScrutinizerPsswd     string = "SCRUTINIZER_PSSWD"
	EmailUsrnm           string = "EMAIL_USRNM"
	EmailPsswd           string = "EMAIL_PSSWD"
	ClientID             string = "silinternational/ga-event-tracker"
	GitHub               string = "github"
	BitBucket            string = "bitbucket"
	IgnoredReposFileName string = "ignored-repos"
	ToEmails             string = "SW_TO_EMAILS"
	S3ObjectKeyEnv       string = "SW_OBJECT_KEY"
	S3BucketEnv          string = "SW_BUCKET"
)
