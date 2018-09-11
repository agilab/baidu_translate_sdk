package baidu_translate_sdk

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type BaiduTranslator struct {
	appid  string // 百度翻译应用 ID
	appkey string // 百度翻译应用秘钥
}

func CreateBaiduTranslator(
	appid string,
	appkey string) *BaiduTranslator {

	return &BaiduTranslator{
		appid:  appid,
		appkey: appkey,
	}
}

type TranslateResult struct {
	From        string        `json:"from"`
	To          string        `json:"to"`
	TransResult []TransResult `json:"trans_result"`
}

type TransResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

func (tr *BaiduTranslator) Translate(query string) (string, error) {
	t := time.Now()
	salt := t.Format("20060102150405")
	signInput := fmt.Sprintf("%s%s%s%s", tr.appid, query, salt, tr.appkey)
	sign := fmt.Sprintf("%x", md5.Sum([]byte(signInput)))

	url := fmt.Sprintf("http://api.fanyi.baidu.com/api/trans/vip/translate?q=%s&from=en&to=zh&appid=%s&salt=%s&sign=%s", url.QueryEscape(query), tr.appid, salt, sign)

	response, err := http.Get(url)
	if err != nil {
		log.Printf("http get 错误：%s", err)
		return "", err
	}
	returnTxt, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("http read 错误：%s", err)
		return "", err
	}

	var result TranslateResult
	if err := json.Unmarshal(returnTxt, &result); err != nil {
		log.Printf("%s", returnTxt)
		return "", err
	}

	if len(result.TransResult) > 0 {
		return result.TransResult[0].Dst, nil
	}

	return "", errors.New("未知错误")
}
