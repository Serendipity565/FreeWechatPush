package data

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type User struct {
	name     string
	birthday time.Time
	place    string
	openid   string
}

type UConfig struct {
	Name   string `yaml:"name"`
	Month  string `yaml:"month"`
	Day    int    `yaml:"day"`
	Place  string `yaml:"place"`
	Openid string `yaml:"openid"`
}

type UserConfig struct {
	Users []UConfig `yaml:"users"`
}

type UserInterface interface {
	GetName() string
	GetBirthday() time.Time
	GetPlace() string
	GetOpenid() string
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetBirthday() time.Time {
	return u.birthday
}

func (u *User) GetPlace() string {
	return u.place
}

func (u *User) GetOpenid() string {
	return u.openid
}

func NewUser(name string, month time.Month, day int, place string, openid string) User {
	today := time.Now()
	year := today.Year()
	return User{
		name:     name,
		birthday: time.Date(year, month, day, 0, 0, 0, 0, today.Location()),
		place:    place,
		openid:   openid,
	}
}

func ReadFromConfig(filename string) []User {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var config UserConfig
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %v", err)
	}

	// 创建用户
	var users []User
	for _, userConfig := range config.Users {
		month, err := time.Parse("January", userConfig.Month)
		if err != nil {
			log.Fatalf("Error parsing birthday: %v", err)
		}
		user := NewUser(userConfig.Name, month.Month(), userConfig.Day, userConfig.Place, userConfig.Openid)
		users = append(users, user)
		// fmt.Println(user)
	}
	return users
}
