package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	GinContextErrorKey = "Error"
	RequestIDInLogName = "request_id"
)

func Error(c *gin.Context) {
	c.Next()
	message, ok := c.Get(GinContextErrorKey)
	if ok {
		b, _ := json.Marshal(message)
		requestID, ok := c.Get(RequestIDHeader)
		if !ok {
			logrus.Errorln(string(b))
			return
		}
		logrus.WithField(RequestIDInLogName, requestID).Errorln(string(b))

	}
}

//当层产生的错误
func SetError(c *gin.Context, err error) error {
	c.Set(GinContextErrorKey, err)
	return errors.WithStack(err)
}
