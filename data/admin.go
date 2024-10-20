package data

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Admin struct {
	appID     string
	appSecret string
}

type AConfig struct {
	AppID     string `yaml:"appID"`
	AppSecret string `yaml:"appSecret"`
}

type AdminConfig struct {
	Admins []AConfig `yaml:"admins"`
}

func NewAdmin(appID, appSecret string) Admin {
	return Admin{
		appID:     appID,
		appSecret: appSecret,
	}
}

type AdminInterface interface {
	GetAppID() string
	GetAppSecret() string
}

func (a Admin) GetAppID() string {
	return a.appID
}

func (a Admin) GetAppSecret() string {
	return a.appSecret
}

func CreateAdmin(filename string) Admin {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}
	var config AdminConfig
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %v", err)
	}

	return NewAdmin(config.Admins[0].AppID, config.Admins[0].AppSecret)
}
