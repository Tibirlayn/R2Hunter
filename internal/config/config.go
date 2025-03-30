package config

import (
	"log"
	"os"
	"time"

	 //"path/filepath"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	HTTPServer  HTTPServer    `yaml:"http_server"`
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
	// Получаем путь к исполняемому файлу
/* 	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Не удалось получить путь к исполняемому файлу: %v", err)
	}

	// Путь к файлу config.yaml в той же директории, что и .exe
	configPath := filepath.Join(filepath.Dir(exePath), "local.yaml")
	configDBPath := filepath.Join(filepath.Dir(exePath), "config.yaml")

	// Проверка существования config.yaml
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Файл конфигурации не найден: %s", configPath)
	}

	// Загрузка основного конфигурационного файла
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}

	// Проверка существования config_db.yaml
	if _, err := os.Stat(configDBPath); os.IsNotExist(err) {
		log.Fatalf("Файл конфигурации БД не найден: %s", configDBPath)
	}

	// Загрузка конфигурации базы данных
	var cfgdb ConfigDB
	if err := cleanenv.ReadConfig(configDBPath, &cfgdb); err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию БД: %v", err)
	} */


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

/* 	// Получаем путь к исполняемому файлу
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Не удалось получить путь к исполняемому файлу: %v", err)
	}
	
	// Путь к файлу config.yaml в той же директории, что и .exe
	configDBPath := filepath.Join(filepath.Dir(exePath), "config.yaml")

	// Проверка существования config.yaml
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Файл конфигурации не найден: %s", configPath)
	} */

	var cfgdb ConfigDB

	if err := cleanenv.ReadConfig(configDBPath, &cfgdb); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}



	return &cfg, &cfgdb
}

// TODO: parm.yaml 
