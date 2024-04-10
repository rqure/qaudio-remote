package main

import (
	"fmt"

	qmq "github.com/rqure/qmq/src"
	"google.golang.org/protobuf/types/known/anypb"
)

type AudioRequestToAnyTransformer struct {
	logger qmq.Logger
}

func NewAudioRequestToAnyTransformer(logger qmq.Logger) qmq.Transformer {
	return &AudioRequestToAnyTransformer{
		logger: logger,
	}
}

func (t *AudioRequestToAnyTransformer) Transform(i interface{}) interface{} {
	d, ok := i.(*qmq.AudioRequest)
	if !ok {
		t.logger.Error(fmt.Sprintf("AudioRequestToAnyTransformer: invalid input %T", i))
		return nil
	}

	a, err := anypb.New(d)
	if err != nil {
		t.logger.Error(fmt.Sprintf("AudioRequestToAnyTransformer: failed to marshal AudioRequest into anypb: %v", err))
		return nil
	}

	return a
}
