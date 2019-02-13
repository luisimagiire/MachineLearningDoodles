package gibbs

import (
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat/distuv"
	"math/rand"
)

func Gibbs(dataset []float64, timeSteps int) ([]float64,[]float64,[]float64){

	// Initialize parameters
	sigmaZero := 1.5  // Lets start we fixed variance parameters and focus on finding the mean
	sigmaOne := sigmaZero

	muZero := dataset[rand.Intn(len(dataset))]
	muOne := dataset[rand.Intn(len(dataset))]
	piHat := rand.Float64()
	round := 1

	piHatEstimates := make([]float64, timeSteps+1)
	muZeroEstimates := make([]float64, timeSteps+1)
	muOneEstimates := make([]float64, timeSteps+1)

	piHatEstimates[0] = piHat
	muZeroEstimates[0] = muZero
	muOneEstimates[0] = muOne

	for ; round <= timeSteps;{

		// Simulate with fixed normals
		distOne := distuv.Normal{
			Mu:    muOneEstimates[round-1],
			Sigma: sigmaOne,
		}

		distZero := distuv.Normal{
			Mu:    muZeroEstimates[round-1],
			Sigma: sigmaZero,
		}

		tmpPiSample := GroupSampler(dataset, distZero, distOne)
		piHatEstimates[round] = floats.Sum(tmpPiSample) / float64(len(tmpPiSample))
		muZeroEstimates[round],muOneEstimates[round] = MeanSampler(tmpPiSample, dataset)
		round++
	}
	return piHatEstimates, muZeroEstimates, muOneEstimates
}

func MeanSampler(assignments []float64, dataset []float64) (float64,float64){

	var tmpTotalMeanZero []float64
	var tmpTotalMeanOne []float64

	for i := range dataset{
		if assignments[i] == 1.0{
			tmpTotalMeanZero = append(tmpTotalMeanZero, dataset[i])
		}else{
			tmpTotalMeanOne = append(tmpTotalMeanOne, dataset[i])
		}
	}

	meanZero := floats.Sum(tmpTotalMeanZero) / float64(len(tmpTotalMeanZero))
	meanOne := floats.Sum(tmpTotalMeanOne) / float64(len(tmpTotalMeanOne))

	return meanZero, meanOne
}


func GroupSampler(dataset []float64, gaussianZero distuv.Normal, gaussianOne distuv.Normal) []float64{
	tmpAssigns := make([]float64, len(dataset))
	for i, elem := range dataset{
		if gaussianZero.Prob(elem) > gaussianOne.Prob(elem){
			tmpAssigns[i] = 1
		}else{
			tmpAssigns[i] = 0
		}
	}
	return tmpAssigns
}



