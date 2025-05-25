#!/bin/bash

set -e

OUTPUT=${1:-bin/app}

echo "📦 Gerando build para $OUTPUT..."

mkdir -p $(dirname "$OUTPUT")
go build -o "$OUTPUT" ./cmd/main.go

echo "✅ Build gerado em $OUTPUT"
