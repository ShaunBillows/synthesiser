package app

import (
	"math"
	"synth/app/signal"
)

func NewAMinorChord(duration, volume float64) *signal.Signal {
	aFreq := 440.0
	cFreq := 440 * math.Pow(2, 3.0/12) // C5
	eFreq := 440 * math.Pow(2, 7.0/12) // E5

	outputSignal := signal.NewSineWave(aFreq, duration, volume, SampleRate).Generate()
	outputSignal = outputSignal.Superpose(
		signal.NewSineWave(cFreq, duration, volume, SampleRate).Generate(),
		signal.NewSineWave(eFreq, duration, volume, SampleRate).Generate(),
	)
	return outputSignal
}

func NewEMinorChord(duration, volume float64) *signal.Signal {
	eFreq := 440 * math.Pow(2, 7.0/12)  // E5
	gFreq := 440 * math.Pow(2, 10.0/12) // G5
	bFreq := 440 * math.Pow(2, 14.0/12) // B5

	outputSignal := signal.NewSineWave(eFreq, duration, volume, SampleRate).Generate()
	outputSignal = outputSignal.Superpose(
		signal.NewSineWave(gFreq, duration, volume, SampleRate).Generate(),
		signal.NewSineWave(bFreq, duration, volume, SampleRate).Generate(),
	)
	return outputSignal
}

func NewDMinorChord(duration, volume float64) *signal.Signal {
	dFreq := 440 * math.Pow(2, 5.0/12) // D5
	fFreq := 440 * math.Pow(2, 9.0/12) // F5
	aFreq := 440.0                     // A4

	outputSignal := signal.NewSineWave(dFreq, duration, volume, SampleRate).Generate()
	outputSignal = outputSignal.Superpose(
		signal.NewSineWave(fFreq, duration, volume, SampleRate).Generate(),
		signal.NewSineWave(aFreq, duration, volume, SampleRate).Generate(),
	)
	return outputSignal
}

func NewGMajorChord(duration, volume float64) *signal.Signal {
	gFreq := 440 * math.Pow(2, -2.0/12) // G4
	bFreq := 440 * math.Pow(2, 2.0/12)  // B4
	dFreq := 440 * math.Pow(2, 5.0/12)  // D5

	outputSignal := signal.NewSineWave(gFreq, duration, volume, SampleRate).Generate()
	outputSignal = outputSignal.Superpose(
		signal.NewSineWave(bFreq, duration, volume, SampleRate).Generate(),
		signal.NewSineWave(dFreq, duration, volume, SampleRate).Generate(),
	)
	return outputSignal
}

func NewAHarmonics(duration, volume float64) *signal.Signal {
	var outputSignal *signal.Signal
	aFreqs := []float64{
		27.5,  // A0
		55.0,  // A1
		110.0, // A2
		220.0, // A3
		440.0, // A4
		880.0, // A5
	}

	for _, freq := range aFreqs {
		if outputSignal == nil {
			outputSignal = signal.NewSineWave(freq, duration, volume, SampleRate).Generate()
		} else {
			outputSignal = outputSignal.Superpose(signal.NewSineWave(freq, duration, volume, SampleRate).Generate())
		}
	}

	return outputSignal
}
