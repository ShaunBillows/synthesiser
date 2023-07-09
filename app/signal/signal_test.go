package signal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalise(t *testing.T) {
	testCases := []struct {
		name             string
		inputSignal      *Signal
		expectedSignal   []float64
	}{
		{
			name: "Can normalise positive values",
			inputSignal: &Signal{
				Data:   []float64{1, 2, 3, 4},
				SampleRate: 4,
			},
			expectedSignal: []float64{0.25, 0.5, 0.75, 1},
		},
		{
			name: "Can normalise negative values",
			inputSignal: &Signal{
				Data:   []float64{-4, -3, -2, -1},
				SampleRate: 4,
			},
			expectedSignal: []float64{-1, -0.75, -0.5, -0.25},
		},
		{
			name: "Can normalise mixed values",
			inputSignal: &Signal{
				Data:   []float64{-3, 2, -1, 4},
				SampleRate: 4,
			},
			expectedSignal: []float64{-0.75, 0.5, -0.25, 1},
		},
		{
			name: "Can normalise constant values",
			inputSignal: &Signal{
				Data:   []float64{100, 100, 100, 100},
				SampleRate: 4,
			},
			expectedSignal: []float64{1, 1, 1, 1},
		},
		{
			name: "Can normalise 0 values",
			inputSignal: &Signal{
				Data:   []float64{0, 0, 0, 0},
				SampleRate: 4,
			},
			expectedSignal: []float64{0, 0, 0, 0},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.inputSignal.Normalise()

			assert.Equal(t, tt.expectedSignal, tt.inputSignal.Data)
		})
	}
}

func TestSetVolume(t *testing.T) {
	testCases := []struct {
		name             string
		volume 			 float64
		inputSignal      *Signal
		expectedSignal   []float64
	}{
		{
			name: "Can set volume",
			volume: 0.1,
			inputSignal: &Signal{
				Data:   []float64{1, 2},
				SampleRate: 4,
			},
			expectedSignal: []float64{0.1, 0.2},
		},
		{
			name: "Cannot set volume over 0.5",
			volume: 100,
			inputSignal: &Signal{
				Data:   []float64{1},
				SampleRate: 4,
			},
			expectedSignal: []float64{0.5},
		},
		{
			name: "Cannot set negative volume",
			volume: -100,
			inputSignal: &Signal{
				Data:   []float64{1},
				SampleRate: 4,
			},
			expectedSignal: []float64{0},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.inputSignal.SetVolume(tt.volume)

			assert.Equal(t, tt.expectedSignal, tt.inputSignal.Data)
		})
	}
}

func TestSuperpose(t *testing.T) {
	testCases := []struct {
		name          string
		inputSignal   *Signal
		signalsToAdd  []*Signal
		expectedSignal []float64
	}{
		{
			name: "Cna add two signals",
			inputSignal: &Signal{
				Data:   []float64{1, -2, 3, -4},
				SampleRate: 4,
			},
			signalsToAdd: []*Signal{
				{
					Data:   []float64{1, 2, 3, -4},
					SampleRate: 4,
				},
			},
			expectedSignal: []float64{2, 0, 6, -8},
		},
		{
			name: "Can add multiple signals",
			inputSignal: &Signal{
				Data:   []float64{1, 2, 3, 4},
				SampleRate: 4,
			},
			signalsToAdd: []*Signal{
				{
					Data:   []float64{1, 2, 3, 4},
					SampleRate: 4,
				},
				{
					Data:   []float64{1, 2, 3, 4},
					SampleRate: 4,
				},
			},
			expectedSignal: []float64{3, 6, 9, 12},
		},
		{
			name: "Can add signals of different length",
			inputSignal: &Signal{
				Data:   []float64{1, 2, 3, 4},
				SampleRate: 4,
			},
			signalsToAdd: []*Signal{
				{
					Data:   []float64{1, 2},
					SampleRate: 4,
				},
			},
			expectedSignal: []float64{2, 4, 3, 4},
		},
		{
			name: "Cannot add signals with different sample rates",
			inputSignal: &Signal{
				Data:   []float64{1, 2, 3, 4},
				SampleRate: 4,
			},
			signalsToAdd: []*Signal{
				{
					Data:   []float64{1, 2, 3, 4},
					SampleRate: 5,
				},
			},
			expectedSignal: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.inputSignal.Superpose(tt.signalsToAdd...)
			assert.Equal(t, tt.expectedSignal, got.Data)
		})
	}
}

func TestWrite(t *testing.T) {

	mockOutput := NewMockOutput()

	testCases := []struct {
		name         string
		inputSignal  *Signal
		finalVolume  float64
		expectedData []float64
	}{
		{
			name: "Can write signal",
			inputSignal: &Signal{
				Data:     []float64{1, 2, 3, 4},
				SampleRate: 4,
			},
			finalVolume:  0.5,
			expectedData: []float64{0.125, 0.25, 0.375, 0.5},
		},
		{
			name: "Max volume is 0.5",
			inputSignal: &Signal{
				Data:     []float64{1, 2, 3, 4},
				SampleRate: 4,
			},
			finalVolume:  100,
			expectedData: []float64{0.125, 0.25, 0.375, 0.5},
		},
		{
			name: "Min volume is 0",
			inputSignal: &Signal{
				Data:     []float64{1, 2, 3, 4},
				SampleRate: 4,
			},
			finalVolume:  -100,
			expectedData: []float64{0, 0, 0, 0},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.inputSignal.Write(mockOutput, tt.finalVolume)
			assert.Equal(t, tt.expectedData, mockOutput.data)
		})
	}
}
