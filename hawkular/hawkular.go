package hawkular

import (
	haw "github.com/asiainfoLDP/datafoundry-gw/apirequest"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
	"os"
	"encoding/json"
	"net/http"
	//"fmt"
	"fmt"
)

const (
	CPUURL = "/hawkular/metrics/counters/data?"
	MEMORYURL = "/hawkular/metrics/gauges/data?"
	NETWORKURL = "/hawkular/metrics/gauges/data?"
)

type Tags struct{
	Pod_namespace []string `json:"namespace"`
}

type Response struct{
	Namespace string `json:"namespace"`
	Info []info `json:"info"`
}

type info struct {
	Start int `json:"start"`
	End int `json:"end"`
	Min float32 `json:"min"`
	Avg float32 `json:"avg"`
	Median float32 `json:"median"`
	Max float32 `json:"max"`
	Sum float32 `json:"sum"`
	Samples int `json:"samples"`
	Empty bool `json:"empty"`
}

var log lager.Logger

func init() {
	log = lager.NewLogger("hawkular")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func GainCpu(c *gin.Context) {
	//获取前端传递的Token，无需拼接"Bearer XXXXXXXXXX"
	token := pkg.GetToken(c)
	bucketDuration := c.Query("bucketDuration")
	start := c.Query("start")
	//获取前端参数
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		log.Error("Read request body error ",err)
		c.JSON(http.StatusExpectationFailed,gin.H{"error":err})
		return
	}
	var cpuTags Tags
	var cpuResList []Response
	err = json.Unmarshal(rBody,&cpuTags)
	if err != nil{
		log.Error("request body json.Unmarshal error ",err)
		c.JSON(http.StatusUnsupportedMediaType,gin.H{"error":err})
		return
	}
	for _,v := range cpuTags.Pod_namespace{
		URL := CPUURL + "bucketDuration=" + bucketDuration + "&start=" + start + "&tags=descriptor_name:cpu/usage,pod_namespace:" + v
		req, err := haw.GenHawRequest("GET", URL, token,v, rBody)
		if err != nil  || req.StatusCode != http.StatusOK{
			log.Error("Gain cpu information fail", err)
			c.JSON(req.StatusCode,gin.H{"Namespace":v,"info":err})
			return
		}
		result, err := ioutil.ReadAll(req.Body)
		if err != nil{
			log.Error("read response body error ",err)
			c.JSON(http.StatusInternalServerError,gin.H{"Namespace":v,"info":err})
			return
		}
		var rinfo []info
		err = json.Unmarshal(result,&rinfo)
		if err != nil{
			log.Error("result json.Unmarshal error ",err)
			c.JSON(http.StatusUnsupportedMediaType,gin.H{"error":err})
			return
		}
		defer req.Body.Close()
		var cpuRes Response
		cpuRes.Namespace = v
		cpuRes.Info = rinfo
		cpuResList = append(cpuResList,cpuRes)
	}
	c.JSON(http.StatusOK,cpuResList)
}

/*
func GainMemory(c *gin.Context) {
	//获取前端传递的Token，无需拼接"Bearer XXXXXXXXXX"
	token := pkg.GetToken(c)
	//获取前端参数
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req, err := oapi.GenRequest("POST", "/oapi/v1/projectrequests", token, rBody)
	if err != nil {
		logger.Error("Create A Project Fail", err)
	}
	logger.Info("Create project", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	//返回结果JSON格式
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GainNetwork(c *gin.Context) {
	//获取前端传递的Token，无需拼接"Bearer XXXXXXXXXX"
	token := pkg.GetToken(c)
	//获取前端参数
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req, err := oapi.GenRequest("POST", "/oapi/v1/projectrequests", token, rBody)
	if err != nil {
		logger.Error("Create A Project Fail", err)
	}
	logger.Info("Create project", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	//返回结果JSON格式
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}
*/