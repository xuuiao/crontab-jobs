package jobs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func InitEnterprise() {
	jsonStr := `[
 ]`

	type enterprise struct {
		EnterpriseID int64  `json:"enterprise_id"`
		CorpID       string `json:"corp_id"`
		Domain       string `json:"domain"`
		Hash         string `json:"hash"`
	}
	var requestList []*enterprise
	err := json.Unmarshal([]byte(jsonStr), &requestList)
	if err != nil {
		fmt.Printf("json字符串错误,err:%s \n", err.Error())
	}
	for _, v := range requestList {
		fmt.Printf("request: %v \n", v)
		url := "https://crs-api.vchangyi.com/micro-scrm/v1/callback/enterprise-add"
		bodyStr, _ := json.Marshal(v)
		request, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyStr))
		if err != nil {
			fmt.Printf("生成request失败, err:%s \n", err.Error())
			return
		}
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Printf("请求失败, err:%s \n", err.Error())
			return
		}
		result, _ := ioutil.ReadAll(response.Body)
		fmt.Printf("response: %s \n", string(result))
		response.Body.Close()
		request.Body.Close()
		time.Sleep(5 * time.Second)
	}
}

func InitTag() {
	var idList = []int64{
		1,
	}

	for _, v := range idList {
		fmt.Printf("request: %v \n", v)
		url := "https://crs-api.vchangyi.com/micro-scrm/v1/tag/sync"
		bodyStr, _ := json.Marshal(v)
		request, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyStr))
		if err != nil {
			fmt.Printf("生成request失败, err:%s \n", err.Error())
			return
		}
		request.Header.Add("enterprise-id", strconv.FormatInt(v, 10))
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Printf("请求失败, err:%s \n", err.Error())
			return
		}
		result, _ := ioutil.ReadAll(response.Body)
		fmt.Printf("response: %s \n", string(result))
		response.Body.Close()
		request.Body.Close()
		time.Sleep(2 * time.Second)
	}
}

type TestDefault struct {
}

func (t TestDefault) Show(a string) error {
	return nil
}
