package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type LiKa struct {
	Data []struct {
		UID      int    `json:"uid"`
		Gender   int    `json:"gender"`
		Username string `json:"username"`
	} `json:"data"`
}

func main() {

	timestamp := int32(time.Now().Unix()) //获取当前时间戳
	fmt.Println(timestamp)
	url := fmt.Sprintf("https://api.new.server.ourydc.cn/api/userExtend/friendListByUserEx?t=%v&userId=0002020052816054224000", timestamp)
	//url :="http://sayu-api.chongqiwawa2.com/action/sayu/recommendhourlist"
	payload := strings.NewReader(fmt.Sprintf("n=18E08084734B2E70A42EF2F845FB215175E7FD5C08F4B228D7492868E2A18E3C90CF31040FFE67D5D2BE40DECD3BF74A8F777605701C77B94487C7A2BE3AB94DABDDCDD80BF2089AB810A7F77B9D8720A33E65BA8C0B8EBDF7F7A6FC08072D74AB67F3E710C4C60968D663944E4859B0FA8E422BBEE02882578DEAA3BBE896C8F7DF9E04BB8782C363B7C0897A240172A08A0BEEB22C402677E97E2FD3CEFF6E8303FF205C2D9227725CA78D2C17D893234A1646FBEF5141C7A46DB8125E7EACBE9153F45EC179E8BAA6E2BBF610FE853F3D5FDDEB23B02AA42242BC0216EC09F15EE241BFEDFCDE5E2CDBD00E16C08A1FB171F77E90FD5B05DB5E891D4C235DAF34B6DF1F2DC84D5AB2C924AA06FE65C0AACD471FC284E9E620D5E17EE83E38976B7214ECE11B0C566EE85911396690FD15CC47C7F5442CF393018E11CBCDB3208CB5074D31670C37705A452AD7399B8204F7DF4BC0002B0B1A7C2F3CAFC8C83E48C24DF76A46F602C1E45F54B2378B2AEDA4E0C1B09BC023158AA72F55955A6828B34AA63A337AEA63A4DB18B0B3077DBED0F7E95F5E3F326F524FBFE4ADC525A6FB71738E50E59F47B8EE47E08E7BDF00D3B9B16B9750F59B43B3397D2B65BE8988B64BE4D99A6898D23FEF5FBEB4F284B6E8B024881976AA3561909CB6F2768406B26885D5A88167D739FC5766B2C9CBCEB27FDA1137A7E2F5F15FA4C7EB881C7B99F5D74FD3AA6784B38A23F619D3D441CAB120F6FDDD624795BF26AC64C27F60E1C9E11A54")) //fmt.Sprintf()才能用%v
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Length", "28")
	req.Header.Add("Host", "sayu-api.chongqiwawa2.com")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("sign", "tgspOgaG3er4rQzN0iD3im4dwCDx1xLOs1I1IlNm2g5iHfR2FfjwTacLUSEc20DiL2Oi5jRjo%2F6TGhvbI5GIE9bqVvaucEen1UtUsvLDcrZgJJPxfCUoteIUYWvFwm8C5s6WqKJ2CQIQkqxTcsrwNmKOz%2BqBZbKjDJS4ANeDfGb6nBV159X8W8yWpg82Tn07ECV4sN6cKXVmshzs2%2BxKifrdXf2zAxOw70oJUzt9JTH4D4250e5VM7bue3VHd80U6DcG27bYyP1eEaA%3D")
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var li LiKa
	json.Unmarshal(body, &li)
	//fmt.Println(body)
	//fmt.Println(zhu)
	LikaMsql(li)
}

func LikaMsql(li LiKa) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	stmt, _ := db.Prepare("insert GuaiZhu (n_id ,Name ,sex) values(?,?,?)")
	for i := 0; i < len(li.Data); i++ {
		est, _ := stmt.Exec(li.Data[i].UID, li.Data[i].Username, li.Data[i].Gender)
		fmt.Println(est)
	}
	db.Close()
}
