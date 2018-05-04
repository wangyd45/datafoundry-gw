package apirequest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
	"time"
)

var apiHost string

var httpClientB = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
	Timeout: 0,
}

var httpClientG = &http.Client{
	Transport: httpClientB.Transport,
	Timeout:   time.Duration(10) * time.Second,
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func init() {
	apiHost = getenv("APIHOST")
	//apiHost = "new.dataos.io:8443"
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

func GenRequest(method, url, token string, body []byte) (*http.Response, error) {
	var req *http.Request
	var err error
	url = "https://" + apiHost + url

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

	return httpClientG.Do(req)

}

func WSRequest(url, token string, w http.ResponseWriter, r *http.Request) {
	var conn *websocket.Conn
	var request *http.Request
	var err error

	conn, err = wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Errorf("Failed to set websocket upgrade: %+v", err)
		return
	}

	url = "https://" + apiHost + url
	request, err = http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("request err:", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", token)

	response, _ := httpClientB.Do(request)

	defer response.Body.Close()
	defer conn.Close()
	var data = make([]byte, 0)
	var datatemp = make([]byte, 512)
	lenindex := 0
	for {
		response.Body.Read(datatemp)
		data = append(data, datatemp...)
		len := len(data)
		index := 0
		lenindex++
		for i := 512 * (lenindex - 1); i < len; i++ {
			if json.Valid(data[:i-index]) {
				conn.WriteMessage(1, data[:i-index])
				data = data[i-index:]
				index = i
				lenindex = 0
			}
		}

	}

}

func WSRequestRL(len int, url, token string, w http.ResponseWriter, r *http.Request) {
	var conn *websocket.Conn
	var request *http.Request
	var err error
	var rh http.Header = make(map[string][]string)
	rh.Set("Sec-Websocket-Protocol", r.Header.Get("Sec-Websocket-Protocol"))

	conn, err = wsupgrader.Upgrade(w, r, rh)
	if err != nil {
		fmt.Errorf("Failed to set websocket upgrade: %+v", err)
		return
	}

	url = "https://" + apiHost + url
	request, err = http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("request err:", err)
	}
	request.Header.Set("Authorization", token)

	response, _ := httpClientB.Do(request)

	defer response.Body.Close()
	defer conn.Close()
	var data = make([]byte, len)
	var jstring string
	for {
		n, _ := response.Body.Read(data)
		jstring = "{ \"message\": \"" + string(data[:n]) + "\"}"
		conn.WriteMessage(2, []byte(jstring))
	}

}
