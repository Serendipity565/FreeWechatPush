package sever

import (
	"FreeWechatPush/data"
	"FreeWechatPush/service/weather"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func timeDifference(birthday time.Time) string {
	cst := time.FixedZone("CST", 8*3600)
	today := time.Now().In(cst)
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, cst)
	if today.Equal(birthday) {
		return "今天就是你的生日啦!生日快乐!"
	}
	if today.After(birthday) {
		birthday = time.Date(today.Year()+1, birthday.Month(), birthday.Day(), 0, 0, 0, 0, cst)
	}
	delta := birthday.Sub(today)
	return fmt.Sprintf("还有 %d 天", int(delta.Hours()/24))
}

func SendWeather(user data.User, accessToken string) error {
	cst := time.FixedZone("CST", 8*3600)
	today := time.Now().In(cst).Format("2006年01月02日")

	//sentence, err := weather.GetDailyLove()
	sentence, err := weather.GetHitokoto()

	if err != nil {
		return err
	}
	birthdaySentence := timeDifference(user.GetBirthday())
	weatherDetails, err := weather.GetWeather(user.GetPlace())
	if err != nil {
		return err
	}
	if len(weatherDetails) == 0 {
		return fmt.Errorf("filed to get find city")
	}
	tempRange := weatherDetails["lowTemp"] + " — " + weatherDetails["highTemp"]
	maxLength := 15

	// sentence = "忽然想问你一些秋天的事情，比如桂花的花期是多少天，天气冷了要添几件长袖，还有，能不能趁着秋天喜欢我"

	parts := weather.SplitSentence(sentence, maxLength)

	if len(parts) < 4 {
		for len(parts) < 4 {
			parts = append(parts, "")
		}
	}

	note := parts[0]
	note1 := parts[1]
	note2 := parts[2]
	note3 := parts[3]

	body := map[string]interface{}{
		"touser":      user.GetOpenid(),
		"template_id": data.WeatherTempleId,
		"url":         "https://weixin.qq.com",
		"data": map[string]map[string]string{
			"date": {
				"value": today,
			},
			"region": {
				"value": user.GetPlace(),
			},
			"weather": {
				"value": weatherDetails["weather"],
			},
			"temp": {
				"value": tempRange,
			},
			"wind_dir": {
				"value": weatherDetails["wind_dir"],
			},
			"wind_str": {
				"value": weatherDetails["wind_str"],
			},
			"birthday": {
				"value": birthdaySentence,
			},
			"note": {
				"value": note,
			},
			"note1": {
				"value": note1,
			},
			"note2": {
				"value": note2,
			},
			"note3": {
				"value": note3,
			},
		},
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", accessToken)
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", strings.NewReader(string(jsonBody)))
	if err != nil {
		return err
	}

	// 用处查看发送的信息
	func() {
		fmt.Println(user.GetName())
		for key, value := range weatherDetails {
			fmt.Printf("%s : %s ", key, value)
		}
		fmt.Println()
		fmt.Println(sentence)
		fmt.Println(birthdaySentence)
		fmt.Println()
	}()

	defer resp.Body.Close()
	return nil
}
