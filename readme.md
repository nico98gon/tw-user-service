
# Twitter Ualá Desafío

Este el repositorio de un sistema básico simil twitter que tiene como fin demostrar habilidades técnicas de desarrollo centrado en el lenguaje **Golang**.

## Descripción General

La versión utilizada es **1.23.4** Fue diseñado y contruído con la primicia de ser altamente escabalable, por eso e optó ir hacia una arquitectura de microservicios utilizando la plataforma de AWS para su despliegue de desarrollo y posteriormente en producción.

El sistema consta de 4 microservicios y un repositorio en común para archivoso compartidos:

1. **User Service**: Maneja el registro, autenticación y gestión de relaciones de usuarios.

2. **Tweet Service**: Se encarga de la creación, edición y eliminación de tweets.

3. **Timeline Service**: Recupera y optimiza la línea de tiempo de cada usuario.

4. **Notification Service**: Envía notificaciones a los usuarios (ej: un nuevo seguidor, un like en un tweet).

5. **Shared**: Repositorio en común para archivos compartidos.

---

**Nota Importante**:

En el desafío se intentará llegar a dividir los microservicios en repositorios individuales, pero en un principio estarán todos en un mismo repositorio. En caso de no llegar con el tiempo se dejará en un monorepo dividido en carpetas que estarán listas para ser separadas en microservicios.

---

### Servicios de AWS:

- **Lambda**: Desplegar funciones ejectuables que contendrán los microservicios buildeados en binarios de Go en archivos .zip
- **Api Gateway**: Api que conectar y manejar todas las peticiones y respuestas Rest con los microservicios en Lambda
- **Secret Manager**: Administrar credenciales de la base de datos
- **S3**: Buckets que permiten contener y adminstrar archivos pesados como imagenes que llaman los microservicios en lambda
- **Cloud Watch**: Visualizar logs y métricas de nuestros lambdas

### Base de datos:

La base de datos es no relacional específicamente **MongoDB**, esto para un desarrollo rápido del desafío y por sus ventajas de flexibilidad de los datos. Pero como el desafío pide escalabilidad por si los usuarios escalan rápidamente, se pensó en una arquitectura DDD **Domain Driven Design** que permite centrarce en el dominio y así separar la lógica de negocio de las funciones de la base de datos y servicios externos. Esto es una gran ventaja si luego se requiere migrar algún microservicio a alguna base de datos relacional como **PostgreSQL** ó **MySQL**. Yo particularmente migraría el microservicio de usuarios a Postgres para mantener concistencia de los datos y evitar duplicados.

## Comandos de ejecución

#### Local

Para poder correr la aplicación en local se debe instalar y configurar el CLI de AWS:

``````
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
``````

Configurar credenciales:

`aws configure`

Verificar que las credenciales están configuradas:

`cat ~/.aws/credentials`

#### Producción

Para poder buildear y subir nuestro .zip a lambda en AWS, deberemos ejecutar:

```
`docker build -t user-service-lambda -f Dockerfile.lambda`
`docker run --rm -v $(pwd):/output user-service-lambda cp /output/user-service.zip /output`
```

ó a través de build_lambda.sh:

- Damos permisos de ejecución: `chmod +x build_lambda.sh`
- Corremos el archivo: `./build_lambda.sh`