package main

import (
	"FreeWechatPush/data"
	"FreeWechatPush/service/login"
	"FreeWechatPush/sever"
	"fmt"
)

var (
	userFile  = "users.yaml"
	adminFile = "admin.yaml"
)

func main() {
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
			return
		}
		fmt.Println()
		//weatherDetails, err := service.GetWeather(user.GetPlace())
		//if err != nil {
		//	fmt.Println(err)
		//}
		//fmt.Println(weatherDetails["lowTemp"] + " —— " + weatherDetails["highTemp"])
		//fmt.Println(weatherDetails)
	}
}
