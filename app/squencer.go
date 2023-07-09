package app

import (
	"math"
	"synth/app/signal"
)

func Sequencer(sound *signal.Signal, duration float64, bpm float64, offBeat bool) *signal.Signal {
	beatDuration := 60.0 / bpm

	totalBeats := int(duration / beatDuration)

	sequence := &signal.Signal{
		Data:     []float64{}, 
		SampleRate: SampleRate,
	}

	if offBeat {
		silenceSamples := int((beatDuration / 2) * SampleRate)
		silence := &signal.Signal{
			Data:     make([]float64, silenceSamples),
			SampleRate: SampleRate,
		}

		sequence = sequence.Add(silence)
	}

	for i := 0; i < totalBeats; i++ {
		sequence = sequence.Add(sound)

		silenceDuration := beatDuration - float64(len(sound.Data)) / SampleRate
		silenceDuration = math.Max(silenceDuration, 0)

		silenceSamples := int(silenceDuration * SampleRate)
		silence := &signal.Signal{
			Data:     make([]float64, silenceSamples),
			SampleRate: SampleRate,
		}

		sequence = sequence.Add(silence)
	}

	return sequence
}

