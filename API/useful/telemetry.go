package useful

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type LogEntry struct {
	Timestamp   string                 `json:"timestamp"`
	Level       string                 `json:"level"`
	Message     string                 `json:"message"`
	Service     string                 `json:"service"`
	Environment string                 `json:"environment"`
	TraceID     string                 `json:"trace_id,omitempty"`
	SpanID      string                 `json:"span_id,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
}

func InitTracer() func() {
	endpoint := os.Getenv("TRACE_URL")
	if endpoint == "" {
		log.Fatal("TRACE_URL environment variable not set")
	}

	log.Printf("Initializing tracer with endpoint: %s", endpoint)

	exporter, err := otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithEndpointURL(endpoint),
		otlptracehttp.WithTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatalf("Error creating exporter: %v", err)
	}

	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(os.Getenv("APPLICATION_NAME")),
			semconv.ServiceVersionKey.String("1.0.0"),
			semconv.DeploymentEnvironmentKey.String(os.Getenv("ENV")),
		),
	)
	if err != nil {
		log.Fatalf("Error creating resources: %v", err)
	}

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resources),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := provider.Shutdown(ctx); err != nil {
			log.Printf("Error stopping tracer: %v", err)
		}
	}

}

func SendLog(level string, message string, attributes map[string]any) {
	entry := LogEntry{
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
		Level:       level,
		Message:     message,
		Service:     os.Getenv("APPLICATION_NAME"),
		Environment: os.Getenv("ENV"),
		Attributes:  attributes,
	}

	if traceID, ok := attributes["trace_id"].(string); ok {
		entry.TraceID = traceID
	}
	if spanID, ok := attributes["span_id"].(string); ok {
		entry.SpanID = spanID
	}

	jsonData, err := json.Marshal(entry)
	if err != nil {
		log.Printf("Error marshaling log entry: %v", err)
		return
	}

	endpoint := os.Getenv("LOG_URL")
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating log request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending log: %v", err)
		return
	}
	defer resp.Body.Close()
}
