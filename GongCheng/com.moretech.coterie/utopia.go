// com.moretech.coterie 包，该包是乌托邦爬虫。
//创建人：zhuyelu
//创建时间：20200605

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
	"strings"
)

//用户id
type UserID struct {
	Members []struct {
		ID string `json:"id"`
	} `json:"members"`
}

//通信id
type SignalCommunicationID struct {
	ImC2CGroupID string `json:"im_c2c_group_id"`
}

//主方法
func main() {
	for i := 1; i < 284; i++ {
		ClimbCUserID()
	}
}

//爬出用户id
func ClimbCUserID() {

	state_id := "sBAKAgj"
	request := "GET"
	url := fmt.Sprintf("https://app.quanziapp.com/api/v2/%v/members?&page=1&per_page=100", state_id)
	mapNum := make(map[string]string)
	mapNum["Host"] = "<calculated when request is sent>"
	mapNum["x-app-version"] = "Android Circles 3.5.3"
	mapNum["authorization"] = "token Vnn9DLoPTnscLgoMRcG9eNXT1590739356.5220678"
	body := Agent(request, url, nil, mapNum)
	var state UserID //用结构体
	json.Unmarshal(body, &state)
	//将查到的数据放到结构体中
	for i := 0; i < len(state.Members); i++ {
		userid := state.Members[i].ID
		Climb_SignalommunicationID(state_id, userid)
	}

}

//爬出通信id
func Climb_SignalommunicationID(identifier string, member_id string) { //bang为结构体，接收的是结构体

	request := "POST"
	url := fmt.Sprintf("https://app.quanziapp.com/api/v2/im/init_c2c_group_contact")
	payload := strings.NewReader(fmt.Sprintf(`{"identifier":"%v","member_id":"%v"}`, identifier, member_id)) //传application/json; charset=utf-8
	mapNum := make(map[string]string)                                                                        //用map储存键值对信息
	mapNum["Host"] = "<calculated when request is sent>"
	mapNum["authorization"] = "token Vnn9DLoPTnscLgoMRcG9eNXT1590739356.5220678"
	mapNum["Content-Length"] = "<calculated when request is sent>"
	mapNum["x-app-version"] = "Android Circles 3.5.3"
	mapNum["Content-Type"] = "application/json; charset=utf-8"
	mapNum["accept-language"] = "zh"
	mapNum["user-agent"] = "okhttp/3.14.4"
	body := Agent(request, url, payload, mapNum)
	var tem SignalCommunicationID //用结构体
	json.Unmarshal(body, &tem)    //将查到的数据放到结构体中
	fmt.Println(string(body))
	MysqUtopia(tem)

}

//将数据插入blogdb数据库中utopia
func MysqUtopia(utopia SignalCommunicationID) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8") //链接数据库
	stmt, _ := db.Prepare("INSERT utopia (n_id) values (?)")                          //插入语句   字段不能填错
	for i := 0; i < len(utopia.ImC2CGroupID); i++ {                                   //循环插入
		shuju, _ := stmt.Exec(utopia.ImC2CGroupID) //执行数据存储
		fmt.Println(shuju)
	}
	db.Close()
}

//代理
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
