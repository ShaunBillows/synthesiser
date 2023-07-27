package fourier

type FourierTransform interface {
	Transform(amplitudes []float64) []complex128
	InverseTransform(frequencies []complex128) []float64
}
