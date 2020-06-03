package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	//timestamp :=int32(time.Now().Unix())//获取当前时间戳
	h := md5.New()
	//n1 :=fmt.Sprintf("")

	n2 := "appType%3Dandroid%26appVersion%3D3.3.2%26appkey%3Dshow_android%26appsecret%3Db64939eddd94efa1c750f2563868c2b8%26area%3DALL%26channelType%3Da_xiaomi%26page%3D2%26pageSize%3D20%26session_id%3Dc44e6327-f806-41ec-aac0-d268373e8001%26sex%3DALL%26sortProperty%3DactiveData%26timestamp%3D1591087680%26userId%3D50166072"

	//n3 :="https://yuewowan.yangba.tv/career/queryKeDouIndexList"

	h.Write([]byte(n2)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	//fmt.Println(cipherStr)
	//fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果
	n := hex.EncodeToString(cipherStr) //加密
	fmt.Println(n)
}
