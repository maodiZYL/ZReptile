// com.shangxin.weichat 包，该包是尚信爬虫。
//创建人：zhuyelu
//创建时间：20200608

package main

import (
	reptile_library2 "awesomeProject/reptile_library"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type SignalCommunicationID struct {
	Data []struct {
		Praises []struct {
			UserID int `json:"userId"`
		} `json:"praises"`
		UserID int `json:"userId"`
	} `json:"data"`
}

func main() {
	for i := 0; i < 90; i++ {
		shangxing(i)
	}
}

//功能方法
func shangxing(pageIndex int) {

	request := "GET"
	//qingqu:="POST"
	url := fmt.Sprintf("https://api.sx89.cn:8092/b/circle/msg/pureVideo?access_token=d96ddc73e74e424b89cfb84aa090b3f6&userId=10026521&pageSize=10&pageIndex=%v", pageIndex)
	mapNum := make(map[string]string) //用map储存键值对信息
	mapNum["Host"] = "api.sx89.cn:8092"
	body := reptile_library2.GetAgent(request, url, mapNum)
	var room SignalCommunicationID
	json.Unmarshal(body, &room)
	MysqlShangxin(room)

}

func MysqlShangxin(shang SignalCommunicationID) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	stmt, _ := db.Prepare("insert shangxin(n_id) values(?)")
	for i := 0; i < len(shang.Data); i++ {
		est, _ := stmt.Exec(shang.Data[i].UserID)
		fmt.Println(est)
	}
	db.Close()
}

//挂代理
/*func Agent(request, url string, to map[string]string) []byte {
	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

	req, _ := http.NewRequest(request, url, nil) //开始请求
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
