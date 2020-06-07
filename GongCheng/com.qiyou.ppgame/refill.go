// com.qiyou.ppgame 包，该包是笔芯语音爬虫。
//创建人：zhuyelu
//创建时间：20200605

package main

import (
	"crypto/md5"
	_ "crypto/md5"
	"crypto/tls"
	"database/sql"
	"encoding/hex"
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

//通信id
type SignalCommunicationID struct {
	Data []struct {
		Userid int `json:"userid"`
	} `json:"data"`
}

func main() {
	for i := 1; i < 100000; i++ {
		Climb_SignalommunicationID(i)
	}

}

//爬通信id
func Climb_SignalommunicationID(pageid int) {
	sign := MD5_Encryption(pageid)
	request := "GET"
	url := fmt.Sprintf("http://apk.qqi2019.com:8001/Api/sound_call/soundcall.aspx?pageid=%v&userid=73002494&sign=%v", pageid, sign)
	mapNum := make(map[string]string) //用map储存键值对信息
	mapNum["Host"] = "apk.qqi2019.com:8001"
	mapNum["User-Agent"] = "okhttp/3.12.1"
	mapNum["Connection"] = "keep-alive"
	body := Agent(request, url, nil, mapNum)
	var tem SignalCommunicationID //用结构体
	json.Unmarshal(body, &tem)    //将查到的数据放到结构体中
	//fmt.Println(string(body))    //打印body内数据
	MysqlRefill(tem)
	time.Sleep(time.Second * 10) //设置时间
}

//MD5加密
func MD5_Encryption(pageid int) (sign string) {
	urlStr := fmt.Sprintf("pageid=%v&userid=73002494&key=kjldsnjdvndkfgsdfnsdnb", pageid)
	//n2 := url.QueryEscape(urlStr)
	h := md5.New()
	h.Write([]byte(urlStr)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	n := hex.EncodeToString(cipherStr) //加密
	return n
}

//将数据插入blogdb数据库中refill
func MysqlRefill(refill SignalCommunicationID) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	stmt, _ := db.Prepare("insert refill (n_id) values(?)")
	for i := 0; i < len(refill.Data); i++ {
		est, _ := stmt.Exec(refill.Data[i].Userid)
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
