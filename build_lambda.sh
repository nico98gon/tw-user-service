#!/bin/bash

ZIP_FILE="user-service.zip"
OUTPUT_DIR="output"
IMAGE_NAME="user-service-lambda"

mkdir -p $OUTPUT_DIR

echo "📦 Construyendo la imagen Docker para AWS Lambda..."
docker build -t $IMAGE_NAME -f Dockerfile.lambda .

echo "🚀 Extrayendo el archivo ZIP desde el contenedor..."
CONTAINER_ID=$(docker create $IMAGE_NAME)
docker cp $CONTAINER_ID:/output/$ZIP_FILE $OUTPUT_DIR/

if [ -f "$OUTPUT_DIR/$ZIP_FILE" ]; then
  echo "✅ Archivo $ZIP_FILE generado exitosamente en $OUTPUT_DIR/"
else
  echo "❌ Error: No se pudo generar el archivo ZIP."
  exit 1
fi
