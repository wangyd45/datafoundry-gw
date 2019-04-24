package others

import (
	oapi "github.com/asiainfoldp/datafoundry-gw/apirequest"
	"github.com/asiainfoldp/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
	"net/http"
	"os"
)

var logger lager.Logger

func init() {
	logger = lager.NewLogger("other_api")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func AnyRequest(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	url := c.Request.URL.String()
	method := c.Request.Method
	body, err := ioutil.ReadAll(c.Request.Body)
	logger.Info("AnyRequest host is  " + host)
	if err != nil {
		logger.Error("AnyRequest Read Request.Body error", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest", "metrics": err})
		return
	}
	code, result, err := Any(method, host + url, token, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError", "metrics": err})
		return
	}
	c.Data(code, "application/json", result)
	return
}

func Any(method, url, token string, body []byte) (code int, result []byte, err error) {
	logger.Debug("any body is " + string(body))
	req, err := oapi.GenRequest(method, url, token, body)
	if err != nil {
		logger.Error("AnyRequest error ", err)
		return req.StatusCode, nil, err
	}
	defer req.Body.Close()
	result, err = ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error("AnyRequest Read req.Body error ", err)
		return http.StatusInternalServerError, result, err
	}
	return req.StatusCode, result, nil
}
