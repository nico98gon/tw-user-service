#!/bin/bash

ZIP_FILE="user-service.zip"
OUTPUT_DIR="output"
BIN_FILE="bootstrap"

mkdir -p $OUTPUT_DIR

if [ ! -f "$BIN_FILE" ]; then
  echo "🔨 Compilando el binario para AWS Lambda..."

  env GOOS=linux GOARCH=arm64 go build -o bootstrap ./cmd/main.go

  if [ ! -f "$BIN_FILE" ]; then
    echo "❌ Error: No se pudo compilar el binario $BIN_FILE."
    exit 1
  fi
fi

if [ -f "$OUTPUT_DIR/$ZIP_FILE" ]; then
  echo "🧹 Borrando el archivo ZIP anterior..."
  rm "$OUTPUT_DIR/$ZIP_FILE"
fi

echo "📦 Creando el archivo ZIP..."
zip -j "$OUTPUT_DIR/$ZIP_FILE" "$BIN_FILE"
rm "$BIN_FILE"

if [ -f "$OUTPUT_DIR/$ZIP_FILE" ]; then
  echo "✅ Archivo $ZIP_FILE generado exitosamente en $OUTPUT_DIR/"
else
  echo "❌ Error: No se pudo generar el archivo ZIP."
  exit 1
fi
