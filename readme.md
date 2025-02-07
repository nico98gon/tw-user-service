
# Twitter Ualá Desafío

Este el repositorio de un sistema básico simil twitter que tiene como fin demostrar habilidades técnicas de desarrollo centrado en el lenguaje **Golang**.

## Descripción General

La versión utilizada es **1.23.4**, fue diseñado y contruído con la primicia de ser altamente escabalable, por eso se optó ir hacia una arquitectura de microservicios utilizando la plataforma de AWS para su despliegue de desarrollo y posteriormente en producción.

El sistema consta de 2 microservicios y un repositorio en común para archivoso compartidos:

1. **User Service**: Maneja el registro, autenticación y gestión de relaciones de usuarios.

2. **Tweet Service**: Se encarga de la creación, edición y eliminación de tweets.

3. **Timeline Service**: Recupera y optimiza la línea de tiempo de cada usuario.

4. **Notification Service**: Envía notificaciones a los usuarios (ej: un nuevo seguidor, un like en un tweet).

5. **Shared**: Repositorio en común para archivos compartidos.

**Estado actual**:
- ✅ User Service: Completo (registro, auth, relaciones)
- ✅ Tweet Service: Completo (CRUD tweets)
- 🚧 Timeline Service: En desarrollo
- ❌ Notification Service: Pendiente

---

**Nota Importante**:

En el desafío se intentará llegar a dividir los microservicios en repositorios individuales, pero en un principio estarán todos en un mismo repositorio. En caso de no llegar con el tiempo se dejará en un monorepo dividido en carpetas que estarán listas para ser separadas en microservicios.

---

### Servicios de AWS:

- **Lambda**: Desplegar funciones serverless que contendrán los microservicios buildeados en binarios de Go en archivos .zip
- **Api Gateway**: Api para conectar y manejar todas las peticiones y respuestas Rest con los microservicios en Lambda
- **Secret Manager**: Administrar credenciales de la base de datos
- **S3**: Buckets que permiten contener y adminstrar archivos pesados como imagenes que llaman los microservicios en lambda
- **Cloud Watch**: Visualizar logs y métricas de las lambdas

### Base de datos:

La base de datos es no relacional específicamente **MongoDB**, esto para un desarrollo rápido del desafío y por sus ventajas de flexibilidad de los datos. Pero como el desafío pide escalabilidad por si los usuarios escalan rápidamente, se pensó en una arquitectura DDD **Domain Driven Design** que permite centrarce en el dominio y así separar la lógica de negocio de las funciones de la base de datos y servicios externos. Esto es una gran ventaja si luego se requiere migrar algún microservicio a alguna base de datos relacional como **PostgreSQL** ó **MySQL**. Yo particularmente migraría el microservicio de usuarios a Postgres para mantener concistencia de los datos y evitar duplicados.


### Arquitectura

┌─────────────────┐ ┌─────────────────┐
│ API Gateway │ │ S3 Bucket │
└───────┬─┬───────┘ └───────┬──────────┘
│ │ │
│ └──────────┐ │
│ │ │
┌───────▼──────┐ ┌───▼────────┐ ┌─▼────────────┐
│ User Service │ │ Tweet │ │ Timeline │
│ (Lambda) │ │ Service │ │ Service │
└───────┬──────┘ └───┬────────┘ └─┬────────────┘
│ │ │
└─────┬──────┘ │
│ │
┌─────▼──────┐ ┌─────▼──────┐
│ MongoDB │ │ Redis │
│ (Primary) │ │ (Cache) │
└────────────┘ └────────────┘

### Estrategias de escalabilidad
- **Caché de lecturas**: Uso de Redis para almacenar timelines frecuentemente accedidos.
- **Sharding en MongoDB**: Particionado de colecciones por rangos de userID.
- **Separación de escrituras/lecturas**: Conexiones a réplicas de MongoDB para queries.

## Variables de entorno

En los microservicios existe un archivo llamado .env.example que contiene las variables de entorno de ejemplo, si no acá también están:

#### Enviroment and ports
APP_ENV=local
PORT=8081 #8082 #8083...

#### AWS
#### AWS_REGION="sa-east-1"
BUCKET_NAME= # Nombre de tu S3 Bucket 
SECRET_NAME= # Nombre de tu Secret Manager
URL_PREFIX= # Prefijo de URL API Gateway

#### BCRYPT & JWT
BCRYPT_COST=6 # 6 para local está bien, 8 para desarrollo y 10 para producción
JWT_SIGN=EsteEsEl-TokenDeNico-PARA-ENTR-uala

#### Database
DB_USERNAME=
DB_PASSWORD=
DB_HOST=
DB_DATABASE=

## Comandos de ejecución

### Requisitos previos
- Go 1.23+
- MongoDB local o en Docker
- AWS CLI configurado

### Local

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

#### Ejecutar con Docker (recomendado):

Lo primero será crear un docker-compose.yml en la ruta principal del proyecto que contiene los microservicios, luego será copiar el contenido de docker-compose-no-exec que está en la ruta principal de user-service en este nuevo docker compose.

Para poder levantar todos los microservicios deberá ejecutar:

`docker-compose up --build`

### Producción

Para poder buildear y subir nuestro .zip a lambda en AWS, deberemos ejecutar:

```
`docker build -t user-service-lambda -f Dockerfile.lambda`
`docker run --rm -v $(pwd):/output user-service-lambda cp /output/user-service.zip /output`
```

