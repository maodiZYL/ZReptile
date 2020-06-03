package main

import (
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导数据库
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
)

type Wutuo struct {
	Members []struct {
		No       string `json:"no"`
		Nickname string `json:"nickname"`
	} `json:"members"`
}

func main() {
	for i := 1; i < 284; i++ {
		url := fmt.Sprintf("https://app.quanziapp.com/api/v2/spjZndj/members?&page=%v&per_page=100", i)
		//创建代理
		auth := proxy.Auth{
			User:     "itemb123",
			Password: "kIl8Jl3aKej",
		}
		address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
		dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

		req, _ := http.NewRequest("GET", url, nil) //开始请求
		req.Header.Set("Host", "<calculated when request is sent>")
		req.Header.Set("x-app-version", "Android Circles 3.5.3")
		req.Header.Set("authorization", "token Vnn9DLoPTnscLgoMRcG9eNXT1590739356.5220678")

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
		var tem Wutuo                        //用结构体
		json.Unmarshal(body, &tem)           //将查到的数据放到结构体中
		//fmt.Println(string(body))    //打印body内数据
		//save1(tem)
		WutuoMql(tem)
		//fmt.Println(tem)

	}
}
func WutuoMql(wu Wutuo) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	stmt, _ := db.Prepare("insert wutuo(n_id ,Name) values(?,?)")
	for i := 0; i < len(wu.Members); i++ {
		est, _ := stmt.Exec(wu.Members[i].No, wu.Members[i].Nickname)
		fmt.Println(est)
	}
	db.Close()
}
