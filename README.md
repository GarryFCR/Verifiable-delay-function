# Verifiable-delay-function(Ongoing)

VDF’s are functions that require a moderate amount of sequential computation to evaluate, but once a solution is found, it is easy for anyone to verify that it is correct.This delay prevents malicious actors from influencing the output of the pseudorandom generator, since all inputs will be finalized before anyone can finish computing the VDF.
This is an implemtation of a VDF by Krzysztof Pietrzak from the paper : [Simple Verifiable Delay Functions](https://eprint.iacr.org/2018/627.pdf)

### Notes

- Generate N=pq where p and q are safe primes. p and q should be discarded for securit purposes.
- Generate signed quadratic residue modulo group(QRN+). This takes time to generate and should be generated beforehand and stored.
- x <---- $QRN+
- We take T = 2^t
- s can be {1,2,3}. s is a pre-computation parameter as mentioned in the paper.

### Implemtation

Run the Vdf as follows:

> go run run.go

### Improvement

- Implement Jacobi for testing QRN+ membership
- Reduce QRN+ generation time or implement a better method to do x <---- $QRN+
