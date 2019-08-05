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

// OSS sso sevice config
type OSS struct {
	APIKey     string `json:"api_key"`
	APISecret  string `json:"api_secret"`
	Endpoint   string `json:"endpoint"`
	BucketName string `json:"bucket_name"`
}

// Aliyun aliyun services
type Aliyun struct {
	OSS `json:"oss"`
}

// Auth authing fields
type Auth struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Services external services like Mysql
type Services struct {
	Redis  `json:"redis"`
	MySQL  `json:"mysql"`
	Aliyun `json:"aliyun"`
}

// Ebook ebook related config
type Ebook struct {
	Width            float64 `json:"width"`
	Height           float64 `json:"height"`
	OriginDir        string  `json:"origin_dir"`
	PDFDestDir       string  `json:"pdf_dest_dir"`
	ImageDestDir     string  `json:"img_dest_dir"`
	MergeTargetDir   string  `json:"merge_target_dir"`
	MergeDestDir     string  `json:"merge_dest_dir"`
	ImageLoadTimeout int64   `json:"image_load_timeout"`
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
	Auth 		`json:"auth"`
	Services     `json:"services"`
	Ebook        `json:"ebook"`
	ServiceNames `json:"servicenames"`
	ServiceMeta  `json:"servicemeta"`
}

// GetMaxConnectionCount get redis max connection count
func (r *Redis) GetMaxConnectionCount() int {
	if r.Size == 0 {
		return 100
	}
	return r.Size
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
