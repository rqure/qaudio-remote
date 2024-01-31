package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	qmq "github.com/rqure/qmq/src"
)

func main() {
	app := qmq.NewQMQApplication("audio-remote")
	app.Initialize()
	defer app.Deinitialize()

	app.AddProducer("audio-remote:exchange").Initialize(10)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	log.SetFlags(log.Lmicroseconds)

	app.Producer("audio-remote:exchange").Push(&qmq.QMQAudioRequest{
		Filename: os.Getenv("AUDIO_FILE"),
	})
}
