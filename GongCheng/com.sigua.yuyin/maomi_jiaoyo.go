// com.sigua.yuyin 包，该包是猫咪交友爬虫。
//创建人：zhuyelu
//创建时间：20200605
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

// MaoMi ，用户对象，定义了用户私信ID
type SignalCommunicationID struct {
	Datas []struct {
		MemberID int `json:"memberId"`
	} `json:"datas"`
}

//主方法
func main() {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second * 2) //设置时间
		for j := 1; j <= 30; j++ {
			Climb_SignalommunicationID(i, j)
			time.Sleep(time.Second * 2) //设置时间
		}
	}
}

//爬猫咪交友私信id
func Climb_SignalommunicationID(pageIndex int, typechange int) {

	url := "https://higumeng.cn/memberAnchor/memberAnchors"

	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

	payload := strings.NewReader(fmt.Sprintf("pageIndex=%v&type=%v&channel=qita", pageIndex, typechange)) //fmt.Sprintf()才能用%v
	req, _ := http.NewRequest("POST", url, payload)
	//请求头
	req.Header.Add("content-length", "31")
	req.Header.Add("Host", "<calculated when request is sent>")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "gzip")

	//使用代理
	var resp *http.Response
	httpTransport := &http.Transport{ //跳过证书验证
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: httpTransport}
	if dialer != nil {
		httpTransport.Dial = dialer.Dial
	}
	resp, _ = httpClient.Do(req) //处理请求

	//resp, _ := http.DefaultClient.Do(req)//不用代理的处理请求
	defer resp.Body.Close()              //关闭请求
	body, _ := ioutil.ReadAll(resp.Body) //读取响应
	var kitty SignalCommunicationID      //应用结构体
	json.Unmarshal(body, &kitty)         //将数据存放到结构体中
	//fmt.Println(string(body)) //控制台输出body
	//fmt.Println(mao)          //控制台输出mao
	MsqlKitty(kitty)
}

//将数据插入blogdb数据库的kitty
func MsqlKitty(mao SignalCommunicationID) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	stmt, _ := db.Prepare("insert kitty(n_id ) values(?)")
	for i := 0; i < len(mao.Datas); i++ {
		est, _ := stmt.Exec(mao.Datas[i].MemberID)
		fmt.Println(est)
	}
	db.Close()
}
