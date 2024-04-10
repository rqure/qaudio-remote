package main

import qmq "github.com/rqure/qmq/src"

type TransformerProviderFactory struct{}

func (t *TransformerProviderFactory) Create(components qmq.EngineComponentProvider) qmq.TransformerProvider {
	transformerProvider := qmq.NewDefaultTransformerProvider()
	transformerProvider.Set("producer:audio-player:file:exchange", []qmq.Transformer{
		NewAudioRequestToAnyTransformer(components.WithLogger()),
		qmq.NewAnyToMessageTransformer(components.WithLogger()),
	})
	transformerProvider.Set("producer:audio-player:tts:exchange", []qmq.Transformer{
		NewTtsRequestToAnyTransformer(components.WithLogger()),
		qmq.NewAnyToMessageTransformer(components.WithLogger()),
	})
	return transformerProvider
}
