package main

import (
	"bytes"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
)

type tongxing3 struct {
	ImC2CGroupID string `json:"im_c2c_group_id"`
}

func main() {
	url := fmt.Sprintf("https://app.quanziapp.com/api/v2/im/init_c2c_group_contact")

	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

	//sBAKAgj
	//jRMr01
	requestBody := fmt.Sprintf(`{"identifier":"spjZMl5","member_id":"5Vl8Pk"}`) //传application/json; charset=utf-8
	var jsonStr = []byte(requestBody)

	//payload := strings.NewReader(fmt.Sprintf("befollow_id=%v&user_id=98236258&is_login=0",bang.Members[i].ID)) //fmt.Sprintf()才能用%v  //传application/x-www-form-urlencoded

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr)) //开始请求
	req.Header.Set("Host", "<calculated when request is sent>")
	req.Header.Set("authorization", "token Vnn9DLoPTnscLgoMRcG9eNXT1590739356.5220678")
	req.Header.Set("Content-Length", "<calculated when request is sent>")
	req.Header.Set("x-app-version", "Android Circles 3.5.3")
	req.Header.Set("content-type", "application/json; charset=utf-8")
	req.Header.Set("accept-language", "zh")
	req.Header.Set("user-agent", "okhttp/3.14.4")

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
	body, _ := ioutil.ReadAll(resp.Body) //读取响应
	var tem tongxing3                    //用结构体
	json.Unmarshal(body, &tem)           //将查到的数据放到结构体中
	//fmt.Println(string(body))
	fmt.Println(string(body))
	Mysql4(tem)
}
func Mysql4(wubang tongxing3) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8") //链接数据库
	stmt, _ := db.Prepare("INSERT wutuo1 (n_id) values (?)")                          //插入语句   字段不能填错
	for i := 0; i < len(wubang.ImC2CGroupID); i++ {                                   //循环插入
		shuju, _ := stmt.Exec((wubang.ImC2CGroupID)) //执行数据存储
		fmt.Println(shuju)
	}
	db.Close()
}
