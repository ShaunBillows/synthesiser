package iir

import (
	"log"
	"math"
)

type LowPassFilter struct {
	a0, a1, a2, b1, b2                     float64
	x1, x2, y1, y2                         float64
	SampleRate, CutoffFrequency, Resonance float64
}

func NewLowPassFilter(SampleRate float64) *LowPassFilter {
	cutOffFrequency := 20.0
	resonance := 1.0
	f := &LowPassFilter{
		SampleRate:      SampleRate,
		CutoffFrequency: cutOffFrequency,
		Resonance:       resonance,
	}
	return f.calculateCoefficients()
}

func (f *LowPassFilter) calculateCoefficients() *LowPassFilter {
	// Normalized cutoff frequency: 2 * pi * cutoff frequency / sampling frequency
	wc := 2 * math.Pi * f.CutoffFrequency / f.SampleRate

	// Pre-warp the frequency for the bilinear transform
	wcPreWarped := 2 * f.SampleRate * math.Tan(wc/2)

	// Coefficients for the Butterworth filter
	a := math.Sqrt(2) * wcPreWarped
	b := math.Pow(wcPreWarped, 2)

	f.a0 = b / (b + a + math.Pow(f.SampleRate, 2))
	f.a1 = 2 * f.a0
	f.a2 = f.a0
	f.b1 = 2 * (b - math.Pow(f.SampleRate, 2)) / (b + a + math.Pow(f.SampleRate, 2))
	f.b2 = (b - a + math.Pow(f.SampleRate, 2)) / (b + a + math.Pow(f.SampleRate, 2))

	return f
}

func (f *LowPassFilter) Apply(input []float64) []float64 {
	output := make([]float64, len(input))
	for i := 2; i < len(input); i++ {
		output[i] = f.a0*input[i] + f.a1*input[i-1] + f.a2*input[i-2] - f.b1*output[i-1] - f.b2*output[i-2]
	}
	return output
}

func (f *LowPassFilter) SetCutoffFrequency(frequency float64) *LowPassFilter {
	if frequency < 0 || frequency > f.SampleRate/2 {
		log.Println("Invalid cutoff frequency. Must be between 0 and Nyquist frequency (SampleRate/2)")
		return f
	}
	f.CutoffFrequency = frequency
	return f.calculateCoefficients()
}

func (f *LowPassFilter) SetResonance(Resonance float64) *LowPassFilter {
	if Resonance <= 0 {
		log.Println("Invalid Resonance value. It should be greater than 0")
		return f
	}
	f.Resonance = Resonance
	return f.calculateCoefficients()
}
