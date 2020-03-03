package tracer

import (
	"os"

	"github.com/opentracing/opentracing-go"

	"github.com/hiromaily/go-tracer/pkg/config"
)

// NewTracer starts and returns a opentracing tracer
func NewTracer(conf *config.TracerConfig) opentracing.Tracer {
	switch conf.Type {
	case "jaeger":
		return StartJaegerTracers(conf.Jaeger)
	case "datadog":
		// environment variable DD_AGENT_HOST should be set to use on GCP environment
		if os.Getenv("DD_AGENT_HOST") == "" {
			return NoopTracer()
		}
		return StartDatadogTracer(conf.Datadog, os.Getenv("DD_AGENT_HOST"))
	}
	return NoopTracer()
}

// NoopTracer is no tracer
func NoopTracer() opentracing.Tracer {
	return opentracing.NoopTracer{}
}
