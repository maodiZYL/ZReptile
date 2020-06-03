package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	//"time"
)

func main() {
	//timestamp :=int32(time.Now().Unix())

	Md5JM(2, 1591087680)
	//fmt.Println(yeshu,t1)
}
func Md5JM(ys int, time int) {
	//ysh:=string(ys)
	//time1:=string(time)
	urlStr := fmt.Sprintf("appType=android&appVersion=3.3.2&appkey=show_android&appsecret=b64939eddd94efa1c750f2563868c2b8&area=ALL&channelType=a_xiaomi&page=%v&pageSize=20&session_id=c44e6327-f806-41ec-aac0-d268373e8001&sex=ALL&sortProperty=activeData&timestamp=%v&userId=50166072", ys, time)

	n2 := url.QueryEscape(urlStr)
	fmt.Println(n2)

	h := md5.New()

	h.Write([]byte(n2)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	//fmt.Println(cipherStr)
	//fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果
	n := hex.EncodeToString(cipherStr) //加密
	fmt.Println(n)
}
