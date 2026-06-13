package custom_http

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func BuildBaseURL(c *gin.Context) string {
	scheme := "https"
	if c.Request.TLS == nil {
		scheme = "http"
	}
	return fmt.Sprintf("%s://%s", scheme, c.Request.Host)
}
