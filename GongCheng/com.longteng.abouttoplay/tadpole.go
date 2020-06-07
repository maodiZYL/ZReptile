// com.longteng.abouttoplay 包，该包是蝌蚪语音爬虫。
//创建人：zhuyelu
//创建时间：20200605

package main

import (
	"crypto/md5"
	"crypto/tls"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//通信id
type SignalCommunicationID struct {
	Data []struct {
		UserID string `json:"userId"`
	} `json:"data"`
}

//主方法
func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10000; j++ {
			Climb_SignalommunicationID(i, j)
		}
	}
}

//爬通信id
func Climb_SignalommunicationID(i int, page int) {

	timestamp := int32(time.Now().Unix())   //获取当前时间戳
	sign := MD5_Encryption(page, timestamp) //调用加密方法
	request := "POST"                       //请求方式
	url := "https://yuewowan.yangba.tv/career/queryKeDouIndexList"
	payload := strings.NewReader(fmt.Sprintf("sortProperty=activeData&sign=%v&appkey=show_android&userId=50166072&channelType=a_xiaomi&pageSize=20&timestamp=%v&appType=android&appVersion=3.3.2&area=ALL&session_id=c44e6327-f806-41ec-aac0-d268373e8001&sex=ALL&page=%v", sign, timestamp, page)) //fmt.Sprintf()才能用%v
	shu := 256 + i
	cl := fmt.Sprintf("%v", shu)
	mapNum := make(map[string]string) //用map储存键值对信息
	mapNum["Content-Length"] = cl
	mapNum["Host"] = "yuewowan.yangba.tv"
	mapNum["content-type"] = "application/x-www-form-urlencoded"
	mapNum["Connection"] = "keep-alive"
	body := Agent(request, url, payload, mapNum)
	var tadpole SignalCommunicationID
	json.Unmarshal(body, &tadpole)
	MysqlTadpole(tadpole)
	//fmt.Println(tadpole)
}

//MD5加密
func MD5_Encryption(page int, timestamp int32) (sign string) {
	//time1 := int(time)
	urlStr := fmt.Sprintf("appType=android&appVersion=3.3.2&appkey=show_android&appsecret=b64939eddd94efa1c750f2563868c2b8&area=ALL&channelType=a_xiaomi&page=%v&pageSize=20&session_id=c44e6327-f806-41ec-aac0-d268373e8001&sex=ALL&sortProperty=activeData&timestamp=%v&userId=50166072", page, timestamp)
	n2 := url.QueryEscape(urlStr) //url加密
	h := md5.New()
	h.Write([]byte(n2)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	n := hex.EncodeToString(cipherStr) //加密
	return n
}

//将数据插入blogdb数据库中tadpole
func MysqlTadpole(tadpole SignalCommunicationID) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	stmt, _ := db.Prepare("insert tadpole (n_id) values(?)")
	for i := 0; i < len(tadpole.Data); i++ {
		est, _ := stmt.Exec(tadpole.Data[i].UserID)
		fmt.Println(est)
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
