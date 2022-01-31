package config

import (
	"log"
	"strconv"
	"time"
)

type AppConfig interface {
	GetHost() string
	GetUsername() string
	GetPassword() string
	GetDialTimeout() int
	GetReadTimeout() int
	GetWriteTimeout() int
	GetDb() string
}

func NewAppConfig() AppConfig {
	return &AppConfigImpl{}
}

type AppConfigImpl struct {
	Host         string `mapstructure:"Host"`
	Username     string `mapstructure:"Username"`
	Password     string `mapstructure:"Password"`
	DialTimeout  string `mapstructure:"DialTimeout"`
	ReadTimeout  string `mapstructure:"ReadTimeout"`
	WriteTimeout string `mapstructure:"WriteTimeout"`
	Db           string `mapstructure:"Db"`
}

func (c *AppConfigImpl) GetHost() string {
	return c.Host
}

func (c *AppConfigImpl) GetUsername() string {
	return c.Username
}
func (c *AppConfigImpl) GetPassword() string {
	return c.Password
}
func (c *AppConfigImpl) GetDialTimeout() int {
	val, err := strconv.Atoi(c.DialTimeout)
	if err != nil {
		log.Fatal("failed loading DialTimeout value from env file")
	}
	return val * int(time.Second)
}
func (c *AppConfigImpl) GetReadTimeout() int {
	val, err := strconv.Atoi(c.ReadTimeout)
	if err != nil {
		log.Fatal("failed loading ReadTimeout value from env file")
	}
	return val * int(time.Second)
}
func (c *AppConfigImpl) GetWriteTimeout() int {
	val, err := strconv.Atoi(c.WriteTimeout)
	if err != nil {
		log.Fatal("failed loading WriteTimeout value from env file")
	}
	return val * int(time.Second)
}

func (c *AppConfigImpl) GetDb() string {
	return c.Db
}
