
# Gibbs sampler sketch

We will use the backbone of the gaussian mixture simulator used in the EM algorithm.
This time we will try to converge to the true parameters sampling from the conditional
distribuition - in other words, not using the analytical solution for the expected step of
the EM algo, but direct sampling from the "frozen/marginalized" distribution.

