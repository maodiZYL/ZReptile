// com.szhibo.ulive 包，该包是n号房爬虫。
//创建人：zhuyelu
//创建时间：20200607

package main

import (
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type SignalCommunicationID struct {
	Data struct {
		Data []struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	} `json:"data"`
}

func main() {
	for i := 1; i < 1000000; i++ {
		time.Sleep(time.Second * 1) //设置等待时间
		Climb_SignalommunicationID()
		time.Sleep(time.Second * 1)
	}
}

func Climb_SignalommunicationID() { //bang为结构体，接收的是结构体

	request := "POSt"
	url := fmt.Sprintf("http://alicdn.wangran.live/api/home/find")
	payload := strings.NewReader(fmt.Sprintf(`{"page":"1","type":"1","package_type":"12","sign":"d55863b18dd2dec83cb0e922498e6187","time":"1591498073659","user_id":"1504151"}`)) //fmt.Sprintf()才能用%v  //传application/x-www-form-urlencoded   传application/json; charset=utf-8

	mapNum := make(map[string]string)
	mapNum["Host"] = "alicdn.wangran.live"
	mapNum["Content-Type"] = "application/json; charset=UTF-8"
	mapNum["Content-Length"] = "128"

	body := Agent(request, url, payload, mapNum)
	fmt.Println(string(body))
	var room SignalCommunicationID
	json.Unmarshal(body, &room)
	MysqlRoom_N(room)
}

//将数据插入blogdb数据库kiss表
func MysqlRoom_N(room SignalCommunicationID) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	stmt, _ := db.Prepare("insert room(n_id) values(?)")
	for i := 0; i < len(room.Data.Data); i++ {
		est, _ := stmt.Exec(room.Data.Data[i].UserID)
		fmt.Println(est)
	}
	db.Close()
}

func Agent(request string, url string, payload *strings.Reader, mapNum map[string]string) []byte {
	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

	req, _ := http.NewRequest(request, url, payload) //开始请求
	for key, Value := range mapNum {
		req.Header.Add(key, Value)
	}

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

	return body
}
