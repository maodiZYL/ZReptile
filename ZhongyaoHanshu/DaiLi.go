package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
)

func main() {
	url := fmt.Sprintf("https://httpbin.org/get")
	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

	req, _ := http.NewRequest("GET", url, nil) //开始请求

	//使用代理
	var resp *http.Response
	httpTransport := &http.Transport{ //跳过证书验证
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: httpTransport}
	if dialer != nil {
		httpTransport.Dial = dialer.Dial
	}

	resp, _ = httpClient.Do(req)
	//resp, _ := client.Do(req)   //处理请求
	body, _ := ioutil.ReadAll(resp.Body) //读取响应
	fmt.Println(string(body))

}
