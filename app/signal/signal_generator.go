package signal

import (
	"math"
	"math/rand"
	"time"
)

// Sine wave

type SineWaveGenerator struct {
	Frequency float64
	Duration  float64
	Volume    float64
	SampleRate float64
}

func NewSineWave(frequency float64, duration float64, volume float64, sampleRate float64) *SineWaveGenerator {
	return &SineWaveGenerator{Frequency: frequency, Duration: duration, Volume: volume, SampleRate: sampleRate}
}

func (g *SineWaveGenerator) Generate() *Signal {
	samples := g.SampleRate * g.Duration

	signalData := make([]float64, int(samples))

	for sampleNo := range signalData {
		time := float64(sampleNo)/g.SampleRate
		signalData[sampleNo] = g.Volume * math.Sin(2 * math.Pi * g.Frequency * time)
	}

	return &Signal{
		Data:       signalData,
		SampleRate: g.SampleRate,
	}
}

// Sawtooth wave

type SawtoothWaveGenerator struct {
	Frequency float64
	Duration  float64
	Volume    float64
	SampleRate float64
}

func NewSawtoothWave(frequency float64, duration float64, volume float64, sampleRate float64) *SawtoothWaveGenerator {
	return &SawtoothWaveGenerator{Frequency: frequency, Duration: duration, Volume: volume, SampleRate: sampleRate}
}

func (g *SawtoothWaveGenerator) Generate() *Signal {
	samples := g.SampleRate * g.Duration
	signalData := make([]float64, int(samples))

	for sampleNo := range signalData {
		time := float64(sampleNo) / g.SampleRate
		signalData[sampleNo] = g.Volume * 2 * (time*g.Frequency - math.Floor(0.5+time*g.Frequency))
	}

	return &Signal{
		Data:       signalData,
		SampleRate: g.SampleRate,
	}
}

// Square wave

type SquareWaveGenerator struct {
	Frequency float64
	Duration  float64
	Volume    float64
	SampleRate float64
}

func NewSquareWave(frequency float64, duration float64, volume float64, sampleRate float64) *SquareWaveGenerator {
	return &SquareWaveGenerator{Frequency: frequency, Duration: duration, Volume: volume, SampleRate: sampleRate}
}

func (g *SquareWaveGenerator) Generate() *Signal {
	samples := g.SampleRate * g.Duration
	signalData := make([]float64, int(samples))

	for sampleNo := range signalData {
		time := float64(sampleNo) / g.SampleRate
		sample := math.Sin(2 * math.Pi * g.Frequency * time)
		if sample >= 0 {
			signalData[sampleNo] = g.Volume
		} else {
			signalData[sampleNo] = -g.Volume
		}
	}
	return &Signal{
		Data:       signalData,
		SampleRate: g.SampleRate,
	}
}

// White noise

type WhiteNoiseGenerator struct {
	Duration  float64
	Volume    float64
	SampleRate float64
}

func NewWhiteNoiseGenerator(duration, volume float64, sampleRate float64) *WhiteNoiseGenerator {
	return &WhiteNoiseGenerator{Duration: duration, Volume: volume, SampleRate: sampleRate}
}

func (g *WhiteNoiseGenerator) Generate() *Signal {
	totalSamples := int(g.SampleRate * g.Duration)
	signalData := make([]float64, totalSamples)
	rand.Seed(time.Now().UnixNano())

	for i := range signalData {
		signalData[i] = g.Volume * (2*rand.Float64() - 1)
	}

	return &Signal{
		Data:       signalData,
		SampleRate: g.SampleRate,
	}
}