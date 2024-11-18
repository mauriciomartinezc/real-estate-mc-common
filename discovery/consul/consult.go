package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/mauriciomartinezc/real-estate-mc-common/discovery"
	"log"
	"os"
	"strconv"
)

type consulApi struct {
	client       *api.Client
	podID        string
	serverPort   int
	isProduction bool
	httpProtocol string
}

func NewConsultApi() discovery.DiscoveryClient {
	consulAddress, podIp, isProduction := getConsulEnv()
	client, err := api.NewClient(&api.Config{Address: consulAddress})

	if err != nil {
		log.Fatalf("Error creando cliente de Consul: %v", err)
	}

	serverPort := getServerPort()
	httpProtocol := getHttpProtocol(isProduction)

	return &consulApi{
		client:       client,
		podID:        podIp,
		serverPort:   serverPort,
		isProduction: isProduction,
		httpProtocol: httpProtocol,
	}
}

func (c *consulApi) RegisterService(serviceName string) {
	c.validateClient()

	serviceAddress := c.getHealthCheckURL()

	registration := &api.AgentServiceRegistration{
		Name:    serviceName,
		Port:    c.serverPort,
		Address: c.podID,
		Check: &api.AgentServiceCheck{
			HTTP:     serviceAddress,
			Interval: "10s",
			Timeout:  "5s",
		},
	}

	err := c.client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("Error registrando el servicio en Consul: %v", err)
	}
	log.Printf("Servicio %s registrado en Consul con dirección %s:%d", serviceName, c.podID, c.serverPort)
}

func (c *consulApi) GetServiceAddress(serviceName string) (string, error) {
	c.validateClient()

	services, _, err := c.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return "", fmt.Errorf("error al consultar Consul para el servicio %s: %v", serviceName, err)
	}

	if len(services) == 0 {
		return "", fmt.Errorf("no se encontraron instancias saludables para el servicio %s", serviceName)
	}

	service := services[0]
	address := fmt.Sprintf("%s://%s:%d", c.httpProtocol, service.Service.Address, service.Service.Port)
	log.Printf("Servicio encontrado: %s en %s", serviceName, address)
	return address, nil
}

func (c *consulApi) validateClient() {
	if c.client == nil {
		log.Fatal("Error: el cliente de Consul no está configurado")
	}
}

func (c *consulApi) getHealthCheckURL() string {
	return fmt.Sprintf("%s://%s:%d/health", c.httpProtocol, c.podID, c.serverPort)
}

func getConsulEnv() (string, string, bool) {
	consulAddress := os.Getenv("CONSUL_ADDRESS")
	if consulAddress == "" {
		log.Fatal("Error: la variable de entorno CONSUL_ADDRESS no está configurada")
	}

	podIP := os.Getenv("POD_IP")
	if podIP == "" {
		log.Fatal("Error: la variable de entorno POD_IP no está configurada")
	}

	isProduction := os.Getenv("APP_ENV") == "production"
	return consulAddress, podIP, isProduction
}

func getServerPort() int {
	serverPortStr := os.Getenv("SERVER_PORT")
	serverPort, err := strconv.Atoi(serverPortStr)
	if err != nil {
		log.Fatalf("Error: la variable de entorno SERVER_PORT debe ser un número válido. Valor actual: %s", serverPortStr)
	}
	return serverPort
}

func getHttpProtocol(isProduction bool) string {
	if isProduction {
		return "https"
	}
	return "http"
}
