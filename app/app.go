package app

import (
	"synth/app/output"
)

const (
	SampleRate float64 = 44100.0
	Volume float64 = 0.5 // 0.5 is the max
	OutputFile string = "/Users/shaunbillows/coding/projects/synth/app/track.wav"
	Bpm = 120 // beats per minute
	Bar = 60.0 / Bpm * 4 // seconds
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	output := output.NewOtoPlayer(SampleRate) // Write to player
	// output := output.NewWavWriter(OutputFile, SampleRate, true) // Write to WAV file
	defer output.Close()

	highHat := NewHighHat(1)
	kickDrum := NewKickDrum(2)
	snare := NewSnare(1)

	kickDrumLoop := Sequencer(kickDrum, 4 * Bar, Bpm, false)
	snareLoop := Sequencer(snare, 4 * Bar, Bpm/2, true)

	highHatLoopStraight := Sequencer(highHat, 4 * Bar, Bpm, true)
	HighHatLoopPolyrhythm := Sequencer(highHat, 4 * Bar, Bpm/3, true)
	highHatLoop := highHatLoopStraight.Superpose(HighHatLoopPolyrhythm)

	chord1 := NewAMinorChord(Bar, 1)
	chord2 := NewEMinorChord(Bar, 1)
	chord3 := NewDMinorChord(Bar, 1)
	chord4 := NewGMajorChord(Bar, 1)

	chordProgression := chord1.Add(chord2).Add(chord3).Add(chord4)

	verse := chordProgression.Superpose(kickDrumLoop, snareLoop)
	chorus := verse.Superpose(highHatLoop)
	verse2 := verse.Superpose(snare, highHat, kickDrum)
	chorus2 := chorus.Superpose(highHat)

	track := verse.Add(chorus).Add(verse2).Add(chorus2)

	track.Write(output, Volume)
}
