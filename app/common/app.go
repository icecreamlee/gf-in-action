package common

var (
	ENV      string
	RootPath string
	RunMode  string
)

const (
	ENVDev     = "dev"
	ENVTest    = "test"
	ENVBeta    = "beta"
	ENVProd    = "prod"
	RunModeCLI = "CLI"
	RunModeWeb = "WEB"
)
