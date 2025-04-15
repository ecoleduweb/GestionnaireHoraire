package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
)

func Telemetry() gin.HandlerFunc {
	return func(c *gin.Context) {
		propagator := otel.GetTextMapPropagator()
		ctx := propagator.Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

		tracer := otel.GetTracerProvider().Tracer("gin-server")
		spanName := fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path)
		ctx, span := tracer.Start(ctx, spanName)
		defer span.End()

		header := make(http.Header)
		propagator.Inject(ctx, propagation.HeaderCarrier(header))
		for k, v := range header {
			c.Writer.Header()[k] = v
		}

		c.Request = c.Request.WithContext(ctx)
		c.Next()

		span.SetAttributes(
			attribute.String("http.method", c.Request.Method),
			attribute.String("http.url", c.Request.URL.String()),
			attribute.String("http.user_agent", c.Request.UserAgent()),
			attribute.String("http.flavor", c.Request.Proto),
			attribute.String("http.scheme", "http"),
			attribute.String("http.host", c.Request.Host),
		)

		span.SetAttributes(
			attribute.Int("http.status_code", c.Writer.Status()),
			attribute.Int64("response.size", int64(c.Writer.Size())),
		)
	}
}
