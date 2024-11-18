package discovery

type DiscoveryClient interface {
	RegisterService(serviceName string)
	GetServiceAddress(serviceName string) (string, error)
}
