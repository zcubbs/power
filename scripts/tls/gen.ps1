param(
    [string]$outputDir = ".",
    [string]$certificateConfig = "certificate.conf"
)

# Create the output directory if it doesn't exist
if (-not (Test-Path $outputDir)) {
    New-Item -ItemType Directory -Path $outputDir | Out-Null
}

# Define the names of the certificate and key files
$certFile = Join-Path $outputDir "cert.pem"
$keyFile = Join-Path $outputDir "key.pem"

# Generate a new private key
openssl genpkey -algorithm RSA -out $keyFile

# Generate a certificate signing request (CSR) using the private key
openssl req -new -key $keyFile -out csr.pem -config $certificateConfig

# Generate a self-signed certificate using the CSR
openssl x509 -req -days 365 -in csr.pem -signkey $keyFile -out $certFile -extensions req_ext -extfile $certificateConfig

# Remove the CSR
Remove-Item csr.pem
