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
	url := "http://open2.gamemm.com/index/searchUserList"
	Climb_SignalommunicationID("POST", url, &postData)
}

//爬取私信Id
func Climb_SignalommunicationID(method, url string, postData *map[string]string) {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range *postData {
		w.WriteField(k, v)
	}
	w.Close()

	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Content-Type", w.FormDataContentType())
	req.Header.Add("Host", "open2.gamemm.com")
	//req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("User-Agent", "okhttp/3.11.0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Length", "1190")

	//使用代理
	var resp *http.Response
	httpTransport := &http.Transport{ //跳过证书验证
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: httpTransport}
	if dialer != nil {
		httpTransport.Dial = dialer.Dial
	}

	resp, _ = httpClient.Do(req)
	data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var kiss SignalCommunicationID
	json.Unmarshal(data, &kiss)
	fmt.Println(string(data))
	//fmt.Println(kiss)
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
