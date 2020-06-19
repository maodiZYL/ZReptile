// com.miaomi.qinba 包，该包是亲吧爬虫。
//创建人：zhuyelu
//创建时间：20200605

package main

import (
	"bytes"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	//"crypto/tls"
	///	"database/sql"
	//"encoding/json"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	//	"golang.org/x/net/proxy"
	///"io/ioutil"
	//"net/http"
	//"strings"
	//"time"
)

type SignalCommunicationID struct {
	Result struct {
		Users []struct {
			PerfectNumber int `json:"perfect_number"`
		} `json:"users"`
	} `json:"result"`
}

func main() {
	FromData()
}

//FromData传值
func FromData() {

	postData := make(map[string]string)
	postData["user_code"] = "9c535d06fcd37ec6dc576abc73581bc6"
	postData["source"] = "android"
	postData["device_id"] = "883dba42c33f6e900ec1f579ec6fc72f"
	postData["app_style"] = "1"
	postData["version"] = "9"
	postData["ads_idf"] = "d28c935848e22dedfbe0e345acb3ed04"
	postData["channel"] = "oppo"
	postData["keywords"] = "7"
	postData["sign"] = "cc79c19f7a46f644122da76f39e758be"
	Climb_SignalommunicationID(&postData)

}

//爬取私信Id
func Climb_SignalommunicationID(postData *map[string]string) {

	request := "POST"
	url := "http://open2.gamemm.com/index/searchUserList"
	payload := new(bytes.Buffer)
	w := multipart.NewWriter(payload)
	for k, v := range *postData {
		w.WriteField(k, v)
	}

	mapNum := make(map[string]string) //用map储存键值对信息
	mapNum["Content-Type"] = w.FormDataContentType()
	mapNum["Host"] = "open2.gamemm.com"
	mapNum["User-Agent"] = "okhttp/3.11.0"
	mapNum["Connection"] = "keep-alive"
	mapNum["Content-Length"] = "1190"

	body := Agent(request, url, payload, mapNum)
	var kiss SignalCommunicationID
	json.Unmarshal(body, &kiss)
	fmt.Println(string(body))
	MysqlKiss(kiss)

}

//将数据插入blogdb数据库kiss表
func MysqlKiss(kiss SignalCommunicationID) {

	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	stmt, _ := db.Prepare("insert kiss(n_id) values(?)")
	for i := 0; i < len(kiss.Result.Users); i++ {
		est, _ := stmt.Exec(kiss.Result.Users[i].PerfectNumber)
		fmt.Println(est)
	}
	db.Close()

}

//代理
func Agent(request, url string, l *bytes.Buffer, to map[string]string) []byte {

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
