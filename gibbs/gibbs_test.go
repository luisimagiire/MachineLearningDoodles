package gibbs

import (
	"machineLearningBB/simulation"
	"testing"
)

func TestGibbs(t *testing.T) {
	means := []float64{-3,3}
	numObs := 1000
	sigma := 1.5
	responsability := 0.7
	maxRounds := 200
	seed := int64(1)

	dataset := simulation.GenerateMGDataset(means, sigma, responsability, numObs, seed)



}
