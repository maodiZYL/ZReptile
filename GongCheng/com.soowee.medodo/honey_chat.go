//com.soowee.medodo 包，该包是蜜聊爬虫。
//创建人:zhuyelu
//创建时间:20200609

package main

import (
	_ "crypto/md5"
	"crypto/tls"
	"database/sql"
	_ "encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
	_ "net/url"
	"strings"
	"time"
)

//定义储存数据的结构体
type Chat struct {
	Data []struct {
		Userid string `json:"userid"`
	} `json:"data"`
}

func main() {
	for j := 1; j < 10000; j++ {
		time.Sleep(time.Second * 2) //设置时间
		for i := 1; i < 1000; i++ {
			time.Sleep(time.Second * 2) //设置时间
			Climb_SignalommunicationID(i)
			time.Sleep(time.Second * 2) //设置时间
		}
	}
}

//爬通信id
func Climb_SignalommunicationID(page int) {
	url := "https://api.3wee.cn/friend/get_user_list.php"
	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

	payload := strings.NewReader(fmt.Sprintf("XDEBUG_SESSION_START=F7EDF33B7DD3A49B426C65F3C3EF5F08&tab=recomm&page=%v", page)) //fmt.Sprintf()才能用%v
	req, _ := http.NewRequest("POST", url, payload)
	//req.Header.Add("Content-Length", "<calculated when request is sent>")
	req.Header.Add("Host", "<calculated when request is sent>")
	req.Header.Add("x-api-password", "WQA150g9Bs50")
	req.Header.Add("x-api-userid", "5504995")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("content-length", "71")

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
	//resp,_ :=  http.DefaultClient.Do(req)
	//defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	var chat Chat
	json.Unmarshal(body, &chat)
	fmt.Println(chat)
	MysqlHoneyChat(chat)
}

//将数据插入blogdb数据库中refill
func MysqlHoneyChat(chat Chat) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	stmt, _ := db.Prepare("insert chat (n_id) values(?)")
	for i := 0; i < len(chat.Data); i++ {
		est, _ := stmt.Exec(chat.Data[i].Userid)
		fmt.Println(est)
	}
	db.Close()
}

func Agent(request, url string, payload *strings.Reader, mapNum map[string]string) []byte {

	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

	req, _ := http.NewRequest(request, url, payload) //开始请求
	for key, value := range mapNum {
		req.Header.Set(key, value)
	}

	//使用代理
	var resp *http.Response
	httpTransport := &http.Transport{ //跳过证书验证
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: httpTransport}
	if dialer != nil {
		httpTransport.Dial = dialer.Dial
	}

	resp, _ = httpClient.Do(req)         //处理请求
	body, _ := ioutil.ReadAll(resp.Body) //读取响应
	return body

}
