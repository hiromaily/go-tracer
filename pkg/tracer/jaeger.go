package tracer

import (
	"fmt"
	"log"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/hiromaily/go-tracer/pkg/config"
)

// StartJaegerTracers is to get Jaeger tracer
func StartJaegerTracers(conf *config.TracerDetailConfig) opentracing.Tracer {
	var suffix string
	jType := jaeger.SamplerTypeConst
	probability := conf.Sampling

	if probability < 1 && probability > 0 {
		jType = jaeger.SamplerTypeProbabilistic
	}

	//NODE_HOST
	endpoint := conf.CollectorEndpoint
	if os.Getenv("NODE_HOST") != "" {
		endpoint = fmt.Sprintf("http://%s:14268/api/traces", os.Getenv("NODE_HOST"))
	}

	cfg := jaegercfg.Configuration{
		ServiceName: conf.ServiceName + suffix,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jType,
			Param: probability, // if value is 1, it always records
		},
		Reporter: &jaegercfg.ReporterConfig{
			CollectorEndpoint: endpoint,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	opts := []jaegercfg.Option{
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	}
	//if isDebug {
	//	opts = append(opts, jaegercfg.NoDebugFlagOnForcedSampling(false))
	//}

	t, _, err := cfg.NewTracer(opts...)
	if err != nil {
		log.Fatalf("Fail to start jaeger client: %v", err)
	}

	return t
}
