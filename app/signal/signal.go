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
	Data       []float64
	SampleRate float64
}

func NewSignal(generator SignalGenerator) *Signal {
	return generator.Generate()
}

func (s *Signal) Normalise() {
	maxAmplitude := 0.0

	for _, sample := range s.Data {
		absSample := math.Abs(sample)
		if absSample > maxAmplitude {
			maxAmplitude = absSample
		}
	}

	if maxAmplitude == 0 {
		return
	}

	for i, sample := range s.Data {
		normalizedSample := sample / maxAmplitude
		s.Data[i] = normalizedSample
	}
}

func (s *Signal) SetVolume(volume float64) {
	if volume > 0.5 {
		volume = 0.5 // Warning: Amplitudes exceeding this value are extremely loud!
	}
	if volume < 0 {
		volume = 0
	}

	for i, sample := range s.Data {
		s.Data[i] = sample * volume
	}
}

func (s *Signal) Superpose(signals ...*Signal) *Signal {
	newSignal := make([]float64, len(s.Data))

	for _, signal := range signals {
		if signal.SampleRate != s.SampleRate {
			log.Printf("Error: Cannot add signals with different sample rates.")
			return &Signal{}
		}
	}

	for i := range newSignal {
		newSignal[i] = s.Data[i]
		for _, signal := range signals {
			if i < len(signal.Data) {
				newSignal[i] += signal.Data[i]
			}
		}
	}

	return &Signal{
		Data:       newSignal,
		SampleRate: s.SampleRate,
	}
}

func (s *Signal) Write(output Output, finalVolume float64) *Signal {
	s.Normalise()
	s.SetVolume(finalVolume)
	output.Write(s.Data)
	return s
}

func (s *Signal) ADSR(attackTime, decayTime, sustainLevel, releaseTime float64) *Signal {
	maxAmplitude := 0.0

	for _, sample := range s.Data {
		absSample := math.Abs(sample)
		if absSample > maxAmplitude {
			maxAmplitude = absSample
		}
	}

	if maxAmplitude == 0.0 {
		return s
	}

	totalSamples := len(s.Data)

	for i, sample := range s.Data {
		t := float64(i) / s.SampleRate

		if t < attackTime {
			s.Data[i] = (maxAmplitude / attackTime) * t * sample
		} else if t < attackTime+decayTime {
			s.Data[i] = (maxAmplitude - ((maxAmplitude-sustainLevel)/decayTime)*(t-attackTime)) * sample
		} else if t < float64(totalSamples)/s.SampleRate-releaseTime {
			s.Data[i] = sustainLevel * sample
		} else {
			s.Data[i] = (sustainLevel - (sustainLevel/releaseTime)*(t-float64(totalSamples)/s.SampleRate)) * sample
		}
	}
	return s
}

func (s *Signal) Add(signal *Signal) *Signal {
	newSignal := make([]float64, len(s.Data))
	copy(newSignal, s.Data)
	newSignal = append(newSignal, signal.Data...)
	return &Signal{
		Data:       newSignal,
		SampleRate: s.SampleRate,
	}
}

func (s *Signal) Copy() *Signal {
	newData := make([]float64, len(s.Data))
	copy(newData, s.Data)
	return &Signal{
		Data:       newData,
		SampleRate: s.SampleRate,
	}
}
