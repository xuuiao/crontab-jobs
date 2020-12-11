package cronjob

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"net/http"
	"time"
)

func NewJob() *jobServer {
	return &jobServer{}
}

type jobServer struct {
	cronTab *cron.Cron
}

func (j *jobServer) Start() {
	j.cronTab = cron.New()
	_, _ = j.cronTab.AddFunc("*/1 * * * *", func() {
		fmt.Printf("%s 每1分钟执行:调取数据迁移", time.Now().Format("15:04:05"))
		url := "https://crs-api.vchangyi.com/micro-scrm/v1/members/migrate-process"
		body := struct {
			Size int `json:"size"`
		}{
			Size: 500,
		}
		bodyStr, _ := json.Marshal(body)
		request, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyStr))
		if err != nil {
			fmt.Printf("生成request失败, err:%s \n", err.Error())
			return
		}
		defer request.Body.Close()
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Printf("请求失败, err:%s \n", err.Error())
			return
		}
		defer response.Body.Close()
		result, _ := ioutil.ReadAll(response.Body)
		fmt.Printf("response: %s \n", string(result))
	})

	j.cronTab.Start()
}

func (j *jobServer) ShutDown() error {
	stop := j.cronTab.Stop()

	select {
	case <-stop.Done():
		return nil
	}
}
