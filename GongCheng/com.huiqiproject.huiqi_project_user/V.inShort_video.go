// com.huiqiproject.huiqi_project_user包，该包是n号房爬虫。
//创建人：zhuyelu
//创建时间：20200607

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
)

type SignalCommunicationID struct {
	Obj struct {
		Rows []struct {
			PublisherID string `json:"publisherId"`
		} `json:"rows"`
	} `json:"obj"`
}

func main() {
	for i := 1; i < 10000; i++ {
		Climb(i)
	}
}

func Climb(pageIndex int) { //bang为结构体，接收的是结构体

	request := "GET" //请求方式
	url := fmt.Sprintf("http://47.104.97.117:9080/qx-app/userVideoBaseApi/listGoodStuffRecommendationPage?pageIndex=%v&pageSize=20", pageIndex)
	mapNum := make(map[string]string) //用map储存键值对信息
	mapNum["Host"] = "47.104.97.117:9080"
	body := Agent(request, url, mapNum) //调用Agent()
	var vin SignalCommunicationID       //用结构体
	json.Unmarshal(body, &vin)
	fmt.Println(vin)
	MysqlVin(vin)

}

func MysqlVin(s SignalCommunicationID) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8") //链接数据库
	stmt, _ := db.Prepare("INSERT vin (n_id) values (?)")                             //插入语句   字段不能填错
	for i := 0; i < len(s.Obj.Rows); i++ {                                            //循环插入
		shuju, _ := stmt.Exec(s.Obj.Rows[i].PublisherID) //执行数据存储
		fmt.Println(shuju)
	}
	db.Close()
}
func Agent(request, url string, mapNum map[string]string) []byte {

	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

	req, _ := http.NewRequest(request, url, nil) //开始请求
	for key, value := range mapNum {
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
