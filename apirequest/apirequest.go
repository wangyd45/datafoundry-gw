package apirequest

import (
	"os"
	"time"
	"net/http"
	"github.com/gorilla/websocket"
	"bytes"
	"crypto/tls"
	"fmt"
	"encoding/json"
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
	Timeout:   time.Duration(10) * time.Second,
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin:func(r *http.Request)bool{return true},
}

func init() {
	//apiHost = getenv("APIHOST")
	apiHost = "new.dataos.io:8443"
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
	url = "https://"+apiHost+url

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

func WSRequest(url, token string,w http.ResponseWriter, r *http.Request) {
	var conn *websocket.Conn
	var request *http.Request
	var err error

	conn, err = wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	url = "https://"+apiHost+url
	request,err = http.NewRequest("GET", url, nil)
	if err !=nil{
		fmt.Println("request err:",err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", token)

	response,_:=httpClientB.Do(request)


	//var data = make([]byte,10240)
	defer response.Body.Close()
	defer conn.Close()
	var data = make([]byte,0)
	var datatemp = make([]byte,512)
	for{
		response.Body.Read(datatemp)
		data = append(data,datatemp...)
		len :=len(data)
		index :=0
		//println("len ===%d",len)
		for i:=0;i<len;i++{
			if json.Valid(data[:i-index]){
				//println("-------------")
				//println(string(data[:i-index]))
				conn.WriteMessage(1,data[:i-index])
				data = data[i-index:]
				index = i
			}
		}

		/*
		n,_:=response.Body.Read(data)
		index :=0
		for i:=0;i<n;i++{
			if json.Valid(data[index:i]){
				println("index=%d",index)
				println(string(data[index:i]))
				conn.WriteMessage(1,data[index:i])
				index = i
			}

		}
		*/

	}

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

	return httpClientG.Do(req)

}


