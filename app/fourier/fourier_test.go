package fourier

import (
	"math/cmplx"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscreteFourierTransform(t *testing.T) {
	testCases := []struct {
		name                        string
		inputAmplitudes             []float64
		expectedFrequencyMagnitudes []complex128
	}{
		{
			name:                        "Can transform constant signal",
			inputAmplitudes:             []float64{3, 3, 3},
			expectedFrequencyMagnitudes: []complex128{9 + 0i, 0 + 0i, 0 + 0i},
		},
		{
			name:                        "Can transform zero signal",
			inputAmplitudes:             []float64{0, 0, 0},
			expectedFrequencyMagnitudes: []complex128{0 + 0i, 0 + 0i, 0 + 0i},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			FT := NewFourierTransformer()
			got := FT.Transform(tt.inputAmplitudes)

			for i := range got {
				expectedLowerBound := cmplx.Abs(tt.expectedFrequencyMagnitudes[i]) - 0.00001
				expectedUpperBound := cmplx.Abs(tt.expectedFrequencyMagnitudes[i]) + 0.00001

				assert.GreaterOrEqual(t, cmplx.Abs(got[i]), expectedLowerBound)
				assert.LessOrEqual(t, cmplx.Abs(got[i]), expectedUpperBound)
			}
		})
	}
}

func TestInverseDiscreteFourierTransform(t *testing.T) {
	testCases := []struct {
		name                     string
		inputFrequencyMagnitudes []complex128
		expectedAmplitudes       []float64
	}{
		{
			name:                     "Can inverse transform constant signal",
			inputFrequencyMagnitudes: []complex128{9 + 0i, 0 + 0i, 0 + 0i},
			expectedAmplitudes:       []float64{3, 3, 3},
		},
		{
			name:                     "Can inverse transform zero signal",
			inputFrequencyMagnitudes: []complex128{0 + 0i, 0 + 0i, 0 + 0i},
			expectedAmplitudes:       []float64{0, 0, 0},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			FT := NewFourierTransformer()
			got := FT.InverseTransform(tt.inputFrequencyMagnitudes)

			for i := range got {
				expectedLowerBound := tt.expectedAmplitudes[i] - 0.00001
				expectedUpperBound := tt.expectedAmplitudes[i] + 0.00001

				assert.GreaterOrEqual(t, got[i], expectedLowerBound)
				assert.LessOrEqual(t, got[i], expectedUpperBound)
			}
		})
	}
}

func TestFFT(t *testing.T) {
	testCases := []struct {
		name                        string
		inputAmplitudes             []float64
		expectedFrequencyMagnitudes []complex128
	}{
		{
			name:                        "Can transform constant signal",
			inputAmplitudes:             []float64{3, 3, 3},
			expectedFrequencyMagnitudes: []complex128{9 + 0i, 0 + 0i, 0 + 0i},
		},
		{
			name:                        "Can transform zero signal",
			inputAmplitudes:             []float64{0, 0, 0},
			expectedFrequencyMagnitudes: []complex128{0 + 0i, 0 + 0i, 0 + 0i},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			FT := NewFastFourierTransformer()
			got := FT.Transform(tt.inputAmplitudes)

			for i := range got {
				expectedLowerBound := cmplx.Abs(tt.expectedFrequencyMagnitudes[i]) - 0.00001
				expectedUpperBound := cmplx.Abs(tt.expectedFrequencyMagnitudes[i]) + 0.00001

				assert.GreaterOrEqual(t, cmplx.Abs(got[i]), expectedLowerBound)
				assert.LessOrEqual(t, cmplx.Abs(got[i]), expectedUpperBound)
			}
		})
	}
}

func TestIFFT(t *testing.T) {
	testCases := []struct {
		name                     string
		inputFrequencyMagnitudes []complex128
		expectedAmplitudes       []float64
	}{
		{
			name:                     "Can inverse transform constant signal",
			inputFrequencyMagnitudes: []complex128{9 + 0i, 0 + 0i, 0 + 0i},
			expectedAmplitudes:       []float64{3, 3, 3},
		},
		{
			name:                     "Can inverse transform zero signal",
			inputFrequencyMagnitudes: []complex128{0 + 0i, 0 + 0i, 0 + 0i},
			expectedAmplitudes:       []float64{0, 0, 0},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			FT := NewFastFourierTransformer()
			got := FT.InverseTransform(tt.inputFrequencyMagnitudes)

			for i := range got {
				expectedLowerBound := tt.expectedAmplitudes[i] - 0.00001
				expectedUpperBound := tt.expectedAmplitudes[i] + 0.00001

				assert.GreaterOrEqual(t, got[i], expectedLowerBound)
				assert.LessOrEqual(t, got[i], expectedUpperBound)
			}
		})
	}
}

func generateRandomFloats(n int) []float64 {
	data := make([]float64, n)
	for i := range data {
		data[i] = rand.Float64()
	}
	return data
}

func BenchmarkFastFourierTransform(b *testing.B) {
	fft := NewFastFourierTransformer()
	data := generateRandomFloats(44100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fft.Transform(data)
	}
}

func BenchmarkFourierTransform(b *testing.B) {
	ft := NewFourierTransformer()
	data := generateRandomFloats(44100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ft.Transform(data)
	}
}
