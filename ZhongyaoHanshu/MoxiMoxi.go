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
	"strings"
	"time"
)

type moximoxi struct {
	Data []struct {
		UID int `json:"uid"`
	} `json:"data"`
}

func main() {
	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

	timestamp := int32(time.Now().Unix()) //获取当前时间戳
	h := md5.New()
	n1 := fmt.Sprintf("myUid=285386&notUidList=&pageSize=3&pub_timestamp=%v&uid=264335&key=70d26f6a5c214d3b858f3f8daad7a161", timestamp)
	h.Write([]byte(n1)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	n := hex.EncodeToString(cipherStr) //加密
	fmt.Println(n)
	url := fmt.Sprintf("https://kelo.wduoo.com/api/moxi/follow/fans?pub_timestamp=%v&pub_sign=%v", timestamp, n)
	n2 := fmt.Sprintf("uid=264335&myUid=285386&pageSize=20&pageSize=3")
	payload := strings.NewReader(n2)
	//client := &http.Client{}       //导入请求的包
	req, _ := http.NewRequest("POST", url, payload) //开始请求
	//req.Header.Add("User-Agent","zidingyi")  //自定义表头
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("Host", "<calculated when request is sent>")
	req.Header.Set("content-length", "45")
	req.Header.Set("pub_uid", "285386")

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
	//defer resp.Body.Close()      //关闭
	body, err := ioutil.ReadAll(resp.Body) //读取响应
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	var tem moximoxi           //用结构体
	json.Unmarshal(body, &tem) //将查到的数据放到结构体中

	fmt.Println(tem)
	//moximoxiMysql(tem)

}
func moximoxiMysql(s moximoxi) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8") //链接数据库
	stmt, _ := db.Prepare("INSERT moximoxi (n_id) values (?)")                        //插入语句   字段不能填错
	for i := 0; i < len(s.Data); i++ {                                                //循环插入
		shuju, _ := stmt.Exec(s.Data[i].UID) //执行数据存储
		fmt.Println(shuju)
	}
	db.Close()
}
