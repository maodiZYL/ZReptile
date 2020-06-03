package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	//_ "github.com/go-sql-driver/mysql" //导数据库
	"golang.org/x/net/proxy" //导代理
	"io/ioutil"
	"net/http"
	"strings"
)

type Auto struct {
	Data struct {
		NewFans []struct {
			BefollowID int `json:"befollow_id"`
		} `json:"newFans"`
	} `json:"data"`
}

func main() {
	for i := 1; i <= 784; i++ {
		url := "https://app.jinglantech.tech/message/newfans"
		//创建代理
		auth := proxy.Auth{
			User:     "itemb123",
			Password: "kIl8Jl3aKej",
		}
		address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
		dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

		payload := strings.NewReader(fmt.Sprintf("reqUserId=68347218&pageNo=%v&pageSize=60&other_user_id=100004", i)) //fmt.Sprintf()才能用%v
		req, _ := http.NewRequest("POST", url, payload)
		req.Header.Add("Content-Length", "<calculated when request is sent>")
		req.Header.Add("Host", "<calculated when request is sent>")
		req.Header.Add("user_id", "68347218")
		req.Header.Add("content-type", "application/x-www-form-urlencoded")
		req.Header.Add("token", "QFUMhufACrpSRP9JbWZhsQ==")

		//使用代理
		var resp *http.Response
		httpTransport := &http.Transport{
			//跳过证书验证
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient := &http.Client{Transport: httpTransport}
		if dialer != nil {
			httpTransport.Dial = dialer.Dial
		}
		resp, _ = httpClient.Do(req)
		//resp,_ :=  http.DefaultClient.Do(req)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		var yao Auto
		json.Unmarshal(body, &yao)
		fmt.Println(yao)
	}
}

/*func yaowang(userlist Auto)  {
	db, _:= sql.Open("mysql","root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	stmt,_ := db.Prepare("insert yaowang (n_id ,Name ,sex) values(?,?,?)")
	for i:=0;i<len(userlist.Data.NewFans);i++ {
		est,_ := stmt.Exec(userlist.Data.NewFans[i].ID,userlist.Data.NewFans[i].Name,userlist.Data.NewFans[i].Sex)
		fmt.Println(est)
	}
	db.Close()
}*/
