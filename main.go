package main

import (
	"fmt"
	"log"
	"os"

	qmq "github.com/rqure/qmq/src"
)

func main() {
	app := qmq.NewQMQApplication("audio-remote")
	app.Initialize()
	defer app.Deinitialize()

	app.AddProducer("audio-player:file:exchange").Initialize(10)

	log.SetFlags(log.Lmicroseconds)

	audioFile := os.Getenv("AUDIO_FILE")
	if audioFile != "" {
		app.Logger().Advise(fmt.Sprintf("Sending request to play: %s", audioFile))
		app.Producer("audio-player:file:exchange").Push(&qmq.QMQAudioRequest{
			Filename: audioFile,
		})
	}

	text := os.Getenv("TEXT_TO_SPEECH")
	if text != "" {
		app.Logger().Advise(fmt.Sprintf("Sending request to text to speech: %s", text))
		app.Producer("audio-player:tts:exchange").Push(&qmq.QMQTextToSpeechRequest{
			Text: text,
		})
	}
}
