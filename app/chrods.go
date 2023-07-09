package app

import (
	"log"
	"math"
	"synth/app/signal"
)

func BbMajor7thChord(duration, volume float64) *signal.Signal{
	bbFreq := 440 * math.Pow(2, -1.0/2) // Bb4
	dFreq  := 440 * math.Pow(2,  2.0/12) // D5
	fFreq  := 440 * math.Pow(2,  4.0/12) // F5
	aFreq  := 440 * math.Pow(2,  7.0/12) // A5
	eFreq  := 440 * math.Pow(2,  9.0/12) // E6

	outputSignal := signal.NewSineWave(bbFreq, duration, volume, SampleRate).Generate()
	outputSignal.Superpose(
		signal.NewSineWave(dFreq,  duration, volume, SampleRate).Generate(),
		signal.NewSineWave(fFreq,  duration, volume, SampleRate).Generate(),
		signal.NewSineWave(aFreq,  duration, volume, SampleRate).Generate(),
		signal.NewSineWave(eFreq,  duration, volume, SampleRate).Generate(),
	)
	log.Printf("Playing Bb Major 7th Chord.")
	return outputSignal
}

func A7Chord(duration, volume float64) *signal.Signal{
	aFreq := 440.0 // A4
	csFreq  := 440 * math.Pow(2,  4.0/12) // C#5
	eFreq  := 440 * math.Pow(2,  7.0/12) // E5
	gFreq  := 440 * math.Pow(2, 10.0/12) // G5
	cFreq  := 440 * math.Pow(2,  3.0/12) // C5

	outputSignal := signal.NewSineWave(aFreq, duration, volume, SampleRate).Generate()
	outputSignal.Superpose(
		signal.NewSineWave(csFreq,  duration, volume, SampleRate).Generate(),
		signal.NewSineWave(eFreq,  duration, volume, SampleRate).Generate(),
		signal.NewSineWave(gFreq,  duration, volume, SampleRate).Generate(),
		signal.NewSineWave(cFreq,  duration, volume, SampleRate).Generate(),
	)
	log.Printf("Playing A7 Chord.")
	return outputSignal
}

func DMinor7Chord(duration, volume float64) *signal.Signal{
	dFreq := 440 * math.Pow(2, -7.0/12) // D4
	fFreq := 440 * math.Pow(2, -5.0/12) // F4
	aFreq := 440.0 // A4
	cFreq := 440 * math.Pow(2, 3.0/12)  // C5

	outputSignal := signal.NewSineWave(dFreq, duration, volume, SampleRate).Generate()
	outputSignal.Superpose(
		signal.NewSineWave(fFreq, duration, volume, SampleRate).Generate(),
		signal.NewSineWave(aFreq, duration, volume, SampleRate).Generate(),
		signal.NewSineWave(cFreq, duration, volume, SampleRate).Generate(),
	)
	log.Printf("Playing D Minor 7th Chord.")
	return outputSignal
}
