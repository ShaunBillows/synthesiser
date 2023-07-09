package output

import (
	"log"

	"github.com/hajimehoshi/oto"
)

type OtoPlayer struct {
	p *oto.Player
}

func NewOtoPlayer(sampleRate float64) *OtoPlayer {
	channelNum := 1 // Number of channels: 1 for mono and 2 for stereo.
	audioBitDepth := 2 // Audio bit depth: 2 bytes (16 bit). This represents the range of audio signal levels.
	bufferSizeInBytes := 8192 // Size of the audio buffer. This affects latency and performance. 

	otoCtx, err := oto.NewContext(int(sampleRate), channelNum, audioBitDepth, bufferSizeInBytes)
	if err != nil {
		log.Printf("Error: Failed to initiate oto player :%v", err)
		return nil
	}
	
	player := otoCtx.NewPlayer()
	return &OtoPlayer{p: player}
}

func (p *OtoPlayer) Write(data []float64) {
	byteData := make([]byte, len(data)*2)

	for i, v := range data {
		val := int16(v * 32767.0) // Convert float64 samples in range -1.0 to 1.0 to a 16-bit integer in range -32767 to 32767.
		byteData[i*2] = byte(val)
		byteData[i*2+1] = byte(val >> 8)
	}

	_, err := p.p.Write(byteData)
	if err != nil {
		log.Printf("Error: Writing to oto player : %v", err)
	}
}

func (p *OtoPlayer) Close() {
	p.p.Close()
}
