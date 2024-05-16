package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string     `yaml:"env" env-default:"local"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HTTPServer  HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address      string        `yaml:"address" env-default:"localhost:8080"`
	Timeout      time.Duration `yaml:"timeout" env-degault:"4s"`
	Idle_Timeout time.Duration `yaml:"idle_timeout" env-defaul:"60s"`
}

type ConfigDB struct {
	Account    ConfigParm `yaml:"account"`
	Battle     ConfigParm `yaml:"battle"`
	Billing    ConfigParm `yaml:"billing"`
	Game       ConfigParm `yaml:"game"`
	Logs       ConfigParm `yaml:"logs"`
	Parm       ConfigParm `yaml:"parm"`
	Statistics ConfigParm `yaml:"statistics"`
}

type ConfigParm struct {
	Server   string `yaml:"server" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	Port     int    `yaml:"port" env-required:"true"`
	NameDB   string `yaml:"namedb" env-required:"true"`
}

func MustLoad() (*Config, *ConfigDB) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set ")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	configDBPath := os.Getenv("CONFIG_PATH_DB")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set ")
	}

	if _, err := os.Stat(configDBPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configDBPath)
	}

	var cfgdb ConfigDB

	if err := cleanenv.ReadConfig(configDBPath, &cfgdb); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg, &cfgdb

}
