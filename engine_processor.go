package main

import (
	"fmt"
	"os"

	qmq "github.com/rqure/qmq/src"
)

type EngineProcessor struct{}

func (ep *EngineProcessor) Process(p qmq.EngineComponentProvider) {
	audioFile := os.Getenv("AUDIO_FILE")
	if audioFile != "" {
		p.WithLogger().Advise(fmt.Sprintf("Sending request to play: %s", audioFile))
		p.WithProducer("audio-player:cmd:play-file").Push(&qmq.AudioRequest{
			Filename: audioFile,
		})
	}

	text := os.Getenv("TEXT_TO_SPEECH")
	if text != "" {
		p.WithLogger().Advise(fmt.Sprintf("Sending request to text to speech: %s", text))
		p.WithProducer("audio-player:cmd:play-tts").Push(&qmq.TextToSpeechRequest{
			Text: text,
		})
	}
}
