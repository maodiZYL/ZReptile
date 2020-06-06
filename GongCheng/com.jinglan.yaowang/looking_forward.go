// com.jinglan.yaowang 包，该包是遥望爬虫。
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
	_ "golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
	"strings"
)

//用户id
type UserID struct {
	Data struct {
		NewFans []struct {
			BefollowID int `json:"befollow_id"`
		} `json:"newFans"`
	} `json:"data"`
}

//通信id
type SignalCommunicationID struct {
	Data struct {
		HuanxinAccount string `json:"huanxin_account"`
	} `json:"data"`
}

//主方法
func main() {
	for i := 1; i < 784; i++ {
		body := ClimbCUserID(i)
		var tem UserID                  //用结构体
		json.Unmarshal(body, &tem)      //将查到的数据放到结构体中
		Climb_SignalommunicationID(tem) //传结构体
	}

}

//爬出用户id
func ClimbCUserID(i int) []byte {

	url := "https://app.jinglantech.tech/message/newfans"
	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

	payload := strings.NewReader(fmt.Sprintf("reqUserId=68347218&pageNo=%v&pageSize=60&other_user_id=100004", i)) //fmt.Sprintf()才能用%v
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Length", "<calculated when request is sent>")
	req.Header.Add("Host", "<calculated when request is sent>")
	req.Header.Add("user_id", "68347218")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("token", "QFUMhufACrpSRP9JbWZhsQ==")

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
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	return body

}

//爬出通信id
func Climb_SignalommunicationID(bang UserID) { //bang为结构体
	for i := 0; i < len(bang.Data.NewFans); i++ { //循环遍历
		url := fmt.Sprintf("https://app.jinglantech.tech/user/findfollowbyid")
		//创建代理
		auth := proxy.Auth{
			User:     "itemb123",
			Password: "kIl8Jl3aKej",
		}
		address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
		dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

		payload := strings.NewReader(fmt.Sprintf("befollow_id=%v&user_id=98236258&is_login=0", bang.Data.NewFans[i].BefollowID)) //fmt.Sprintf()才能用%v

		req, _ := http.NewRequest("POST", url, payload) //开始请求
		req.Header.Set("Host", "<calculated when request is sent>")
		req.Header.Set("user-agent", "okhttp/3.11.0")
		req.Header.Set("content-type", "application/x-www-form-urlencoded")

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
		var tem SignalCommunicationID        //用结构体
		json.Unmarshal(body, &tem)
		MysqlLookingForward(tem)
	}
}

//将数据传到数据库
func MysqlLookingForward(s SignalCommunicationID) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8") //链接数据库
	stmt, _ := db.Prepare("INSERT yaowang (n_id) values (?)")                         //插入语句   字段不能填错
	for i := 0; i < len(s.Data.HuanxinAccount); i++ {                                 //循环插入
		shuju, _ := stmt.Exec(s.Data.HuanxinAccount) //执行数据存储
		fmt.Println(shuju)
	}
	db.Close()
}
