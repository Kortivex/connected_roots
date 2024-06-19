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

type HTTP struct {
	Protocol  string `koanf:"protocol"`
	Host      string `koanf:"host"`
	Port      int    `koanf:"port"`
	Templates string `koanf:"templates"`
	Assets    string `koanf:"assets"`
	Debug     bool   `koanf:"debug"`
	Recover   bool   `koanf:"recover"`
	Body      string `koanf:"body"`
	Timeouts  struct {
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

type Monitoring struct {
	Metrics struct {
		Active     bool `koanf:"active"`
		Prometheus struct {
			Disabled bool   `koanf:"disabled"`
			Service  string `koanf:"service"`
			Path     string `koanf:"path"`
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
	HTTP       `koanf:"http"`
	DB         `koanf:"db"`
	Monitoring `koanf:"monitoring"`
}

func (c *Config) GetHTTPParams() httpserver.Params {
	readTimeout := time.Duration(c.HTTP.Timeouts.Read) * time.Second
	writeTimeout := time.Duration(c.HTTP.Timeouts.Write) * time.Second
	idleTimeout := time.Duration(c.HTTP.Timeouts.Idle) * time.Second

	return httpserver.Params{
		Port:                  fmt.Sprintf("%d", c.HTTP.Port),
		BodyLimit:             c.HTTP.Body,
		PrometheusServiceName: c.Monitoring.Metrics.Prometheus.Service,
		PrometheusDisabled:    &c.Monitoring.Metrics.Prometheus.Disabled,
		WriteTimeout:          &writeTimeout,
		ReadTimeout:           &readTimeout,
		IdleTimeout:           &idleTimeout,
		RecoverDisabled:       !c.Recover,
	}
}
