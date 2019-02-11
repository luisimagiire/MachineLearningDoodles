package simulation

import (
	"gonum.org/v1/gonum/stat/distuv"
	"math/rand"
)

func GenerateMGDataset(means []float64, sigma float64, responsability float64, numObs int, seed int64) []float64{
	rand.Seed(seed)
	var gaussians = make([][]float64, len(means))

	for i := 0; i<len(means); i++{
		var tmpDist = distuv.Normal{Mu:float64(means[i]), Sigma: sigma}
		tmpvec := make([]float64, numObs)
		for j := 0; j<numObs; j++{
			tmpvec[j] = tmpDist.Rand()
		}
		gaussians[i] = tmpvec
	}

	return GenerateMixtureGaussians(responsability, gaussians)
}

func GenerateMixtureGaussians(responsability float64, gaussianData [][]float64) []float64{
	if responsability < 0 || responsability > 1 {
		panic("Responsability must be between 0 and 1!")
	}
	var ber = distuv.Bernoulli{P:responsability}

	var numObs = len(gaussianData[0])
	var finalData = make([]float64, numObs)
	for vec:=0; vec< numObs; vec ++{
		tmpDraw := ber.Rand()
		if tmpDraw == 1{
			finalData[vec] = gaussianData[0][vec]
		}else{
			finalData[vec] = gaussianData[1][vec]
		}
	}
	return finalData
}