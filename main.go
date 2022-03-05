package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	url = "https://mstdn.jp/api/v1/timelines/public/?local=true"
)

// json をパースする際の定義
type Status struct {
	Content string  `json:"Content"`
	Account Account `json:"Account"`
}

type Account struct {
	Username string `json:"Username"`
}

func main() {
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
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	// 返ってきた response の Body を取得
	byteArray, _ := ioutil.ReadAll(resp.Body)
	jsonBytes := ([]byte)(string(byteArray))

	// パース規則の変数
	data := new([]Status)

	// パース
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal erroe: ", err)
		return
	}
	fmt.Printf("%#+v\n", *data)
}
