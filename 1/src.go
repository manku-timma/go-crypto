package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("Start")
	private_key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Println("Private key creation failed: ", err)
		return
	}
	log.Println("Public key: modulus (N): ", private_key.N)
	log.Println("Public key: exponent (E): ", private_key.E)
	log.Println("Private key: exponent (D): ", private_key.D)
	log.Println("Private key: prime factors of public key modulus (P,Q): ", private_key.Primes)
	log.Println("Private key: precomputed: D mod P-1", private_key.Precomputed.Dp)
	log.Println("Private key: precomputed: D mod Q-1", private_key.Precomputed.Dq)
	log.Println("Private key: precomputed: Q^-1 mod P", private_key.Precomputed.Qinv)
	log.Println("Private key: precomputed: CRTvalues",
		private_key.Precomputed.CRTValues)
}
