package config

import (
	"fmt"
	"time"

	"github.com/Kortivex/connected_roots/pkg/httpserver"
)

type App struct {
	Env      string `koanf:"env"`
	Name     string `koanf:"name"`
	LogLevel string `koanf:"loglevel"`
}

type API struct {
	Protocol string `koanf:"protocol"`
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	Debug    bool   `koanf:"debug"`
	Recover  bool   `koanf:"recover"`
	Body     string `koanf:"body"`
	APIKey   string `koanf:"api"`
	Timeouts struct {
		Shutdown int `koanf:"shutdown"`
		Read     int `koanf:"read"`
		Write    int `koanf:"write"`
		Idle     int `koanf:"idle"`
	} `koanf:"timeouts"`
	Health struct {
		Frequency int `koanf:"frequency"`
	} `koanf:"health"`
}

type Frontend struct {
	Protocol  string `koanf:"protocol"`
	Host      string `koanf:"host"`
	Port      int    `koanf:"port"`
	Templates string `koanf:"templates"`
	Assets    string `koanf:"assets"`
	Cookie    struct {
		Name   string `koanf:"name"`
		MaxAge int    `koanf:"maxage"`
		Key    string `koanf:"key"`
		Table  string `koanf:"table"`
	} `koanf:"cookie"`
	I18n struct {
		Path string `koanf:"path"`
		En   string `koanf:"en"`
		Es   string `koanf:"es"`
	}
	Debug    bool   `koanf:"debug"`
	Recover  bool   `koanf:"recover"`
	Body     string `koanf:"body"`
	Timeouts struct {
		Shutdown int `koanf:"shutdown"`
		Read     int `koanf:"read"`
		Write    int `koanf:"write"`
		Idle     int `koanf:"idle"`
	} `koanf:"timeouts"`
	Health struct {
		Frequency int `koanf:"frequency"`
	} `koanf:"health"`
}

type DB struct {
	Postgres struct {
		DSN    string `koanf:"dsn"`
		Logger struct {
			SlowThreshold             int  `koanf:"slowthreshold"`
			IgnoreRecordNotFoundError bool `koanf:"ignorerecordnotfounderror"`
			Colorful                  bool `koanf:"colorful"`
		} `koanf:"logger"`
		Connection struct {
			MaxIdleConns    int `koanf:"maxidleconns"`
			MaxOpenConns    int `koanf:"maxopenconns"`
			ConnMaxIdleTime int `koanf:"connmaxidletime"`
			ConnMaxLifetime int `koanf:"connmaxlifetime"`
		} `koanf:"connection"`
		Version int `koanf:"version"`
		Health  struct {
			Frequency int `koanf:"frequency"`
		} `koanf:"health"`
	} `koanf:"postgres"`
}

type Thirds struct {
	SDK struct {
		Verbose               bool `koanf:"verbose"`
		ConnectedRootsService struct {
			Host   string `koanf:"host"`
			APIKey string `koanf:"api"`
		} `koanf:"connectedrootsservice"`
	} `koanf:"sdk"`
}

type Monitoring struct {
	Metrics struct {
		Active     bool `koanf:"active"`
		Prometheus struct {
			Disabled        bool   `koanf:"disabled"`
			ServiceBackend  string `koanf:"servicebackend"`
			ServiceFrontend string `koanf:"servicefrontend"`
			Path            string `koanf:"path"`
		} `koanf:"prometheus"`
	} `koanf:"metrics"`
	Observability struct {
		Active bool `koanf:"active"`
		Otel   struct {
			Disabled    bool   `koanf:"sdkdisabled"`
			DumpEnabled bool   `koanf:"bodydumpenabled"`
			Service     string `koanf:"servicename"`
			Addr        string `koanf:"exporterotlpendpoint"`
		} `koanf:"otel"`
	} `koanf:"observability"`
}

type Config struct {
	App        `koanf:"app"`
	API        `koanf:"api"`
	Frontend   `koanf:"frontend"`
	DB         `koanf:"db"`
	Thirds     `koanf:"thirds"`
	Monitoring `koanf:"monitoring"`
}

func (c *Config) GetAPIParams() httpserver.Params {
	readTimeout := time.Duration(c.API.Timeouts.Read) * time.Second
	writeTimeout := time.Duration(c.API.Timeouts.Write) * time.Second
	idleTimeout := time.Duration(c.API.Timeouts.Idle) * time.Second

	return httpserver.Params{
		Port:                  fmt.Sprintf("%d", c.API.Port),
		BodyLimit:             c.API.Body,
		PrometheusServiceName: c.Monitoring.Metrics.Prometheus.ServiceBackend,
		PrometheusDisabled:    &c.Monitoring.Metrics.Prometheus.Disabled,
		WriteTimeout:          &writeTimeout,
		ReadTimeout:           &readTimeout,
		IdleTimeout:           &idleTimeout,
		RecoverDisabled:       !c.API.Recover,
	}
}

func (c *Config) GetFrontendParams() httpserver.Params {
	readTimeout := time.Duration(c.Frontend.Timeouts.Read) * time.Second
	writeTimeout := time.Duration(c.Frontend.Timeouts.Write) * time.Second
	idleTimeout := time.Duration(c.Frontend.Timeouts.Idle) * time.Second

	return httpserver.Params{
		Port:                  fmt.Sprintf("%d", c.Frontend.Port),
		BodyLimit:             c.Frontend.Body,
		PrometheusServiceName: c.Monitoring.Metrics.Prometheus.ServiceFrontend,
		PrometheusDisabled:    &c.Monitoring.Metrics.Prometheus.Disabled,
		WriteTimeout:          &writeTimeout,
		ReadTimeout:           &readTimeout,
		IdleTimeout:           &idleTimeout,
		RecoverDisabled:       !c.Frontend.Recover,
	}
}