ó a través de build_lambda.sh:

- Damos permisos de ejecución: `chmod +x build_lambda.sh`
- Corremos el archivo: `./build_lambda.sh`


## Endpoints

Se dejará el archivo exportado de postman en la carpeta /postman junto con un readme que contiene las variables del mismo, y se puede acceder al team en postman desde:

https://app.getpostman.com/join-team?invite_code=a9c8ad1d529219cd5c04a27e3bc99d0ae594cac0442bf11c54b336008aeddd5d&target_code=7727bc70ac0f7b184e82cf86cd76a3f9

`Nota Importante`:

La forma más rápida y sencilla de probar la aplicación es utilizando las variables de entorno en postman para apuntar a las ya desplegadas lambdas. Pueden encontrar estas variables en los readme de la carpeta /postman.

### API (User Service)

Esta colección de Postman contiene una serie de endpoints para interactuar con el User Service, proporcionando funcionalidades para la gestión de usuarios y sus relaciones. Los endpoints están diseñados para realizar operaciones clave como registro, autenticación, actualización de perfil, gestión de avatares y banners, y relaciones entre usuarios (seguidores/seguidos). Todos los endpoints requieren autenticación mediante token (Bearer {{token}}), excepto el "register", "login", "get-profile", "get-avatar", "get-banner" y "get-follow".

Endpoints incluidos:
Profile (GET /get-profile): Obtiene el perfil de un usuario especificado por userID.
List Users (GET /get-users): Lista usuarios según tipo (new o follow) y un criterio de búsqueda opcional (search).
Register User (POST /register): Registra un nuevo usuario con email, password, name y last_name.
Login (POST /login): Inicia sesión y guarda el token en el entorno de Postman para su uso en solicitudes autenticadas.
Update User (PUT /update-profile): Actualiza el perfil del usuario con información como name, bio y location.
Upload Avatar (POST /upload-avatar): Sube un archivo de imagen como avatar del usuario.
Upload Banner (POST /upload-banner): Sube un archivo de imagen como banner del usuario.
Get Avatar (GET /get-avatar): Obtiene el avatar del usuario especificado por userID.
Get Banner (GET /get-banner): Obtiene el banner del usuario especificado por userID.
Register Relation (POST /new-relation): Registra una nueva relación entre el usuario actual y otro usuario (userIDRel).
Delete Relation (DELETE /delete-relation): Elimina una relación existente entre el usuario actual y userIDRel.
Get Relation (GET /get-relation): Verifica si existe una relación entre el usuario actual y userIDRel.
Get Following Users (GET /get-following): Obtiene la lista de usuarios que sigue el usuario especificado por userID.
Get Followers Users (GET /get-followers): Obtiene la lista de seguidores del usuario especificado por userID.

**Nota importante**:
Los endpoints utilizan variables de entorno ({{user-service}}, {{userID}}, {{token}}) para facilitar la configuración y reutilización en diferentes entornos.
Se requiere un token de autenticación para la mayoría de las operaciones, que puede obtenerse a través del endpoint de login.

### API (Tweet Service)

La colección Tweeter Service incluye una serie de endpoints para gestionar tweets dentro del sistema. Los endpoints requieren autenticación mediante un token Bearer. A continuación, se detallan las principales funcionalidades:

Create Tweet (POST /tweet)
Permite crear un nuevo tweet enviando el contenido del mensaje en el cuerpo de la solicitud.
Body (JSON):
{
  "content": "Este es mi primer tweet"
}

Read Tweets (GET /read-tweets)
Recupera los tweets de un usuario específico.
Parámetros de consulta:
id: ID del usuario.
cursor (opcional): para manejar la paginación.

Read Following Tweets (GET /following-tweets)
Devuelve los tweets de los usuarios que sigue el usuario autenticado.
Parámetro de consulta:
cursor (opcional): para manejar la paginación.

Delete Tweet (DELETE /delete-tweet)
Elimina un tweet específico identificándolo mediante su ID.

Parámetro de consulta:
id: ID del tweet a eliminar.

## Eficiencia en Timeline

La obtención de tweets es una de las aristas más cruciales del proyecto, ya que representa una gran carga de trabajo para el mismo. Vemos como podemos optimizar estas peticiones utilizando las siguientes estrategias:

#### Goroutines

Utilizar golang tiene grandes ventajas, una de ellas es la de las goroutines, esto permite manejar ciertas tareas de manera concurrente, como la obtención de tweets de diferentes usuarios o la realización de múltiples operaciones de base de datos en paralelo.

#### Caché con Redis

Esta estrategia permite implementar un sistema que puede reducir significativamente la carga en la base de datos y mejorar los tiempos de respuesta. Permite almacenar en Redis los tweets más recientes o los tweets más populares, y servir esos datos desde la caché en lugar de hacer consultas a la base de datos cada vez que se necesiten.

## Testing

Se han realizado test unitarios a las funcionalidades más críticas de la aplicación como crear usuario, crear tweet y ver el timeline de tweets.

Para poder ejecutar los test deberemos utilizar: `go test ./tests/...`

**Nota**: Siempre estar situados en la terminal sobre el microservicio que se realizará el test

# Tests críticos:
- Registro de usuario (validación contraseña, duplicados)
- Límite de 280 caracteres en tweets
- Consulta de timeline con paginación