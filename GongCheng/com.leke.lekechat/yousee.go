// com.leke.lekechat 包，该包是佑见爬虫。
//创建人：zhuyelu
//创建时间：20200608

package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
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
)

type SignalCommunication struct {
	Data struct {
		List []struct {
			FriendID int `json:"friendId"`
		} `json:"list"`
	} `json:"data"`
}

func main() {
	for i := 1; i < 10000; i++ {
		Climb_SignalommunicationID(i)
	}
}

//功能方法
func Climb_SignalommunicationID(currentPage int) {

	//request:="GET"
	request := "POST"
	url := fmt.Sprintf("http://leikebaijing.com/vv/nearby/friends")

	//body1 := CBC(currentPage)

	body1 := strings.ToUpper(CBC(currentPage)) //将字母改为大写
	//body1 := strings.ToLower(CBC(currentPage))   //将字母改为小写
	//fmt.Println(body1)
	payload := strings.NewReader(fmt.Sprintf("data=%v", body1)) //fmt.Sprintf()才能用%v  //传application/x-www-form-urlencoded   传application/json; charset=utf-8
	mapNum := make(map[string]string)                           //用map储存键值对信息
	mapNum["token"] = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJBUFAiLCJpc3MiOiJTZXJ2aWNlIiwiZXhwIjoxNjA3MTQ0OTk1LCJ1c2VySWQiOiI0OTY5IiwiaWF0IjoxNTkxNTkyOTk1fQ.6gvIBnJ63hWZ311NOQaew0Qz3vbrfIEXRdSf8wf7h_4"
	mapNum["Content-Type"] = "application/x-www-form-urlencoded"
	mapNum["Content-Length"] = "197"
	mapNum["Host"] = "leikebaijing.com"
	body := Agent(request, url, payload, mapNum)
	var yousee SignalCommunication
	json.Unmarshal(body, &yousee)
	Mysqlyousee(yousee)

}

//将数据插入表yousee
func Mysqlyousee(yousee SignalCommunication) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	stmt, _ := db.Prepare("insert yousee(n_id) values(?)")
	for i := 0; i < len(yousee.Data.List); i++ {
		est, _ := stmt.Exec(yousee.Data.List[i].FriendID)
		fmt.Println(est)
	}
	db.Close()
}

//AES CBC加密
func CBC(currentPage int) string {
	origData := []byte(fmt.Sprintf(`{"currentPage":%v,"sex":0,"latitude":28.211165,"longitude":112.900307,"limit":20}`, currentPage)) // 待加密的数据
	key := []byte("lditHOlyqbeeVe16")                                                                                                 // 加密的密钥,密码
	iv := []byte("e4l3eo89ecYsxE34")                                                                                                  //偏移量
	//log.Println("原文：", string(origData))
	//log.Println("------------------ CBC模式 --------------------")
	encrypted := AesEncryptCBC(origData, key, iv)
	//log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	payload := hex.EncodeToString(encrypted)
	return payload
}

func AesEncryptCBC(origData []byte, key []byte, iv []byte) (encrypted []byte) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                 // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)   // 补全码
	blockMode := cipher.NewCBCEncrypter(block, iv) // 加密模式
	encrypted = make([]byte, len(origData))        // 创建数组
	blockMode.CryptBlocks(encrypted, origData)     // 加密
	return encrypted
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//代理
func Agent(request, url string, payload *strings.Reader, to map[string]string) []byte {
	//创建代理
	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}
	address := fmt.Sprintf("%s:%s", "101.133.153.21", "9999")
	dialer, _ := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)

	req, _ := http.NewRequest(request, url, payload) //开始请求
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
