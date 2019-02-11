package gibbs

import (
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat/distuv"
	"math/rand"
)

func Gibbs(dataset []float64, timeSteps int){

	// Initialize parameters
	//sigmaZero := stat.Variance(dataset, nil)
	sigmaZero := 1.0  // Lets start we fixed variance parameters and focus on finding the mean
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

	for ; round < timeSteps;{

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
		muZeroEstimates[round] = MeanSampler(dataset, piHatEstimates[round-1], 0)
		muOneEstimates[round] = MeanSampler(dataset, piHatEstimates[round-1], 1)

	}

}

func MeanSampler(assignments []float64, dataset []float64, distributionId int) float64{
	tmpTotalMeanZero := make([]float64, len(dataset))
	tmpTotalMeanOne := make([]float64, len(dataset))

	//for i, elem := range dataset{
	//	if ber.Rand() == 1 && distributionId == 0{
	//		tmpAssigns[i] = elem
	//	}else{
	//		tmpAssigns[i] = elem
	//	}
	//}

	for i := range dataset{
		if assignments[i] == 1{
			if distributionId == 0{
				tmpTotalMeanZero[i] = dataset[i]
			}else{
				tmpTotalMeanOne[i] = 0
			}
		}else{
			if distributionId == 1{
				tmpTotalMeanOne[i] = dataset[i]
			}else{
				tmpTotalMeanZero[i] = 0
			}
		}
	}


	return floats.Sum(tmpAssigns) / float64(len(tmpAssigns))
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



