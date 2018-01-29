package apirequest

import (
	"os"
	"time"
	"net/http"
	"bytes"
	"crypto/tls"
	"fmt"
)


const (
	MaxRequestTimeout = time.Duration(30) * time.Second
	MinRequestTimeout = time.Duration(10) * time.Second

)

var apiHost string

var httpClientA = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
	Timeout:   0,
}
var httpClientB = &http.Client{
	Transport: httpClientA.Transport,
	Timeout:   MinRequestTimeout,
}
var httpClientC = &http.Client{
	Transport: httpClientA.Transport,
	Timeout:   MaxRequestTimeout,
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

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)


	switch timeout {
	case MaxRequestTimeout:
		return httpClientC.Do(req)
	case MinRequestTimeout:
		return httpClientB.Do(req)
	default:
		return httpClientA.Do(req)
	}
}


