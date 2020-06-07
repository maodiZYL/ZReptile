// com.tb.linzhu 包，该包是邻住爬虫。
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

//结构体
type SignalCommunicationID struct {
	Res []struct {
		Memberid int `json:"memberid"`
	} `json:"res"`
}

func main() {
	for n := 1; n <= 40; n++ {
		for i := 1; i <= 11; i++ {
			time.Sleep(time.Second * 1) //设置时间
			for j := 1; j <= 27; j++ {
				Climb_SignalommunicationID(n, i, j)
				time.Sleep(time.Second * 1) //设置时间
			}
		}
	}
}

//爬邻住私信Id
func Climb_SignalommunicationID(flag int, page int, catid int) {

	url := "https://app.linzhu.net/app/square/getPosts"
	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

	payload := strings.NewReader(fmt.Sprintf("catid=%v&page=%v&flag=%v", catid, page, flag)) //fmt.Sprintf()才能用%v
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-length", "22")
	req.Header.Add("Host", "<calculated when request is sent>")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "okhttp/3.6.0")

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
	//resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var hood SignalCommunicationID
	json.Unmarshal(body, &hood)
	fmt.Println(string(body))
	fmt.Println(hood)
	MysqlNeighborhood(hood)
}

//将数据插入数据库
func MysqlNeighborhood(hood SignalCommunicationID) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	stmt, _ := db.Prepare("insert linzhu(n_id) values(?)")
	for i := 0; i < len(hood.Res); i++ {
		est, _ := stmt.Exec(hood.Res[i].Memberid)
		fmt.Println(est)
	}
	db.Close()
}
