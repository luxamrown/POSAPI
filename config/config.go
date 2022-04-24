package config

import (
	"fmt"
	"os"

	"mohamadelabror.com/posapi/manager"
)

type ApiConfig struct {
	Url string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type Manager struct {
	InfraManager   manager.Infra
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
}

type Config struct {
	Manager
	ApiConfig
	DbConfig
}

func (c Config) readConfigFile() Config {

	c.ApiConfig = ApiConfig{Url: "localhost:3000"}
	c.DbConfig = DbConfig{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Name:     os.Getenv("MYSQL_DBNAME"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
	}
	return c
}

func NewConfig() Config {
	cfg := Config{}
	cfg = cfg.readConfigFile()

	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", cfg.DbConfig.User, cfg.DbConfig.Password, cfg.DbConfig.Host, cfg.DbConfig.Port, cfg.DbConfig.Name)
	// dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s", "root", "stauffenberg", "localhost", "3306", "posapi")
	cfg.InfraManager = manager.NewInfra(dataSourceName)
	cfg.RepoManager = manager.NewRepoManager(cfg.InfraManager)
	cfg.UseCaseManager = manager.NewUseCaseManager(cfg.RepoManager)

	return cfg
}
