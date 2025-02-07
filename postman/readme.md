

## Variables de entorno:

Utilizar las url de AWS para probar las api sin tener que levantar la app.

{{user-service}} = http://localhost:8081 ó https://kzgxdd5pq9.execute-api.us-east-1.amazonaws.com/development
{{token}} = # Se obtiene al realizar login en endpoint de User Service
{{tweet-service}} = http://localhost:8082 ó 
{{userID}} = 679ec4e1c8e9babf6776c0af # Se obtiene al registrar un usuario
{{userIDRel}} = # ID de otro usuario para crear o leer relación

## En el endpoint de Login se puede usar el siguiente script que permite guardar el token en la variable de entorno {{token}}:

pm.test("Save the token response into the token Postman environment variable", function () {
    var jsonData = pm.response.json();
    pm.environment.set("token", jsonData.data.token);
});