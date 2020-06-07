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
	request := "POST"
	url := "https://higumeng.cn/memberAnchor/memberAnchors"
	payload := strings.NewReader(fmt.Sprintf("pageIndex=%v&type=%v&channel=qita", pageIndex, typechange)) //fmt.Sprintf()才能用%v
	mapNum := make(map[string]string)
	mapNum["content-length"] = "31"
	mapNum["Host"] = "<calculated when request is sent>"
	mapNum["content-type"] = "application/x-www-form-urlencoded"
	mapNum["user-agent"] = "gzip"
	body := Agent(request, url, payload, mapNum)
	var kitty SignalCommunicationID //应用结构体
	json.Unmarshal(body, &kitty)    //将数据存放到结构体中
	//fmt.Println(kitty)          //控制台输出mao
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

func Agent(request, url string, l *strings.Reader, to map[string]string) []byte {

	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

	req, _ := http.NewRequest(request, url, l) //开始请求
	for key, value := range to {
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
