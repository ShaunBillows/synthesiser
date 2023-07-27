package fourier

import (
	"math"
)

type FourierTransformer struct{}

func NewFourierTransformer() *FourierTransformer {
	return &FourierTransformer{}
}

func (ft *FourierTransformer) Transform(amplitudes []float64) []complex128 {
	N := len(amplitudes)
	frequencyMagnitudes := make([]complex128, N)
	for k := 0; k < N; k++ {
		var frequencyMagnitude complex128
		frequencyMagnitude = 0
		for n := 0; n < N; n++ {
			theta := -2.0 * math.Pi * float64(k*n) / float64(N)
			frequencyMagnitude += complex(amplitudes[n]*math.Cos(theta), amplitudes[n]*math.Sin(theta))

		}
		frequencyMagnitudes[k] = frequencyMagnitude
	}
	return frequencyMagnitudes
}

func (ft *FourierTransformer) InverseTransform(frequencyMagnitudes []complex128) []float64 {
	N := len(frequencyMagnitudes)
	amplitudes := make([]float64, N)
	for n := 0; n < N; n++ {
		var amplitude complex128
		amplitude = 0
		for k := 0; k < N; k++ {
			theta := 2.0 * math.Pi * float64(k*n) / float64(N)
			amplitude += frequencyMagnitudes[k] * complex(math.Cos(theta), math.Sin(theta))
		}
		amplitudes[n] = real(amplitude) / float64(N)
	}
	return amplitudes
}
