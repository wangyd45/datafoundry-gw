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
)

const (
	CPUURL = "/hawkular/metrics/counters/data?"
	MEMORYURL = "/hawkular/metrics/gauges/data?"
	NETWORKURL = "/hawkular/metrics/gauges/data?"
)

type Tags struct{
	Descriptor_name string `json:"descriptor_name"`
	Pod_namespace string `json:"pod_namespace"`
}
/*
type Response struct {
	Start int `json:"start"`
	End int `json:"end"`
	Min float32 `json:"min"`
	Avg float32 `json:"avg"`
	Median float32 `json:"median"`
	Max float32 `json:"max"`
	Sum float32 `json:"sum"`
	Samples int `json:"samples"`
	Empty bool `json:"empty"`
}*/

type Response struct{
	Namespace string `json:"namespace"`
	result string `json:"result"`
}

var logger lager.Logger

func init() {
	logger = lager.NewLogger("hawkular")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func GainCpu(c *gin.Context) {
	//获取前端传递的Token，无需拼接"Bearer XXXXXXXXXX"
	token := pkg.GetToken(c)
	bucketDuration := c.Param("bucketDuration")
	start := c.Param("start")
	//获取前端参数
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	var cpuTags []Tags
	var cpuResList []Response
	json.Unmarshal(rBody,&cpuTags)
	for _,v := range cpuTags{
		URL := CPUURL + "bucketDuration=" + bucketDuration + "&start=" + start + "&tags=descriptor_name:"+ v.Descriptor_name + ",pod_namespace:" + v.Pod_namespace
		req, err := haw.GenRequest("POST", URL, token, rBody)
		if err != nil {
			logger.Error("Gain cpu information fail", err)
			return
		}
		result, err := ioutil.ReadAll(req.Body)
		if err != nil{
			logger.Error("read response body error ",err)
			return
		}
		defer req.Body.Close()
		var cpuRes Response
		cpuRes.Namespace = v.Pod_namespace
		cpuRes.result = string(result)
		cpuResList = append(cpuResList,cpuRes)
	}
	c.JSON(http.StatusAccepted,gin.H{"cpuinfo":cpuResList})
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