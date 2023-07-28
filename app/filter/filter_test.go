package filter

import (
	"synth/app"
	fourierFilter "synth/app/filter/fourier"
	iirFilter "synth/app/filter/iir"
	"synth/app/fourier"
	fourierMath "synth/app/fourier"
	"synth/app/output"
	"testing"
)

const (
	SampleRate float64 = 44100.0
	Volume     float64 = 1
)

func TestFilters(t *testing.T) {
	player := output.NewOtoPlayer(SampleRate)
	osc1, osc2, osc3, osc4, osc5, osc6 :=
		output.NewOscilloscope("./filter-test/input.png"),
		output.NewOscilloscope("./filter-test/fast-fourier-filter.png"),
		output.NewOscilloscope("./filter-test/iir-filter.png"),
		output.NewOscilloscope("./filter-test/fourier-filter.png"),
		output.NewOscilloscope("./filter-test/test5.png"),
		output.NewOscilloscope("./filter-test/test6.png")

	osc1.Title = "Input Signal"
	osc2.Title = "Fast Fourier Filter"
	osc3.Title = "IIR Filter"
	osc4.Title = "Fourier Filter"

	FFT := fourierMath.NewFastFourierTransformer()
	FT := fourier.NewFourierTransformer()

	lpff := fourierFilter.NewLowPassFilter(SampleRate, FFT)
	hpff := fourierFilter.NewHighPassFilter(SampleRate, FFT)

	lpf := fourierFilter.NewLowPassFilter(SampleRate, FT)
	hpf := fourierFilter.NewHighPassFilter(SampleRate, FT)

	lpiir := iirFilter.NewLowPassFilter(SampleRate)
	hpiir := iirFilter.NewHighPassFilter(SampleRate)

	note := app.NewAHarmonics(0.1, 1)

	filtered1, filtered2, filtered3, filtered4, filtered5, filtered6 := note.Copy(), note.Copy(), note.Copy(), note.Copy(), note.Copy(), note.Copy()

	filtered2.Data = lpff.SetCutoffFrequency(200).Apply(filtered2.Data)
	filtered3.Data = lpiir.SetCutoffFrequency(200).Apply(filtered3.Data)
	filtered4.Data = lpf.SetCutoffFrequency(200).Apply(filtered4.Data)

	filtered1.Write(osc1, Volume).Write(player, Volume)
	filtered2.Write(osc2, Volume).Write(player, Volume)
	filtered3.Write(osc3, Volume).Write(player, Volume)
	filtered4.Write(osc4, Volume).Write(player, Volume)

	_ = player
	_, _, _, _, _, _ = osc1, osc2, osc3, osc4, osc5, osc6
	_, _ = FFT, FT
	_, _ = lpf, hpf
	_, _ = lpff, hpff
	_, _ = lpiir, hpiir
	_, _, _, _, _, _ = filtered1, filtered2, filtered3, filtered4, filtered5, filtered6
}
