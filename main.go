package main

import (
	qmq "github.com/rqure/qmq/src"
)

type NameProvider struct{}

func (np *NameProvider) Get() string {
	return "audio-remote"
}

func main() {
	engine := qmq.NewDefaultEngine(qmq.DefaultEngineConfig{
		NameProvider:               &NameProvider{},
		TransformerProviderFactory: &TransformerProviderFactory{},
		EngineProcessor:            &EngineProcessor{},
	})
	engine.Run()
}
