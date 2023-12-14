#!/bin/bash

# Usage: ./gen_proto.sh <proto_dir> <gen_dir>

# Exit the script if any command fails
set -e

# Check if the correct number of arguments are provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <proto_dir> <gen_dir>"
    exit 1
fi

# Assign input arguments to variables
PROTO_DIR=$1
GEN_DIR=$2
OPENAPI_GEN_DIR=$3

# Check if the proto directory exists
if [ ! -d "$PROTO_DIR" ]; then
    echo "Proto directory not found: $PROTO_DIR"
    exit 1
fi

# Create the gen directory if it doesn't exist
mkdir -p $GEN_DIR
mkdir -p $OPENAPI_GEN_DIR

# Generate Go code from proto files
protoc --proto_path=$PROTO_DIR \
       --go_out=$GEN_DIR \
       --go_opt=paths=source_relative \
       --go-grpc_out=$GEN_DIR \
       --go-grpc_opt=paths=source_relative \
       --grpc-gateway_out=$GEN_DIR \
       --grpc-gateway_opt=paths=source_relative \
       --openapiv2_out=$OPENAPI_GEN_DIR --openapiv2_opt=allow_merge,merge_file_name=api \
       $PROTO_DIR/*.proto

echo "Proto files have been successfully compiled."

