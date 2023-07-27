package output

import (
	"log"
	"os"
	"os/exec"

	"github.com/wcharczuk/go-chart/v2"
)

type Oscilloscope struct {
	FilePath string
	Title    string
}

func NewOscilloscope(path string) *Oscilloscope {
	return &Oscilloscope{
		FilePath: path,
		Title:    "",
	}
}

func (o *Oscilloscope) Write(data []float64) {
	limit := 2000
	if len(data) < 500 {
		limit = len(data)
	}

	var xs []float64
	for i := 0; i < limit; i++ {
		xs = append(xs, float64(i))
	}

	ys := data[:limit]

	graph := chart.Chart{
		Title: o.Title,
		XAxis: chart.XAxis{
			Name: "Time (samples)",
		},
		YAxis: chart.YAxis{
			Name: "Amplitude",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xs,
				YValues: ys,
			},
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top: 20,
			},
		},
	}

	f, err := os.Create(o.FilePath)
	if err != nil {
		log.Printf("Error creating graph file: %v", err)
		return
	}
	defer f.Close()

	err = graph.Render(chart.PNG, f)
	if err != nil {
		log.Printf("Error rendering graph: %v", err)
	}

	cmd := exec.Command("open", o.FilePath)
	if err := cmd.Run(); err != nil {
		log.Printf("Error opening file: %v", err)
	}
}
