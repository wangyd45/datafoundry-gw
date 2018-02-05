package apirequest

import (
	"os"
	"time"
	"net/http"
	"bytes"
	"crypto/tls"
	"fmt"
	"net"
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

var httpClientD = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:        MaxIdleConns,
		MaxIdleConnsPerHost: MaxIdleConnsPerHost,
		IdleConnTimeout:	 time.Duration(IdleConnTimeout)* time.Second,
	},
	Timeout: 20 * time.Second,
}

const (
	MaxIdleConns int = 100
	MaxIdleConnsPerHost int = 100
	IdleConnTimeout int = 90
)



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

	if method == "PATCH"{
		req.Header.Set("Content-Type", "application/json-patch+json")
	}else{
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", token)

	//return httpClientD.Do(req)

	switch timeout {
	case 300:
		return httpClientD.Do(req)
	case MinRequestTimeout:
		return httpClientB.Do(req)
	default:
		return httpClientA.Do(req)
	}

}


