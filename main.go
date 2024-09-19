package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func loadSound(filename string) (beep.StreamSeekCloser, beep.Format) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return streamer, format
}

func main() {
	downSound, format := loadSound("./audio/down.mp3")
	// fullSound, _ := loadSound("full.mp3")
	upSound, _ := loadSound("./audio/up.mp3")
	// soundSound, _ := loadSound("sound.mp3")

	err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Start typing to hear keyboard sounds... Press ESC to quit")

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			log.Println(err)
			continue
		}

		var soundToPlay beep.StreamSeekCloser

		switch key {
		case keyboard.KeyEsc:
			fmt.Println("Exiting...")
			return
		default:
			soundToPlay = downSound
		}

		soundToPlay.Seek(0)
		speaker.Play(soundToPlay)

		// Play the up sound after a short delay
		time.Sleep(10 * time.Millisecond)
		upSound.Seek(0)
		speaker.Play(upSound)
	}
}