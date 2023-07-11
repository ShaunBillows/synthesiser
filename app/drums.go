package app

import (
	"math"
	"synth/app/signal"
)

func NewHighHat(volume float64) *signal.Signal {
	a := 0.0  
	d := 0.01 
	s := 0.1  
	r := 0.02 

	soundDuration := a + d + r

	whiteNoise := signal.NewWhiteNoiseGenerator(soundDuration, volume, SampleRate).Generate()
	return whiteNoise.ADSR(a, d, s, r)
}

func NewKickDrum(volume float64) *signal.Signal {
	a := 0.01  
	d := 0.1   
	s := 0.0  
	r := 0.1   

	soundDuration := a + d + r
	frequency := 60.0
	sineWave := signal.NewSineWave(frequency, soundDuration, volume, SampleRate).Generate()

	for i := range sineWave.Data {
		time := float64(i) / SampleRate
		relativeTime := time / soundDuration
		currentFrequency := frequency * (1.0 - relativeTime)
		sineWave.Data[i] *= math.Sin(2.0 * math.Pi * currentFrequency * time)
	}

	return sineWave.ADSR(a, d, s, r)
}

func NewSnare(volume float64) *signal.Signal {
	a := 0.01 
	d := 0.1 
	s := 0.0  
	r := 0.1 

	soundDuration := a + d + r
	frequency := 200.0
	sineWave := signal.NewSineWave(frequency, soundDuration, volume, SampleRate).Generate()
	noise := signal.NewWhiteNoiseGenerator(soundDuration, volume, SampleRate).Generate()
	mixed := sineWave.Superpose(noise)
	return mixed.ADSR(a, d, s, r)
}
