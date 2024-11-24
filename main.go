package main

import (
	"FreeWechatPush/data"
	"FreeWechatPush/service/login"
	"FreeWechatPush/sever"
	"fmt"
	"time"
)

var (
	userFile  = "users.yaml"
	adminFile = "admin.yaml"
)

func solve() {
	var admin data.Admin
	admin = data.CreateAdmin(adminFile)
	token, err := login.GetAccessToken(admin)
	if err != nil {
		fmt.Println(err)
		return
	}

	users := data.ReadFromConfig(userFile)
	for _, user := range users {
		err = sever.SendWeather(user, token)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func waitUntilEight() {
	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, now.Location())
	if now.After(next) {
		next = next.Add(24 * time.Hour)
	}
	time.Sleep(time.Until(next))
}

func main() {
	for {
		waitUntilEight()
		solve()
	}
}
