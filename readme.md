

### Comandos de ejecución

#### Local

#### Producción

Para poder subir buildear y subir nuestro .zip a lambda en AWS, deberemos ejecutar:

```
- Crear la imagen con el binario .zip dentro: `docker build -t user-service-lambda -f Dockerfile.lambda`
- Extraer el .zip desde el contenedor a la computadora: `docker run --rm -v $(pwd):/output user-service-lambda cp /output/user-service.zip /output`
```

ó a través de build_lambda.sh:

- Damos permisos de ejecución: `chmod +x build_lambda.sh`
- Corremos el archivo: `./build_lambda.sh`