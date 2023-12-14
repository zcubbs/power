# gen_proto.ps1
param (
  [string]$ProtoDir,
  [string]$ProtoGenDir,
  [string]$OpenapiGenDir
)

# Exit if the script encounters any error
$ErrorActionPreference = "Stop"

# Check if the proto directory exists
if (-not (Test-Path -Path $ProtoDir -PathType Container)) {
  Write-Error "Proto directory not found: $ProtoDir"
  exit 1
}

# Create the gen directory if it doesn't exist
if (-not (Test-Path -Path $ProtoGenDir)) {
  New-Item -ItemType Directory -Path $ProtoGenDir | Out-Null
}

# Generate Go code from proto files
$protoFile = $ProtoDir + "\*.proto"
&protoc --proto_path=$ProtoDir `
          --go_out=$ProtoGenDir `
          --go_opt=paths=source_relative `
          --go-grpc_out=$ProtoGenDir `
          --go-grpc_opt=paths=source_relative `
          --grpc-gateway_out=$ProtoGenDir `
          --grpc-gateway_opt=paths=source_relative `
          --openapiv2_out=$OpenapiGenDir --openapiv2_opt=allow_merge=true,merge_file_name=api `
          $protoFile
if ($LASTEXITCODE -ne 0) {
  Write-Error "Failed to generate code from $protoFile"
  exit 1
}

Write-Host "Proto files have been successfully compiled."
