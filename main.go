package main

import (
	"log"
	"os"
	"fmt"

	qmq "github.com/rqure/qmq/src"
)

func main() {
	app := qmq.NewQMQApplication("audio-remote")
	app.Initialize()
	defer app.Deinitialize()

	app.AddProducer("audio-player:file:exchange").Initialize(10)

	log.SetFlags(log.Lmicroseconds)

	audioFile := os.Getenv("AUDIO_FILE")
	app.Logger().Advise(fmt.Sprintf("Sending request to play: %s", audioFile))
	app.Producer("audio-player:file:exchange").Push(&qmq.QMQAudioRequest{
		Filename: audioFile,
	})
}
