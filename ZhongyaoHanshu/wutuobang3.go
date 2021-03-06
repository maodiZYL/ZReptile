package main

import (
	"bytes"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
)

type Wutuobang1 struct {
	Spaces []struct {
		Identifier string `json:"identifier"`
	} `json:"spaces"`
}
type Wutuobang2 struct {
	Members []struct {
		ID string `json:"id"`
	} `json:"members"`
}

type tongxing2 struct {
	ImC2CGroupID string `json:"im_c2c_group_id"`
}

func main() {
}

/*func main() {

	var n string
	for i := 1; i < 20; i++ {
		body := bangName()
		var tem Wutuobang1         //用结构体
		json.Unmarshal(body, &tem) //将查到的数据放到结构体中
		//tem.Spaces[0].Identifier
		for i, _ := range tem.Spaces {
			//fmt.Printf(tem.Spaces[i].Identifier)
			n = tem.Spaces[i].Identifier
			yonghu(n)
		}
	}
}
*/
func bangName() []byte {
	url := fmt.Sprintf("https://app.quanziapp.com/api/v1/recommended_spaces")
	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

	req, _ := http.NewRequest("GET", url, nil) //开始请求
	//req.Header.Set("Host", "<calculated when request is sent>")
	req.Header.Set("x-app-version", "Android Circles 3.5.3")
	req.Header.Set("authorization", "token Vnn9DLoPTnscLgoMRcG9eNXT1590739356.5220678")

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
	body, _ := ioutil.ReadAll(resp.Body) //读取响应
	return body
}

func yonghu(bangming string) {
	for i := 1; i < 284; i++ {
		ne := bangming
		url := fmt.Sprintf("https://app.quanziapp.com/api/v2/%v/members?&page=1&per_page=100", ne)
		//创建代理
		auth := proxy.Auth{
			User:     "itemb123",
			Password: "kIl8Jl3aKej",
		}
		address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
		dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

		req, _ := http.NewRequest("GET", url, nil) //开始请求
		//req.Header.Set("Host", "<calculated when request is sent>")
		req.Header.Set("x-app-version", "Android Circles 3.5.3")
		req.Header.Set("authorization", "token Vnn9DLoPTnscLgoMRcG9eNXT1590739356.5220678")

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
		//resp, _ := client.Do(req)   //处理请求
		body, _ := ioutil.ReadAll(resp.Body) //读取响应
		var tem2 Wutuobang2                  //用结构体
		json.Unmarshal(body, &tem2)
		//将查到的数据放到结构体中
		/*for i := 0; i < len(tem2.Members); i++ {
			n := tem2.Members[i].ID
			Tongxing2(ne, n)
		}*/

		Mysql5(tem2)

	}
}

func Tongxing2(bang string, bang2 string) { //bang为结构体，接收的是结构体
	//循环遍历结构体
	url := fmt.Sprintf("https://app.quanziapp.com/api/v2/im/init_c2c_group_contact")

	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

	//sBAKAgj
	//jRMr01
	requestBody := fmt.Sprintf(`{"identifier":"%v","member_id":"%v"}`, bang, bang2) //传application/json; charset=utf-8
	var jsonStr = []byte(requestBody)

	//payload := strings.NewReader(fmt.Sprintf("befollow_id=%v&user_id=98236258&is_login=0",bang.Members[i].ID)) //fmt.Sprintf()才能用%v  //传application/x-www-form-urlencoded

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr)) //开始请求
	//req.Header.Set("Host", "<calculated when request is sent>")
	req.Header.Set("authorization", "token Vnn9DLoPTnscLgoMRcG9eNXT1590739356.5220678")
	//req.Header.Set("Content-Length", "<calculated when request is sent>")
	req.Header.Set("x-app-version", "Android Circles 3.5.3")
	req.Header.Set("content-type", "application/json; charset=utf-8")
	req.Header.Set("accept-language", "zh")
	req.Header.Set("user-agent", "okhttp/3.14.4")

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
	body, _ := ioutil.ReadAll(resp.Body) //读取响应
	var tem tongxing2                    //用结构体
	json.Unmarshal(body, &tem)           //将查到的数据放到结构体中
	//fmt.Println(string(body))
	fmt.Println(string(body))
	//
	//    Mysql5(tem)
}
func Mysql5(wubang Wutuobang2) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8") //链接数据库
	stmt, _ := db.Prepare("INSERT wutuo1 (n_id) values (?)")                          //插入语句   字段不能填错
	for i := 0; i < len(wubang.Members); i++ {                                        //循环插入
		shuju, _ := stmt.Exec((wubang.Members[i].ID)) //执行数据存储
		fmt.Println(shuju)
	}
	db.Close()
}
