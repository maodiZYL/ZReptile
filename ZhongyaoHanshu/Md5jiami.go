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

	//n2 := "source=android&device_id=883dba42c33f6e900ec1f579ec6fc72f&app_style=1&version=9&ads_idf=d28c935848e22dedfbe0e345acb3ed04&channel=oppo&keywords=2"
	n2 := "latitude=28.217385&longitude=112.906904&page=1&refreshNum=1&size=10×tamp=1592394222834&key=b9b8309ca3a04a7c906c67f 6edlaibai"
	//n3 :="https://yuewowan.yangba.tv/career/queryKeDouIndexList"

	h.Write([]byte(n2)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	//fmt.Println(cipherStr)
	//fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果
	n := hex.EncodeToString(cipherStr) //加密
	fmt.Println(n)
}
