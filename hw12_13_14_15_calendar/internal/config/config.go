package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DB     DB         `json:"db"`
	GRPC   GRPC       `json:"grpc"`
	HTTP   HTTP       `json:"http"`
	Logger LoggerConf `json:"logger"`
}

type LoggerConf struct {
	LogLevel string `json:"log_level"`
}

type DB struct {
	DSN                string `json:"dsn"`
	MaxOpenConnections int32  `json:"max_open_connections"`
}

type GRPC struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type HTTP struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func NewConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (l *LoggerConf) Level() string {
	return l.LogLevel
}

// TODO
