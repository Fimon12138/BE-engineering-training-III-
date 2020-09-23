package midderware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tickethub_service/util/log"
	"time"
)

func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		param := gin.LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}

		// Stop timer
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path

		log.Infof(DefaultLogFormatter(param))
	}
}
func DefaultLogFormatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf(" %s %s %d %s %3d %s\n",
		param.Method,
		param.Path,
		param.StatusCode,
		param.Request.Proto,
		param.Latency,
		param.ErrorMessage,
	)
}
