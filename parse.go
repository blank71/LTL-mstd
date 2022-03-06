package main

import (
	"encoding/json"
	"fmt"
)

// json をパースする際の定義
type Status struct {
	Content string  `json:"Content"`
	Account Account `json:"Account"`
}

type Account struct {
	Username string `json:"Username"`
}

func Parse(jsonBytes []byte) (data *[]Status, err error) {
	// パース
	data = new([]Status)
	err = json.Unmarshal(jsonBytes, data)
	if err != nil {
		fmt.Println("JSON Unmarshal erroe: ", err)
		return
	}
	return
}
