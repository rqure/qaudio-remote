package main

import qmq "github.com/rqure/qmq/src"

type TransformerProviderFactory struct{}

func (t *TransformerProviderFactory) Create(components qmq.EngineComponentProvider) qmq.TransformerProvider {
	transformerProvider := qmq.NewDefaultTransformerProvider()
	transformerProvider.Set("producer:audio-player:cmd:play-file", []qmq.Transformer{
		NewAudioRequestToAnyTransformer(components.WithLogger()),
		qmq.NewAnyToMessageTransformer(components.WithLogger(), qmq.AnyToMessageTransformerConfig{
			SourceProvider: components.WithNameProvider(),
		}),
	})
	transformerProvider.Set("producer:audio-player:cmd:play-tts", []qmq.Transformer{
		NewTtsRequestToAnyTransformer(components.WithLogger()),
		qmq.NewAnyToMessageTransformer(components.WithLogger(), qmq.AnyToMessageTransformerConfig{
			SourceProvider: components.WithNameProvider(),
		}),
	})
	return transformerProvider
}
