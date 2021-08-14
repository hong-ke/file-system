package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"time"
)

func Logger(c *gin.Context) {
	start := time.Now().UTC()
	path := c.Request.URL.Path
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	idr := ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	request, err := readBody(idr)
	if err != nil {
		logrus.Errorf("read body bytes err:%v", err)
		return
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	ip := c.ClientIP()

	c.Next()

	end := time.Now().UTC()
	latency := end.Sub(start)
	requestID, ok := c.Get(RequestIDHeader)
	if !ok {
		logrus.Infof("%s | %s| %s |request: %s", latency, ip, path, request)
	} else {
		logrus.WithField(RequestIDInLogName, requestID).Infof("%s | %s| %s |request: %s", latency, ip, path, request)
	}

}

func readBody(reader io.Reader) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return "", errors.WithStack(err)
	}
	s := buf.String()
	return s, nil
}
