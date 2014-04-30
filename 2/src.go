package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"io/ioutil"
	"log"
	"math/big"
	"time"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("Start")
	// Create a 1024 bit RSA key.
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

	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1729),
		Subject: pkix.Name{
			CommonName:         "admin@example.com",
			Country:            []string{"India"},
			Organization:       []string{"Example Ltd."},
			OrganizationalUnit: []string{"Example Engineering"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(2, 0, 0),
		SubjectKeyId:          []byte{1, 2, 3, 4, 5, 6},
		BasicConstraintsValid: true,
		IsCA: true,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth,
			x509.ExtKeyUsageServerAuth},
		KeyUsage: x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	signed_ca_cert, err := x509.CreateCertificate(rand.Reader,
		cert, cert, &private_key.PublicKey, private_key)
	if err != nil {
		log.Println("CA cert creation failed", err)
		return
	}
	signed_ca_cert_file := "ca.pem"
	log.Println("Writing self signed CA cert to:", signed_ca_cert_file)
	ioutil.WriteFile(signed_ca_cert_file, signed_ca_cert, 0600)

	log.Println("Writing private key to: ca.key")
	ioutil.WriteFile("ca.key",
		x509.MarshalPKCS1PrivateKey(private_key),
		0600)

}
