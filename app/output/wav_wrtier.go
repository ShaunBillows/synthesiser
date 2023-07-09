package output

import (
	"log"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

type WavWriter struct {
	w          *wav.Encoder
	sampleRate int
}

func NewWavWriter(path string, sampleRate float64, appendMode bool) *WavWriter {
	file, err := os.Create(path)
	if err != nil {
		log.Panicf("Error: Failed to create WAV file :%v", err)
		return nil
	}

	enc := wav.NewEncoder(file, int(sampleRate), 16, 1, 1)
	return &WavWriter{w: enc, sampleRate: int(sampleRate)}
}

func (p *WavWriter) Write(data []float64) {
	intData := make([]int, len(data))

	for i, v := range data {
		intData[i] = int(v * 32767.0)
	}
	
	buffer := &audio.IntBuffer{Data: intData, Format: &audio.Format{SampleRate: p.sampleRate, NumChannels: 1}}
	err := p.w.Write(buffer)
	if err != nil {
		log.Printf("Error: Writing to file : %v", err)
	}
}

func (p *WavWriter) Close() {
	p.w.Close()
}
