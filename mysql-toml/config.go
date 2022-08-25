package main

import (
	"log"

	"github.com/BurntSushi/toml"
	validator "gopkg.in/go-playground/validator.v9"
)

type Config struct {
	Mysql  MysqlConfig  `toml:"mysql"`
	Proxy  ProxyConfig  `toml:"proxy"`
	Worker WorkerConfig `toml:"worker"`
}

type MysqlConfig struct {
	Host     string `toml:"host" validate:"required"`
	Port     int    `toml:"port" validate:"required"`
	Username string `toml:"username" validate:"required"`
	Password string `toml:"password"`
	Schema   string `toml:"schema" validate:"required"`
}

type ProxyConfig struct {
	BaseUrl string `toml:"base_url" validate:"required"`
}

type WorkerConfig struct {
	WorkerNum int `toml:"worker" validate:"required"`
	QueueLen  int `toml:"queue" validate:"required"`
}

var config *Config

// 配置文件config.toml加载
func load_config() {
	var err error
	var validate = validator.New()

	config = &Config{}
	if _, err = toml.DecodeFile("config.toml", config); err != nil {
		log.Fatal(err)
	}

	if err = validate.Struct(config); err != nil {
		log.Fatal(err)
	}
	log.Println("Load config from config.toml---------------------------------------------------------------------")
}
