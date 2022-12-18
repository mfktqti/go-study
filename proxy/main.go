package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/lunny/log"
)

func main() {
	proxyAddr := "http://218.59.139.238:80"
	httpUrl := "http://www.baidu.com"
	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		log.Fatal(err)
	}
	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(50),
	}
	httpClient := &http.Client{
		Timeout:   time.Second * 100,
		Transport: netTransport,
	}
	res, err := httpClient.Get(httpUrl)
	fmt.Printf("err: %v\n", err)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("res.StatusCode: %v\n", res.StatusCode)
	if res.StatusCode != http.StatusOK {
		log.Println(err)
		return
	}
	c, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(c))
}
