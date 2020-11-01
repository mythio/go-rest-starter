package middleware

import (
	"github.com/gin-gonic/gin"
)

// Headers adds general security headers for basic security measures
func Headers() func(*gin.Context) {
	return func(c *gin.Context) {
		// Protects from MimeType Sniffing
		c.Header("X-Content-Type-Options", "nosniff")
		// Prevents browser from prefetching DNS
		c.Header("X-DNS-Prefetch-Control", "off")
		// Denies website content to be served in an iframe
		c.Header("X-Frame-Options", "DENY")
		c.Header("Strict-Transport-Security", "max-age=5184000; includeSubDomains")
		// Prevents Internet Explorer from executing downloads in site's context
		c.Header("X-Download-Options", "noopen")
		// Minimal XSS protection
		c.Header("X-XSS-Protection", "1; mode=block")

		c.Next()
	}
}
