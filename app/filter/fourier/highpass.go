package fourier

import (
	"log"
)

type HighPassFilter struct {
	CutoffFrequency float64
	Resonance       float64
	FT              FourierTransform
	SampleRate      float64
}

func NewHighPassFilter(sampleRate float64, ft FourierTransform) *HighPassFilter {
	cutoffFrequency := 20.0
	resonance := 1.0
	return &HighPassFilter{
		CutoffFrequency: cutoffFrequency,
		Resonance:       resonance,
		FT:              ft,
		SampleRate:      sampleRate,
	}
}

func (f *HighPassFilter) Apply(amplitudes []float64) []float64 {
	frequencies := f.FT.Transform(amplitudes)
	N := len(amplitudes)
	for i := range frequencies {
		frequency := float64(i) * f.SampleRate / float64(N)
		if frequency < f.CutoffFrequency {
			frequencies[i] = 0
		}
	}
	return f.FT.InverseTransform(frequencies)
}

func (f *HighPassFilter) SetCutoffFrequency(frequency float64) *HighPassFilter {
	if frequency < 0 || frequency > f.SampleRate/2 {
		log.Println("Invalid cutoff frequency. Must be between 0 and Nyquist frequency (SampleRate/2)")
		return f
	}
	f.CutoffFrequency = frequency
	return f
}

func (f *HighPassFilter) SetResonance(resonance float64) *HighPassFilter {
	if resonance <= 0 {
		log.Println("Invalid resonance value. It should be greater than 0")
		return f
	}
	f.Resonance = resonance
	return f
}
