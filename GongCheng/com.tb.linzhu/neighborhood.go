// com.tb.linzhu 包，该包是邻住爬虫。
//创建人：zhuyelu
//创建时间：20200605
package main

import (
	"awesomeProject/reptile_library"
	"awesomeProject/write_data_to_text"
	"encoding/json"
	"fmt"
	"strconv"
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

	request := "POST"
	url := "https://app.linzhu.net/app/square/getPosts"
	payload := strings.NewReader(fmt.Sprintf("catid=%v&page=%v&flag=%v", catid, page, flag)) //fmt.Sprintf()才能用%v
	mapNum := make(map[string]string)
	mapNum["content-length"] = "22"
	//mapNum["Host"] = "<calculated when request is sent>"
	mapNum["content-type"] = "application/x-www-form-urlencoded"
	mapNum["user-agent"] = "okhttp/3.6.0"

	body := reptile_library.PostAgent(request, url, payload, mapNum)
	fmt.Println(string(body))
	var hood SignalCommunicationID
	json.Unmarshal(body, &hood)
	route := "GongCheng/com.tb.linzhu" //文件路径
	file_name := "邻住私信id"              //文件名
	for _, v := range hood.Res {
		//调用write_data_to_text包中的TraceFile()方法，将数据写入文件
		write_data_to_text.TraceFile(strconv.Itoa(v.Memberid), route, file_name)
	}
	fmt.Println(hood)

}

/*
//将数据插入数据库
func MysqlNeighborhood(hood SignalCommunicationID) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	stmt, _ := db.Prepare("insert neighborhood(n_id) values(?)")
	for i := 0; i < len(hood.Res); i++ {
		est, _ := stmt.Exec(hood.Res[i].Memberid)
		fmt.Println(est)
	}
	db.Close()
}
*/

/*func Agent(request, url string, l *strings.Reader, to map[string]string) []byte {

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

}*/
