package em

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

// Test K-means on gaussian generated values
func TestGenerateMVGaussian(t *testing.T){
	gaussianMeans := []float64{1.0, -1.0}
	sigma := mat.NewDiagDense(2, []float64{1,1})
	numObs := 10

	var dataSet = GenerateMVGaussian(gaussianMeans, sigma, numObs,uint64(99))
	fmt.Printf("2x2 - MV Gaussian:")
	for _, vec := range dataSet{
		fmt.Printf("%v \n", vec)
	}
}
