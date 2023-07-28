package iir

import (
	"log"
	"math"
)

type HighPassFilter struct {
	a0, a1, a2, b1, b2                     float64
	x1, x2, y1, y2                         float64
	SampleRate, CutoffFrequency, Resonance float64
}

func NewHighPassFilter(SampleRate float64) *HighPassFilter {
	cutOffFrequency := 20.0
	resonance := 1.0
	f := &HighPassFilter{
		SampleRate:      SampleRate,
		CutoffFrequency: cutOffFrequency,
		Resonance:       resonance,
	}
	return f.calculateCoefficients()
}

func (f *HighPassFilter) calculateCoefficients() *HighPassFilter {
	wc := math.Tan(math.Pi * f.CutoffFrequency / f.SampleRate) // wc is the pre-warping frequency

	f.a0 = 1.0 / (1.0 + wc/f.Resonance + wc*wc)
	f.a1 = -2 * f.a0
	f.a2 = f.a0
	f.b1 = 2.0 * (wc*wc - 1.0) * f.a0
	f.b2 = (1.0 - wc/f.Resonance + wc*wc) * f.a0
	return f
}

func (f *HighPassFilter) Apply(input []float64) []float64 {
	output := make([]float64, len(input))
	for i, x := range input {
		y := f.a0*x + f.a1*f.x1 + f.a2*f.x2 - f.b1*f.y1 - f.b2*f.y2
		f.x2, f.x1 = f.x1, x
		f.y2, f.y1 = f.y1, y
		output[i] = y
	}
	return output
}

func (f *HighPassFilter) SetCutoffFrequency(frequency float64) *HighPassFilter {
	if frequency < 0 || frequency > f.SampleRate/2 {
		log.Println("Invalid cutoff frequency. Must be between 0 and Nyquist frequency (SampleRate/2)")
		return f
	}
	f.CutoffFrequency = frequency
	return f.calculateCoefficients()
}

func (f *HighPassFilter) SetResonance(Resonance float64) *HighPassFilter {
	if Resonance <= 0 {
		log.Println("Invalid Resonance value. It should be greater than 0")
		return f
	}
	f.Resonance = Resonance
	return f.calculateCoefficients()
}
