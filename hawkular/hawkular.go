package hawkular

import (
	"encoding/json"
	haw "github.com/asiainfoLDP/datafoundry-gw/apirequest"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
	"net/http"
	"os"
	"errors"
)

const (
	CPUURL     = "/hawkular/metrics/counters/data?"
	MEMORYURL  = "/hawkular/metrics/gauges/data?"
	NETWORKURL = "/hawkular/metrics/gauges/data?"
)

type Tags struct {
	Pod_namespace []string `json:"namespace"`
}

type Response struct {
	Namespace string `json:"namespace"`
	Info      []info `json:"info"`
}

type info struct {
	Start   int     `json:"start"`
	End     int     `json:"end"`
	Min     float32 `json:"min"`
	Avg     float32 `json:"avg"`
	Median  float32 `json:"median"`
	Max     float32 `json:"max"`
	Sum     float32 `json:"sum"`
	Samples int     `json:"samples"`
	Empty   bool    `json:"empty"`
}

var log lager.Logger

func init() {
	log = lager.NewLogger("hawkular")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func GainCpu(c *gin.Context) {
	//获取前端传递的Token，无需拼接"Bearer XXXXXXXXXX"
	token := pkg.GetToken(c)
	/*
		bucketDuration := c.Query("bucketDuration")
		start := c.Query("start")
		if bucketDuration == "" {
			bucketDuration = "12mn"
		}
		if start == "" {
			start = "-8h"
		}
	*/
	urlParas := pkg.SliceURL(c.Request.URL.String())
	//获取前端参数
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("Read request body error ", err)
		c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		return
	}
	var cpuTags Tags
	var cpuResList []Response
	err = json.Unmarshal(rBody, &cpuTags)
	if err != nil {
		log.Error("request body json.Unmarshal error ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	for _, v := range cpuTags.Pod_namespace {
		//URL := CPUURL + "bucketDuration=" + bucketDuration + "&start=" + start + "&tags=descriptor_name:cpu/usage,pod_namespace:" + v
		URL := CPUURL + urlParas + "&tags=descriptor_name:cpu/usage,pod_namespace:" + v
		req, err := haw.GenHawRequest("GET", URL, token, v, nil)
		defer req.Body.Close()
		if err != nil || req.StatusCode != http.StatusOK {
			log.Error("Gain cpu information fail", err)
			c.JSON(req.StatusCode, gin.H{"Namespace": v, "info": err})
			return
		}
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("read response body error ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"Namespace": v, "info": err})
			return
		}
		var rinfo []info
		err = json.Unmarshal(result, &rinfo)
		if err != nil {
			log.Error("result json.Unmarshal error ", err)
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": err})
			return
		}
		var cpuRes Response
		cpuRes.Namespace = v
		cpuRes.Info = rinfo
		cpuResList = append(cpuResList, cpuRes)
	}
	c.JSON(http.StatusOK, cpuResList)
}

func GainMemory(c *gin.Context) {
	//获取前端传递的Token，无需拼接"Bearer XXXXXXXXXX"
	token := pkg.GetToken(c)
	/*
		bucketDuration := c.Query("bucketDuration")
		start := c.Query("start")
		if bucketDuration == "" {
			bucketDuration = "12mn"
		}
		if start == "" {
			start = "-8h"
		}
	*/
	urlParas := pkg.SliceURL(c.Request.URL.String())
	//获取前端参数
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("Read request body error ", err)
		c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		return
	}
	var memoryTags Tags
	var memoryResList []Response
	err = json.Unmarshal(rBody, &memoryTags)
	if err != nil {
		log.Error("request body json.Unmarshal error ", err)
		c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": err})
		return
	}
	for _, v := range memoryTags.Pod_namespace {
		//URL := CPUURL + "bucketDuration=" + bucketDuration + "&start=" + start + "&tags=descriptor_name:memory/usage,pod_namespace:" + v
		URL := MEMORYURL + urlParas + "&tags=descriptor_name:memory/usage,pod_namespace:" + v
		req, err := haw.GenHawRequest("GET", URL, token, v, nil)
		defer req.Body.Close()
		if err != nil || req.StatusCode != http.StatusOK {
			log.Error("Gain memory information fail", err)
			c.JSON(req.StatusCode, gin.H{"Namespace": v, "info": err})
			return
		}
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("read response body error ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"Namespace": v, "info": err})
			return
		}
		var rinfo []info
		err = json.Unmarshal(result, &rinfo)
		if err != nil {
			log.Error("result json.Unmarshal error ", err)
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": err})
			return
		}

		var memoryRes Response
		memoryRes.Namespace = v
		memoryRes.Info = rinfo
		memoryResList = append(memoryResList, memoryRes)
	}
	c.JSON(http.StatusOK, memoryResList)
}

func GainNetwork(c *gin.Context) {
	//获取前端传递的Token，无需拼接"Bearer XXXXXXXXXX"
	token := pkg.GetToken(c)
	sigin := c.Param("sigin")
	var network string
	if sigin == "rx" {
		network = "rx_rate"
	} else if sigin == "tx" {
		network = "tx_rate"
	}else{
		log.Error("Read request body error ", errors.New("network param error"))
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "network param error"})
		return
	}
	/*
		bucketDuration := c.Query("bucketDuration")
		start := c.Query("start")
		if bucketDuration == "" {
			bucketDuration = "12mn"
		}
		if start == "" {
			start = "-8h"
		}
	*/
	urlParas := pkg.SliceURL(c.Request.URL.String())
	//获取前端参数
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("Read request body error ", err)
		c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		return
	}
	var networkTags Tags
	var networkResList []Response
	err = json.Unmarshal(rBody, &networkTags)
	if err != nil {
		log.Error("request body json.Unmarshal error ", err)
		c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": err})
		return
	}
	for _, v := range networkTags.Pod_namespace {
		//URL := CPUURL + "bucketDuration=" + bucketDuration + "&start=" + start + "&tags=descriptor_name:network/"+ network + ",pod_namespace:" + v
		URL := NETWORKURL + urlParas + "&tags=descriptor_name:network/" + network + ",pod_namespace:" + v
		req, err := haw.GenHawRequest("GET", URL, token, v, nil)
		defer req.Body.Close()
		if err != nil || req.StatusCode != http.StatusOK {
			log.Error("Gain network information fail", err)
			c.JSON(req.StatusCode, gin.H{"Namespace": v, "info": err})
			return
		}
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("read response body error ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"Namespace": v, "info": err})
			return
		}
		var rinfo []info
		err = json.Unmarshal(result, &rinfo)
		if err != nil {
			log.Error("result json.Unmarshal error ", err)
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": err})
			return
		}
		var networkRes Response
		networkRes.Namespace = v
		networkRes.Info = rinfo
		networkResList = append(networkResList, networkRes)
	}
	c.JSON(http.StatusOK, networkResList)
}
