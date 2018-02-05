package storage

import "github.com/openzipkin/zipkin-go"

type Storage interface {
	SpanStore
	SpanConsumer
}

type SpanStore interface {
}

type SpanConsumer interface {
	Accept([]zipkin.Span)
}
