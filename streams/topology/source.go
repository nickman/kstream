package topology

import (
	"github.com/gmbyapa/kstream/kafka"
	"github.com/gmbyapa/kstream/streams/encoding"
)

type SourceEncoder struct {
	Key, Value encoding.Encoder
}

type Source interface {
	Node
	NodeBuilder
	Encoder() SourceEncoder
	Topic() string
	ShouldCoPartitionedWith(source Source)
	CoPartitionedWith() Source
	GetExistingParent(src Source) Source
	RePartitionedAs() Source
	AutoCreate() bool
	Internal() bool
	InitialOffset() kafka.Offset
}
