package properties

import (
	"fmt"
	"github.com/magiconair/properties"
	"log"
	"os"
	"sync"
)

const serviceEnvKey = "SERVICE_ENV"

type serviceProperties struct {
	Env         string
	Name        string
	ContextPath string
	Port        string
	Host        string
	props       *properties.Properties
}

func (s *serviceProperties) GetValue(key string) string {
	return s.props.GetString(key, "")
}

func (s *serviceProperties) GetValueOrDefault(key string, def string) string {
	return s.props.GetString(key, def)
}

var instance *serviceProperties
var once sync.Once

func GetInstance() *serviceProperties {
	once.Do(func() {
		instance = loadServiceProperties()
		if len(instance.Env) > 0 {
			log.Printf("Established active environment: " + instance.Env)
		}
	})
	return instance
}

func loadServiceProperties() *serviceProperties {
	propEnvAppend := ""
	selectedEnv := getEnv()

	if len(selectedEnv) > 0 {
		propEnvAppend = fmt.Sprintf("-%s", selectedEnv)
	}

	propFile := fmt.Sprintf("service%s.properties", propEnvAppend)
	p := properties.MustLoadFile(propFile, properties.UTF8)

	return &serviceProperties{
		Env:         getEnv(),
		Name:        p.GetString("service.name", "service"),
		ContextPath: p.GetString("service.context-path", ""),
		Port:        p.GetString("service.port", "8080"),
		Host:        p.GetString("service.host", "0.0.0.0"),
		props:       p,
	}
}

func getEnv() string {
	return os.Getenv(serviceEnvKey)
}
