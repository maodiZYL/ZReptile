package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"strings"
)

type GuaiZhu struct {
	Data []struct {
		UID      int    `json:"uid"`
		Gender   int    `json:"gender"`
		Username string `json:"username"`
	} `json:"data"`
}

func main() {
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 4; j++ {

			url := "http://sayu-api.chongqiwawa2.com/action/sayu/recommendhourlist"
			payload := strings.NewReader(fmt.Sprintf("cycle=%v&type=%v", j, i)) //fmt.Sprintf()才能用%v
			req, _ := http.NewRequest("POST", url, payload)
			req.Header.Add("Content-Length", "28")
			req.Header.Add("Host", "sayu-api.chongqiwawa2.com")
			req.Header.Add("content-type", "application/x-www-form-urlencoded")
			req.Header.Add("sign", "tgspOgaG3er4rQzN0iD3im4dwCDx1xLOs1I1IlNm2g5iHfR2FfjwTacLUSEc20DiL2Oi5jRjo%2F6TGhvbI5GIE9bqVvaucEen1UtUsvLDcrZgJJPxfCUoteIUYWvFwm8C5s6WqKJ2CQIQkqxTcsrwNmKOz%2BqBZbKjDJS4ANeDfGb6nBV159X8W8yWpg82Tn07ECV4sN6cKXVmshzs2%2BxKifrdXf2zAxOw70oJUzt9JTH4D4250e5VM7bue3VHd80U6DcG27bYyP1eEaA%3D")
			resp, _ := http.DefaultClient.Do(req)
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			var zhu GuaiZhu
			json.Unmarshal(body, &zhu)
			//fmt.Println(body)
			//fmt.Println(zhu)
			GuaiZhuMsql(zhu)

		}
	}

}

func GuaiZhuMsql(guai GuaiZhu) {
	db, _ := sql.Open("mysql", "root:haosql@tcp(127.0.0.1:3306)/blogdb?charset=utf8")
	stmt, _ := db.Prepare("insert GuaiZhu (n_id ,Name ,sex) values(?,?,?)")
	for i := 0; i < len(guai.Data); i++ {
		est, _ := stmt.Exec(guai.Data[i].UID, guai.Data[i].Username, guai.Data[i].Gender)
		fmt.Println(est)
	}
	db.Close()
}
