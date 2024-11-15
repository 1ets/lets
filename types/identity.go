package types

// Serve information
type IdentityService struct {
	Name        string
	Description string
	ServiceDns  string `desc:"Service DNS"`
	Layer       string `desc:"Microservice Layer"`
	Author      string
}

type IdentitySource struct {
	Repository    string `desc:"Repository URL"`
	Documentation string `desc:"Documentation URL"`
}
