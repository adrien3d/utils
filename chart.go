package utils

import (
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
	"os"
	"time"
)

func DrawChart(timestamps [][]time.Time, datas [][]float64, chartName string, seriesNames []string) {
	var timeseries []chart.Series

	for i := 0; i < len(datas); i++ {
		timeseries = append(timeseries, chart.TimeSeries{
			Name:    chartName + " " + seriesNames[i],
			XValues: timestamps[i],
			YValues: datas[i],
		})
	}

	graph := chart.Chart{
		Background: chart.Style{
			Padding:   chart.Box{Top: 50, Left: 25, Right: 25, Bottom: 10},
			FillColor: drawing.ColorFromHex("efefef"),
		},
		XAxis: chart.XAxis{Name: "Time", NameStyle: chart.StyleShow(), Style: chart.StyleShow(), ValueFormatter: chart.TimeValueFormatterWithFormat(("02/01 3:04PM"))},
		YAxis: chart.YAxis{Name: chartName, AxisType: chart.YAxisSecondary, NameStyle: chart.StyleShow(), Style: chart.StyleShow()},
		//YAxisSecondary: chart.YAxis{Name: chartName", NameStyle: chart.StyleShow(), Style: chart.StyleShow()},
		Series: timeseries,
	}
	graph.Elements = []chart.Renderable{chart.Legend(&graph)}
	file, _ := os.Create("chart-" + chartName + ".png")
	graph.Render(chart.PNG, file)
}
