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
	"time"
)

type Bijgt struct {
	Data []struct {
		Userid int `json:"userid"`
	} `json:"data"`
}

func main() {
	tuijian()

}

func tuijian() {
	for i := 1; i < 100000; i++ {
		s := JM(i)
		//timestamp :=int32(time.Now().Unix())//获取当前时间戳
		//sign := "d49ae6ff5afc3a87be478f2175b409a0"
		url := fmt.Sprintf("http://apk.qqi2019.com:8001/Api/sound_call/soundcall.aspx?pageid=%v&userid=73002494&sign=%v", i, s)
		//url :=  "http://apk.qqi2019.com:8001/Api/sound_call/soundcallcity.aspx?pageid=2&sign=F53A9234989305B458B67D7AAC094AD9&userid=73002494&city=深圳"
		//创建代理
		auth := proxy.Auth{
			User:     "itemb123",
			Password: "kIl8Jl3aKej",
		}
		address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
		dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

		//client := &http.Client{}                   //导入请求的包
		req, _ := http.NewRequest("GET", url, nil) //开始请求
		req.Header.Set("Host", "apk.qqi2019.com:8001")
		req.Header.Set("User-Agent", "okhttp/3.12.1")
		req.Header.Set("Connection", "keep-alive")
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
		//resp, _ := client.Do(req) //处理请求
		defer resp.Body.Close() //关闭

		body, _ := ioutil.ReadAll(resp.Body) //读取响应
		//fmt.Println(string(body))
		//fmt.Println(string(body))
		var tem Bijgt              //用结构体
		json.Unmarshal(body, &tem) //将查到的数据放到结构体中
		//fmt.Println(string(body))    //打印body内数据
		//fmt.Println(tem)				//打印tem内数据
		bixinMysql(tem)
		time.Sleep(time.Second * 10) //设置时间
	}

	//fmt.Println(tem.Data.FansList[0].Nickname)   //打印tem中的第一个数据
	//save1(tem)
	//remaxid,_=strconv.Atoi(tem.Data.FansList[len(tem.Data.FansList)-1].Id)
}
func JM(ys int) (sign string) {
	//time1 := int(time)
	urlStr := fmt.Sprintf("pageid=%v&userid=73002494&key=kjldsnjdvndkfgsdfnsdnb", ys)
	//n2 := url.QueryEscape(urlStr)
	h := md5.New()
	h.Write([]byte(urlStr)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	n := hex.EncodeToString(cipherStr) //加密
	return n
}

func bixinMysql(bi Bijgt) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	stmt, _ := db.Prepare("insert bixin (n_id) values(?)")
	for i := 0; i < len(bi.Data); i++ {
		est, _ := stmt.Exec(bi.Data[i].Userid)
		n := est
		fmt.Println(n)
	}
	db.Close()
}
