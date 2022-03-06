package main

import (
	"fmt"
)

const (
	url = "https://mstdn.jp/api/v1/timelines/public/?local=true"
)

func main() {
	jsonBytes, err := getRespBody(url)
	if err != nil {
		return
	}

	data, err := Parse(jsonBytes)
	if err != nil {
		return
	}

	fmt.Printf("%#+v\n", *data)
}
