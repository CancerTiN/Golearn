package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://baidu.com"
	url = "https://portal.gdc.cancer.gov/files/9d001823-ed72-4afe-9fa8-cc8e387dc04a"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)

	}
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(s))
	_ = resp.Body.Close()
}
