#!/bin/bash

# Default output directory
outputDir="."

# Check if an output directory was provided
if [ "$#" -eq 1 ]; then
  outputDir=$1
fi

# Define the names of the certificate and key files
certFile="${outputDir}/cert.pem"
keyFile="${outputDir}/key.pem"

# Generate a new private key
openssl genpkey -algorithm RSA -out $keyFile

# Generate a certificate signing request (CSR) using the private key
openssl req -new -key $keyFile -out csr.pem -config certificate.conf

# Generate a self-signed certificate using the CSR
openssl x509 -req -days 365 -in csr.pem -signkey $keyFile -out $certFile -extensions req_ext -extfile certificate.conf

# Remove the CSR
rm csr.pem
