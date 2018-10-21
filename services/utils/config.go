package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

var once sync.Once
var instance *Config

// GetConfig get config defined in config.json
func GetConfig() *Config {
	once.Do(func() {
		env := os.Getenv("DF_ENVIROMENT")
		if env == "" {
			env = "dev"
		}

		var config *Config
		var filepath string
		pwd, _ := os.Getwd()

		if flag.Lookup("test.v") == nil {
			// normal run
			filepath = path.Join(pwd, fmt.Sprintf("config.%s.json", strings.ToLower(env)))
		} else {
			// under go test
			filepath = path.Join(pwd, "testdata", "config.unit.test.json")
		}

		configFile, err := os.Open(filepath)
		defer configFile.Close()
		if err != nil {
			panic(err)
		}

		jsonParser := json.NewDecoder(configFile)
		err = jsonParser.Decode(&config)
		if err != nil {
			panic(err)
		}

		instance = config
	})

	return instance
}

// Redis redis config
type Redis struct {
	Host     string `json:"host"`
	Password string `json:"password,omitempty"`
	Size     int    `json:"maxconnectioncount"`
}

// MySQL mysql config
type MySQL struct {
	Host       string `json:"host"`
	DataBase   string `json:"database"`
	User       string `json:"user"`
	Password   string `json:"password"`
	AllowDebug bool   `json:"allow_debug"`
}

// Nats nats config
type Nats struct {
	Host   string `json:"host"`
	Size   int    `json:"maxconnectioncount"`
	Topics string `json:"topics"`
}

// Services external services like Mysql
type Services struct {
	Redis `json:"redis"`
	MySQL `json:"mysql"`
	Nats  `json:"nats"`
}

// ServiceNames servicename config
type ServiceNames struct {
	CoreServer string `json:"core_server"`
}

// ServiceMeta service meta data including service discovery specs
type ServiceMeta struct {
	RegistryTTL       int    `json:"registry_ttl"`
	RegistryHeartbeat int    `json:"registry_heartbeat"`
	Version           string `json:"api_version"`
}

// Config config entry
type Config struct {
	Services     `json:"services"`
	ServiceNames `json:"servicenames"`
	ServiceMeta  `json:"servicemeta"`
}

// GetNatsTopics convert config string to topic array
func GetNatsTopics(topic string) []string {
	return strings.Split(topic, ",")
}

// GetMaxConnectionCount get redis max connection count
func (r *Redis) GetMaxConnectionCount() int {
	if r.Size == 0 {
		return 100
	}
	return r.Size
}

// GetMaxConnectionCount get nats max connection count
func (n *Nats) GetMaxConnectionCount() int {
	if n.Size == 0 {
		return 100
	}
	return n.Size
}

// GetRegistryTTL get registry ttl
func (s *ServiceMeta) GetRegistryTTL() time.Duration {
	if s.RegistryTTL == 0 {
		return 30 * time.Second
	}

	return time.Duration(s.RegistryTTL) * time.Second
}

// GetRegistryHeartbeat get registry heartbeat
func (s *ServiceMeta) GetRegistryHeartbeat() time.Duration {
	if s.RegistryHeartbeat == 0 {
		return 10 * time.Second
	}

	return time.Duration(s.RegistryHeartbeat) * time.Second
}

// GetVersion get api version
func (s *ServiceMeta) GetVersion() string {
	return s.Version
}
