package service

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

// 不知道为什么用不了，每次都是一同个句子
// 存在反爬？

func GetHitokoto() (string, error) {
	url := "https://hitokoto.cn/"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	var sentence, author string
	var sentences []string

	doc.Find("div.word#hitokoto_text").Each(func(i int, s *goquery.Selection) {
		sentence = s.Text()
		sentences = append(sentences, sentence)
	})

	doc.Find("div.author#hitokoto_author").Each(func(i int, s *goquery.Selection) {
		author = s.Text()
	})

	// 合并所有句子
	fullText := strings.Join(sentences, " ")

	return fullText + " - " + author, nil
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
