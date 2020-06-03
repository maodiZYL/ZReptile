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

type Kedo struct {
	Data []struct {
		UserID string `json:"userId"`
	} `json:"data"`
}

func main() {
	Pkedo()
}

func Pkedo() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10000; j++ {
			//time.Sleep(time.Second*10)    //设置时间
			//x:=2
			timestamp := int32(time.Now().Unix()) //获取当前时间戳
			sign := Md5JM(j, timestamp)

			//创建代理
			auth := proxy.Auth{
				User:     "itemb123",
				Password: "kIl8Jl3aKej",
			}
			address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
			dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

			url := "https://yuewowan.yangba.tv/career/queryKeDouIndexList"
			payload := strings.NewReader(fmt.Sprintf("sortProperty=activeData&sign=%v&appkey=show_android&userId=50166072&channelType=a_xiaomi&pageSize=20&timestamp=%v&appType=android&appVersion=3.3.2&area=ALL&session_id=c44e6327-f806-41ec-aac0-d268373e8001&sex=ALL&page=%v", sign, timestamp, j)) //fmt.Sprintf()才能用%v
			req, _ := http.NewRequest("POST", url, payload)
			shu := 256 + i
			cl := fmt.Sprintf("%v", shu)
			req.Header.Add("Content-Length", cl)
			req.Header.Add("Host", "yuewowan.yangba.tv")
			req.Header.Add("content-type", "application/x-www-form-urlencoded")
			req.Header.Add("Connection", "keep-alive")

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

			//resp, _ := client.Do(req)   //处理请求
			//body, _:=ioutil.ReadAll(resp.Body)    //读取响应
			//fmt.Println( string(body))
			//resp,_ :=  http.DefaultClient.Do(req)
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			var ke Kedo
			json.Unmarshal(body, &ke)
			//fmt.Println(string(body))
			//fmt.Println(zhu)
			kedoMysql(ke)
			fmt.Println(ke)
		}
		//time.Sleep(time.Second*2)    //设置时间
	}
}

func Md5JM(ys int, time int32) (sign string) {
	time1 := int(time)
	urlStr := fmt.Sprintf("appType=android&appVersion=3.3.2&appkey=show_android&appsecret=b64939eddd94efa1c750f2563868c2b8&area=ALL&channelType=a_xiaomi&page=%v&pageSize=20&session_id=c44e6327-f806-41ec-aac0-d268373e8001&sex=ALL&sortProperty=activeData&timestamp=%v&userId=50166072", ys, time1)
	n2 := url.QueryEscape(urlStr)
	h := md5.New()
	h.Write([]byte(n2)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	n := hex.EncodeToString(cipherStr) //加密
	return n
}

func kedoMysql(ke Kedo) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	stmt, _ := db.Prepare("insert kedo (n_id) values(?)")
	for i := 0; i < len(ke.Data); i++ {
		est, _ := stmt.Exec(ke.Data[i].UserID)
		n := est
		fmt.Println(n)
	}
	db.Close()
}
