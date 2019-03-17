package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.ptt.cc/bbs/movie/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile("index.html", b, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
