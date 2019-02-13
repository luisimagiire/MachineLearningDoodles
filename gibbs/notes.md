
# Gibbs sampler sketch

We will use the backbone of the gaussian mixture simulator used in the EM algorithm.
This time we will try to converge to the true parameters sampling from the conditional
distribuition - in other words, not using the analytical solution for the expected step of
the EM algo, but direct sampling from the "frozen/marginalized" distribution.

## Remarks

The algorithm works like charm, recovering the three parameters and converging very
quickly.

Even with very little data (50 points), its shown that the algorithm converges
 to regions close from the truth parameters 
 (albeit more steps are needed for convergence).
 
