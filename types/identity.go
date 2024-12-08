package types

import "net"

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

type IdentityNetwork struct {
	Hostname     []string `desc:"Hostname"`
	IPV4         []net.IP `desc:"IPV4 Address"`
	IPV6         []net.IP `desc:"IPV6 Address"`
	NetInterface []string `desc:"Network Interface"`
}

type Replica struct {
	IPV4 []net.IP `desc:"Replica IPV4 Address"`
	IPV6 []net.IP `desc:"Replica IPV6 Address"`
}
