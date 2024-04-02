package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	username := "username"
	password := "password"
	host := "gw-us.nstproxy.com"
	port := 24125
	proxyURL := fmt.Sprintf("http://%s:%s@%s:%v", username, password, host, port)
	proxy, _ := url.Parse(proxyURL)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}

	req, err := http.NewRequest("GET", "API_URL", nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
