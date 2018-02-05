package elasticsearch

import (
	"log"

	"github.com/epy0n0ff/zipkin-go/model"
	"github.com/openzipkin/zipkin-go"
)

type ElasticsearchSpanConsumer struct {
	logger *log.Logger
	spanC  chan *model.SpanModel
}

func (c *ElasticsearchSpanConsumer) Accept(spans []zipkin.Span) {
	if len(spans) == 0 {
		// empty span
	}
	// convert a span to json
	// do posting _bulk endpoint
}
