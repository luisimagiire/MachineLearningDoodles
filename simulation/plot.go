package simulation

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func PlotLine(filePath string, vector []float64, plotRage []float64){
	pts := make(plotter.XYs, len(vector))
	for i, elem := range vector {
		if plotRage != nil {
			pts[i].X = plotRage[i]
		}else{
			pts[i].X = float64(i)
		}
		pts[i].Y = elem
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	//p.Title.Text = "Plotutil example"
	//p.X.Label.Text = "X"
	//p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,pts)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, filePath); err != nil {
		panic(err)
	}
}

func PlotDataset(fileName string, dataset []float64){
	v := make(plotter.Values, len(dataset))
	for i, elem := range dataset{
		v[i] = elem
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	h, err := plotter.NewHist(v, 16)
	if err != nil {
		panic(err)
	}
	p.Add(h)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, fileName); err != nil {
		panic(err)
	}
}