package main

import (
	"fmt"
	"time"
)

func main() {
	timestamp := int32(time.Now().Unix()) //获取当前时间戳
	fmt.Println(timestamp)
}
