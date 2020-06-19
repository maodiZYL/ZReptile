package main

import (
	"crypto/tls"
	"fmt"
	"github.com/Unknwon/goconfig"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
)

func main() {
	Username, Password, Address, Port := conf()
	url := fmt.Sprintf("https://httpbin.org/get")
	//创建代理
	auth := proxy.Auth{
		User:     Username,
		Password: Password,
	}
	address := fmt.Sprintf("%s:%s", Address, Port)
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
func conf() (string, string, string, string) {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		panic("错误1")
	}
	Username, err := cfg.GetValue("data", "Username")
	Password, err := cfg.GetValue("data", "Password")
	Address, err := cfg.GetValue("data", "Address")
	Port, err := cfg.GetValue("data", "Port")
	if err != nil {
		panic("错误2")
	}
	return Username, Password, Address, Port
}
