package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unicode/utf8"
)

type Hitokoto struct {
	ID         int     `json:"id"`
	UUID       string  `json:"uuid"`
	Hitokoto   string  `json:"hitokoto"` // 一言文字内容
	Type       string  `json:"type"`
	From       string  `json:"from"`
	FromWho    *string `json:"from_who"` // 可能为null，使用指针
	Creator    string  `json:"creator"`
	CreatorUID int     `json:"creator_uid"`
	Reviewer   int     `json:"reviewer"`
	CommitFrom string  `json:"commit_from"`
	CreatedAt  string  `json:"created_at"` // 时间戳
	Length     int     `json:"length"`
}

// GetHitokoto 获取一言
func GetHitokoto() (string, error) {
	url := "https://v1.hitokoto.cn"
	// 发起 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 将 JSON 数据解析到结构体
	var hitokoto Hitokoto
	err = json.Unmarshal(body, &hitokoto)
	if err != nil {
		return "", err
	}

	return hitokoto.Hitokoto, nil
}

func GetDailyLove() (string, error) {
	url := "https://api.lovelive.tools/api/SweetNothings/Serialization/Json"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	if sentence, ok := result["returnObj"].([]interface{})[0].(string); ok {
		return sentence, nil
	}
	return "", fmt.Errorf("unable to get daily love")
}

func SplitSentence(sentence string, maxLength int) []string {
	var parts []string
	n := len(sentence)

	// 如果句子长度小于或等于最大长度，直接返回
	if n <= maxLength {
		return []string{sentence}
	}

	start := 0
	for i := 0; i < 4; i++ {
		end := start
		count := 0

		// 计算当前部分的字符数
		for end < n && count < maxLength {
			_, size := utf8.DecodeRuneInString(sentence[end:])
			end += size
			count++
		}

		parts = append(parts, sentence[start:end])
		start = end

		if start >= n { // 防止多余的循环
			break
		}
	}

	return parts
}
