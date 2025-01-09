package middlewares

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func LatencyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		// latency = now - start
		latency := time.Since(start)

		log.Info().
			Str("Latency", latency.String()).Msg("Calculated Latency")
	}
}
