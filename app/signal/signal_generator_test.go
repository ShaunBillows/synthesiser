package signal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSineWaveGenerate(t *testing.T) {
    testCases := []struct {
        name             	 string
        generator        	 *SineWaveGenerator
        expectedLength       int
        expectedFirstSample  float64
        expectedLastSample   float64
        expectedMaxSample    float64
        expectedMinSample    float64
    }{
        {
            name:             "Can make a sine wave",
            generator:        NewSineWave(1, 1, 1, 4),
            expectedLength:       4,
            expectedFirstSample:  0,
            expectedLastSample:   -1,
            expectedMaxSample:    1,
            expectedMinSample:   -1,
        },
    }

    for _, tt := range testCases {
        t.Run(tt.name, func(t *testing.T) {
            got := tt.generator.Generate()

            assert.Equal(t, tt.expectedLength, len(got.Signal), "Expected length of 4")
            assert.Equal(t, tt.generator.SampleRate, got.SampleRate, "Expected original sample rate")
            assert.Equal(t, tt.expectedFirstSample, got.Signal[0], "Expect to start at 0")
			assert.Equal(t, tt.expectedLastSample, got.Signal[3], "Expect to end at -1")
			assert.Equal(t, tt.expectedMaxSample, got.Signal[1], "Expect max value to be at first index")
			assert.Equal(t, tt.expectedMinSample, got.Signal[3], "Expect min value to be at last index")
        })
    }
}

func TestSawtoothWaveGenerate(t *testing.T) {
    testCases := []struct {
        name             	 string
        generator        	 *SawtoothWaveGenerator
        expectedLength       int
        expectedFirstSample  float64
        expectedLastSample   float64
        expectedMaxSample    float64
        expectedMinSample    float64
    }{
        {
            name:             "Can make a sawtooth wave",
            generator:        NewSawtoothWave(1, 1, 1, 4),
            expectedLength:       4,
            expectedFirstSample:  0,
            expectedLastSample:   -0.5,
            expectedMaxSample:    0.5,
            expectedMinSample:   -1,
        },
    }

    for _, tt := range testCases {
        t.Run(tt.name, func(t *testing.T) {
            got := tt.generator.Generate()

            assert.Equal(t, tt.expectedLength, len(got.Signal), "Expect length of 4")
            assert.Equal(t, tt.generator.SampleRate, got.SampleRate, "Expect original sample rate")
            assert.Equal(t, tt.expectedFirstSample, got.Signal[0], "Expect to start at 0")
			assert.Equal(t, tt.expectedLastSample, got.Signal[3], "Expect to end at -0.5")
			assert.Equal(t, tt.expectedMaxSample, got.Signal[1], "Expect max value to be at first index")
			assert.Equal(t, tt.expectedMinSample, got.Signal[2], "Expect min value to be at second index")
        })
    }
}

func TestSquareWaveGenerate(t *testing.T) {
    testCases := []struct {
        name             	 string
        generator        	 *SquareWaveGenerator
        expectedLength       int
        expectedFirstSample  float64
        expectedLastSample   float64
        expectedMaxSample    float64
        expectedMinSample    float64
    }{
        {
            name:             "Can make a square wave",
            generator:        NewSquareWave(1, 1, 1, 4),
            expectedLength:       4,
            expectedFirstSample:  1,
            expectedLastSample:   -1,
            expectedMaxSample:    1,
            expectedMinSample:   -1,
        },
    }

    for _, tt := range testCases {
        t.Run(tt.name, func(t *testing.T) {
            got := tt.generator.Generate()

            assert.Equal(t, tt.expectedLength, len(got.Signal), "Expect length of 4")
            assert.Equal(t, tt.generator.SampleRate, got.SampleRate, "Expect original sample rate")
            assert.Equal(t, tt.expectedFirstSample, got.Signal[0], "Expect to start at 1")
			assert.Equal(t, tt.expectedLastSample, got.Signal[3], "Expect to end at -1")
			assert.Equal(t, tt.expectedMaxSample, got.Signal[0], "Expect max value to be at first index")
			assert.Equal(t, tt.expectedMinSample, got.Signal[3], "Expect min value to be at last index")
        })
    }
}
