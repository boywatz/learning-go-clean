package middlewares

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (rw responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		brw := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}

		defer func() {
			logrus.Infof("url: %s duration: %v sec", c.Request.URL.Path, time.Since(start).Seconds())
			logrus.Infof("Request Body: %v", c.Request.Body)
			logrus.Infof("Response Body: %v", brw.body.String())
		}()

		c.Writer = brw
		c.Next()
	}
}

func GreetingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Infof("Greeting Middleware Start.")
		c.Next()
		logrus.Infof("Greeting Middleware End.")
	}
}

func CommonMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Infof("Common Middleware Start.")
		c.Next()
		logrus.Infof("Common Middleware End.")
	}
}
