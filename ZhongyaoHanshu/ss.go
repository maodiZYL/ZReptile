package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
	"strings"
)

type yaowang1 struct {
	Data struct {
		UserName       string `json:"user_name"`
		HuanxinAccount string `json:"huanxin_account"`
	} `json:"data"`
}

func main() {
	url := fmt.Sprintf("https://app.jinglantech.tech/user/findfollowbyid")
	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)
	bangming := 92157839
	payload := strings.NewReader(fmt.Sprintf("befollow_id=%v&user_id=98236258&is_login=0", bangming)) //fmt.Sprintf()才能用%v

	req, _ := http.NewRequest("POST", url, payload) //开始请求
	req.Header.Set("Host", "<calculated when request is sent>")
	req.Header.Set("user-agent", "okhttp/3.11.0")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", "48")
	//使用代理
	var resp *http.Response
	httpTransport := &http.Transport{
		//跳过证书验证
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: httpTransport}
	if dialer != nil {
		httpTransport.Dial = dialer.Dial
	}

	resp, _ = httpClient.Do(req)
	//resp, _ := client.Do(req)   //处理请求
	body, _ := ioutil.ReadAll(resp.Body) //读取响应
	var tem yaowang1                     //用结构体
	json.Unmarshal(body, &tem)           //将查到的数据放到结构体中
	//fmt.Println(string(body))    //打印body内数据
	//save1(tem)
	//Mysqlyao(tem)
	fmt.Println(tem)
}
