package handler

import (
	"github.com/pivotal-golang/lager"
	"os"
	"sync/atomic"
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	"bytes"
	"crypto/tls"
	"fmt"
	"errors"
	"github.com/openshift/origin/pkg/cmd/util/tokencmd"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
)

//==============================================================
//
//==============================================================

type OpenshiftClient struct {
	host string
	oapiUrl string
	kapiUrl string
	namespace string
	username  string
	password  string
	bearerToken atomic.Value
}

var theOC *OpenshiftClient
var logger lager.Logger

type E string
func (e E) Error() string {
	return string(e)
}
const (
	NotFound = E("not found")
)

func newOpenshiftClient(host, username, password, defaultNamespace string) *OpenshiftClient {
	host = "https://" + host
	oc := &OpenshiftClient{
		host: host,
		oapiUrl: host + "/oapi/v1",
		kapiUrl: host + "/api/v1",

		namespace: defaultNamespace,
		username:  username,
		password:  password,
	}
	oc.bearerToken.Store("")

	go updateBearerToken()

	return oc
}

func updateBearerToken() {
	for {
		clientConfig := &kclient.Config{}
		clientConfig.Host = theOC.host
		clientConfig.Insecure = true

		token, err := tokencmd.RequestToken(clientConfig, nil, theOC.username, theOC.password)
		if err != nil {
			println("RequestToken error: ", err.Error())

			time.Sleep(15 * time.Second)
		} else {

			theOC.bearerToken.Store(token)

			println("RequestToken token: ", token)

			time.Sleep(3 * time.Hour)
		}
	}
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

func init() {
	logger = lager.NewLogger("OcApi")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))

	theOC = newOpenshiftClient(
		getenv("OCAPIHOST"),
		getenv("OCAPIUSER"),
		getenv("OCAPIPW"),
		getenv("OCNAMESPACE"),
	)
}

//==============================================================


func GetProjects(c *gin.Context){

	token := theOC.bearerToken.Load().(string)
	if token == "" {
		errors.New("token is blank")
		return
	}
	request(GeneralRequestTimeout, "GET", "/oapi/v1/projects", token,nil)
	return
}

const (
	GeneralRequestTimeout = time.Duration(30) * time.Second
	DfRequestTimeout = time.Duration(8) * time.Second

)
var httpClientA = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
	Timeout:   0,
}
var httpClientB = &http.Client{
	Transport: httpClientA.Transport,
	Timeout:   DfRequestTimeout,
}
var httpClientC = &http.Client{
	Transport: httpClientA.Transport,
	Timeout:   GeneralRequestTimeout,
}

func request(timeout time.Duration, method, url, bearerToken string, body []byte) (*http.Response, error) {
	var req *http.Request
	var err error
	if len(body) == 0 {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewReader(body))
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+bearerToken)


	switch timeout {
	case GeneralRequestTimeout:
		return httpClientC.Do(req)
	case DfRequestTimeout:
		return httpClientB.Do(req)
	default:
		return httpClientA.Do(req)
	}
}


