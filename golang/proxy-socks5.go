package main

import (
	"fmt"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"time"
)

func main() {
	username := "username"
	password := "password"
	host := "gw-us.nstproxy.com"
	port := 24125
	proxyAddress := fmt.Sprintf("%s:%v", host, port)

	auth := proxy.Auth{
		User:     username,
		Password: password,
	}

	dialer, err := proxy.SOCKS5("tcp", proxyAddress, &auth, proxy.Direct)
	if err != nil {
		panic(err)
	}

	httpTransport := &http.Transport{
		Dial: dialer.Dial,
	}

	client := &http.Client{Transport: httpTransport, Timeout: time.Second * 10}

	resp, err := client.Get("API_URL")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
}
