package fourier

import (
	"github.com/mjibson/go-dsp/fft"
)

type FastFourierTransformer struct{}

func NewFastFourierTransformer() *FastFourierTransformer {
	return &FastFourierTransformer{}
}

func (ft *FastFourierTransformer) Transform(amplitudes []float64) []complex128 {
	data := make([]complex128, len(amplitudes))
	for i, v := range amplitudes {
		data[i] = complex(v, 0)
	}

	return fft.FFT(data)
}

func (ft *FastFourierTransformer) InverseTransform(frequencyMagnitudes []complex128) []float64 {
	data := fft.IFFT(frequencyMagnitudes)

	result := make([]float64, len(data))
	for i, v := range data {
		result[i] = real(v)
	}

	return result
}
