package apirequest

import (
	"os"
	"time"
	"net/http"
	"github.com/gorilla/websocket"
	"bytes"
	"crypto/tls"
	"fmt"
)


var apiHost string

var httpClientB = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
	Timeout:   0,
}

var httpClientG = &http.Client{
	Transport: httpClientB.Transport,
	Timeout:   time.Duration(8) * time.Second,
}

func init() {
	//apiHost = getenv("APIHOST")
}

//获取环境变量
func getenv(env string) string {
	env_value := os.Getenv(env)
	if env_value == "" {
		fmt.Println("FATAL: NEED ENV", env)
		fmt.Println("Exit...........")
		os.Exit(2)
	}
	fmt.Println("ENV:", env, env_value)
	return env_value
}



func Request(timeout time.Duration, method, url, token string, body []byte) (*http.Response, error) {
	var req *http.Request
	var err error
	//url = apiHost + url
	url = "https://new.dataos.io:8443" + url
	if len(body) == 0 {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewReader(body))
	}

	if err != nil {
		return nil, err
	}

	if method == "PATCH" {
		req.Header.Set("Content-Type", "application/json-patch+json")
	} else {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", token)
	//req.Header.Set("Connection", "upgrade")
	//req.Header.Set("Upgrade", "websocket")
	//req.Header.Set("Sec-WebSocket-Version","13")
	//req.Header.Set("Sec-WebSocket-Key", "")


	return httpClientG.Do(req)

}

var wsupgraderT = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
}

func wshandlerT(w http.ResponseWriter, r *http.Request) {

	conn, err := wsupgraderT.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	url := "https://new.dataos.io:8443/oapi/v1/watch/projects/wutest001"
	request,err1 := http.NewRequest("GET", url, nil)
	if err1 !=nil{
		fmt.Println("request err:",err1)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer p0AdoasmlrVlrFSkAk9ywjns_YQvcig5dTr1GTVzdXU")
	//request.Header.Set("Connection", "close")

	response,_:=httpClientB.Do(request)

	//我也不知道有多大
	var data = make([]byte,1024)
	defer response.Body.Close()
	defer conn.Close()
	for{

		response.Body.Read(data)
		conn.WriteMessage(1,data)
	}

}


