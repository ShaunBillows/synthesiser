package signal

import (
	"log"
	"math"
)

type Output interface {
	Write(signal []float64)
}

type SignalGenerator interface {
	Generate() *Signal
}

type Signal struct {
	Signal   []float64
	SampleRate float64
}

func NewSignal(generator SignalGenerator) *Signal {
	return generator.Generate()
}

func (s *Signal) Normalise() {
	maxAmplitude := 0.0

	for _, sample := range s.Signal {
		absSample := math.Abs(sample)
		if absSample > maxAmplitude {
			maxAmplitude = absSample
		}
	}

	if maxAmplitude == 0 {
		return
	}

	for i, sample := range s.Signal {
		normalizedSample := sample / maxAmplitude
		s.Signal[i] = normalizedSample
	}
}

func (s *Signal) SetVolume(volume float64) {
	if volume > 0.5 {
		volume = 0.5 // Warning: Amplitudes exceeding this value are extremely loud!
	}
	if volume < 0 {
		volume = 0
	}

    for i, sample := range s.Signal {
        s.Signal[i] = sample * volume
    }
}

func (s *Signal) Superpose(signals ...*Signal) *Signal {
	newSignal := make([]float64, len(s.Signal))

	for _, signal := range signals {
		if signal.SampleRate != s.SampleRate {
			log.Printf("Error: Cannot add signals with different sample rates.")
			return &Signal{}
		}
	}

	for i := range newSignal {
		newSignal[i] = s.Signal[i]
		for _, signal := range signals {
			if i < len(signal.Signal) {
				newSignal[i] += signal.Signal[i]
			}
		}
	}

	s.Signal = newSignal
	return &Signal{
		Signal:       newSignal,
		SampleRate: s.SampleRate,
	}
}

func (s *Signal) Write(output Output, finalVolume float64) *Signal {
    s.Normalise()
    s.SetVolume(finalVolume)
	output.Write(s.Signal)
	return s
}
