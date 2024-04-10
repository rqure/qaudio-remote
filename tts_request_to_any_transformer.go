package main

import (
	"fmt"

	qmq "github.com/rqure/qmq/src"
	"google.golang.org/protobuf/types/known/anypb"
)

type TtsRequestToAnyTransformer struct {
	logger qmq.Logger
}

func NewTtsRequestToAnyTransformer(logger qmq.Logger) qmq.Transformer {
	return &TtsRequestToAnyTransformer{
		logger: logger,
	}
}

func (t *TtsRequestToAnyTransformer) Transform(i interface{}) interface{} {
	d, ok := i.(*qmq.TextToSpeechRequest)
	if !ok {
		t.logger.Error(fmt.Sprintf("TtsRequestToAnyTransformer: invalid input %T", i))
		return nil
	}

	a, err := anypb.New(d)
	if err != nil {
		t.logger.Error(fmt.Sprintf("TtsRequestToAnyTransformer: failed to marshal AudioRequest into anypb: %v", err))
		return nil
	}

	return a
}
