package gibbs

import (
	"machineLearningBB/simulation"
	"testing"
)

func TestGibbs(t *testing.T) {
	means := []float64{-3,3}
	numObs := 50
	sigma := 1.5
	responsability := 0.7
	maxRounds := 2000
	seed := int64(1)

	dataset := simulation.GenerateMGDataset(means, sigma, responsability, numObs, seed)
	simulation.PlotDataset("dataset.png", dataset)
	piHatEstimates, muZeroEstimates, muOneEstimates := Gibbs(dataset, maxRounds)

	simulation.PlotLine("piHatTrace.png", piHatEstimates, nil)
	simulation.PlotLine("muZeroTrace.png", muZeroEstimates, nil)
	simulation.PlotLine("muOneTrace.png", muOneEstimates, nil)

}
