package main

import (
	"io/ioutil"
	"net/http"
)

func getResp(url string) (resp *http.Response, err error) {
	// server に投げる request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	// Header を追加しないと 200 が返って来ない
	req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)

	// req を server に投げて response を得る
	client := new(http.Client)
	resp, err = client.Do(req)
	return
}

func getRespBody(url string) (jsonBytes []byte, err error) {
	resp, err := getResp(url)
	if err != nil {
		return
	}

	// 返ってきた response の Body を取得
	byteArray, err := ioutil.ReadAll(resp.Body)
	jsonBytes = ([]byte)(string(byteArray))
	return
}
