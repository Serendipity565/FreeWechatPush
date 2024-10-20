package weather

import (
	"FreeWechatPush/data"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetWeather(myCity string) (map[string]string, error) {
	url := fmt.Sprintf("https://www.weather.com.cn/weather/%s.shtml", data.CityCode[myCity])
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	text := string(body)
	//fmt.Println("text:", text)
	// 提取天气信息
	weatherInfo := make(map[string]string)

	// 提取日期
	startIdx := strings.Index(text, "<h1>")
	if startIdx != -1 {
		endIdx := strings.Index(text[startIdx:], "（今天）</h1>")
		if endIdx != -1 {
			date := strings.TrimSpace(text[startIdx+4 : startIdx+endIdx])
			weatherInfo["date"] = date
		}
	}

	// 提取天气类型
	startIdx = strings.Index(text, "<p title=")
	if startIdx != -1 {
		endIdx := strings.Index(text[startIdx:], " class=")
		if endIdx != -1 {
			weather := strings.TrimSpace(text[startIdx+10 : startIdx+endIdx-1])
			weatherInfo["weather"] = weather
		}
	}

	// 提取温度信息
	tempStart := strings.Index(text, "<p class=\"tem\">")
	if tempStart != -1 {
		tempEnd := strings.Index(text[tempStart:], "</p>")
		if tempEnd != -1 {
			tempSection := text[tempStart : tempStart+tempEnd]
			lowTempStart := strings.Index(tempSection, "<i>")
			lowTempEnd := strings.Index(tempSection, "</i>")
			if lowTempStart != -1 && lowTempEnd != -1 {
				lowTemp := strings.TrimSpace(tempSection[lowTempStart+3 : lowTempEnd])
				weatherInfo["lowTemp"] = lowTemp
				weatherInfo["highTemp"] = lowTemp

			}

			highTempStart := strings.Index(tempSection, "<span>")
			highTempEnd := strings.Index(tempSection, "</span>")
			if highTempStart != -1 && highTempEnd != -1 {
				highTemp := strings.TrimSpace(tempSection[highTempStart+6 : highTempEnd])
				weatherInfo["highTemp"] = highTemp
			}
		}
	}

	// 获取风向
	// 提取风力信息
	windStart := strings.Index(text, "<p class=\"win\">")
	if windStart != -1 {
		windEnd := strings.Index(text[windStart:], "</p>")
		if windEnd != -1 {
			windSection := text[windStart : windStart+windEnd]
			// fmt.Println(windSection)
			windDirectionStart := strings.Index(windSection, `<span title="`)
			windDirectionEnd := strings.Index(windSection, `" class=`)
			if windDirectionStart != -1 && windDirectionEnd != -1 {
				windDirection := strings.TrimSpace(windSection[windDirectionStart+13 : windDirectionEnd])
				weatherInfo["wind_dir"] = windDirection
			}
			windStrengthStart := strings.Index(windSection, "<i>")
			windStrengthEnd := strings.Index(windSection, "</i>")
			if windStrengthStart != -1 && windStrengthEnd != -1 {
				wind := strings.TrimSpace(windSection[windStrengthStart+3 : windStrengthEnd])
				weatherInfo["wind_str"] = wind
			}
		}
	}
	return weatherInfo, nil
}
