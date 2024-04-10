package main

import (
	"embed"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/vorbis"
)

//go:embed sounds/*/*.ogg
var sounds embed.FS

type SoundType string

const (
	SOUNDTYPE_RAIN         SoundType = "rain"
	SOUNDTYPE_RAINSTORM    SoundType = "rainstorm"
	SOUNDTYPE_THUNDERSTORM SoundType = "thunderstorm"
	SOUNDTYPE_WATER        SoundType = "water"
	SOUNDFILE_FIRST        SoundType = "001.ogg"
	SOUNDFILE_SECOND       SoundType = "002.ogg"
)

func player(t SoundType) func() {
	f1, err := sounds.Open("sounds/" + string(t) + "/" + string(SOUNDFILE_FIRST))
	if err != nil {
		log.Fatal("Failed to open file", "error", err)
	}
	streamer1, format1, err := vorbis.Decode(f1)
	if err != nil {
		log.Fatal("Failed to decode file", "error", err)
	}
	defer streamer1.Close()
	speaker.Init(format1.SampleRate, format1.SampleRate.N(time.Second/10))
	ctrl1 := &beep.Ctrl{Streamer: beep.Loop(-1, streamer1), Paused: true}
	speaker.Play(ctrl1)

	f2, err := sounds.Open("sounds/" + string(t) + "/" + string(SOUNDFILE_SECOND))
	if err != nil {
		log.Fatal("Failed to open file", "error", err)
	}
	streamer2, format2, err := vorbis.Decode(f2)
	if err != nil {
		log.Fatal("Failed to decode file", "error", err)
	}
	defer streamer2.Close()
	speaker.Init(format2.SampleRate, format2.SampleRate.N(time.Second/10))
	// this is to avoid the sound to start playing at the same time
	// and create a more natural sound
	speaker.Lock()
	err = streamer2.Seek(format2.SampleRate.N(time.Second * 3))
	if err != nil {
		log.Fatal("Failed to seek file", "error", err)
	}
	speaker.Unlock()
	ctrl2 := &beep.Ctrl{Streamer: beep.Loop(-1, streamer2), Paused: true}
	speaker.Play(ctrl2)

	toggleFunc := func() {
		speaker.Lock()
		ctrl1.Paused = !ctrl1.Paused
		ctrl2.Paused = !ctrl2.Paused
		speaker.Unlock()
	}
	return toggleFunc
}