package app

import (
	"synth/app/output"
)

const (
	SampleRate float64 = 44100.0
	FinalVolume float64 = 0.5 // 0.5 is the max
	OutputFile string = "/Users/shaunbillows/coding/projects/synth/app/sounds.wav"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	output := output.NewOtoPlayer(SampleRate) // Write to player
	// output := output.NewWavWriter(OutputFile, SampleRate, true) // Write to WAV file
	defer output.Close()

	BbMajor7thChord(1.5, 1).Write(output, FinalVolume)
    A7Chord(1.5, 1).Write(output, FinalVolume)
    DMinor7Chord(3, 1).Write(output, FinalVolume)
	BbMajor7thChord(1.5, 1).Write(output, FinalVolume)
    A7Chord(1.5, 1).Write(output, FinalVolume)
    DMinor7Chord(3, 1).Write(output, FinalVolume)
}
