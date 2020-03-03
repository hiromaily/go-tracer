package config

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

// Config is root of toml config
type Config struct {
	Server *ServerConfig `validate:"required"`
	Tracer *TracerConfig `validate:"required"`
}

// ServerConfig is for server information
type ServerConfig struct {
	Port int `toml:"port" validate:"required"`
}

// TracerConfig is for tracer information
type TracerConfig struct {
	Type    string              `toml:"type" validate:"required"`
	Jaeger  *TracerDetailConfig `toml:"jaeger"`
	Datadog *TracerDetailConfig `toml:"datadog"`
}

// TracerDetailConfig is detail of tracer
type TracerDetailConfig struct {
	ServiceName       string  `toml:"service_name" validate:"required"`
	CollectorEndpoint string  `toml:"collector_endpoint" validate:"required"`
	Sampling          float64 `toml:"sampling_probability" validate:"required"`
}

// load configfile
func loadConfig(path string) (*Config, error) {
	if path == "" {
		return nil, errors.New("file path is required")
	}

	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to read %s", path)
	}

	var conf Config
	md, err := toml.Decode(string(d), &conf)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to parse %s: %v", path, md)
	}

	if err = conf.validate(); err != nil {
		return nil, errors.Wrap(err, "fail to validate config")
	}

	return &conf, nil
}

func (c *Config) validate() error {
	v := validator.New()
	return v.Struct(c)
}

// New is for creating config instance
func New(file string) (*Config, error) {
	conf, err := loadConfig(file)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
