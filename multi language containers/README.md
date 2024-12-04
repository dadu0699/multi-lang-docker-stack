# Contenedores Docker Multi-Lenguaje

Este proyecto demuestra cómo contenedorar aplicaciones utilizando Docker. Incluye ejemplos de contenedores para aplicaciones en diferentes lenguajes de programación, como Python, Java, Go y JavaScript (Node.js y Deno).

## Tabla de Contenidos

1. [¿Qué es Docker?](#qué-es-docker)
2. [Cómo construir una imagen Docker](#cómo-construir-una-imagen-docker)
3. [Cómo correr una imagen Docker](#cómo-correr-una-imagen-docker)
4. [Cómo subir una imagen a Docker Hub](#cómo-subir-una-imagen-a-docker-hub)
5. [Explicación del Dockerfile](#explicación-del-dockerfile)

## ¿Qué es Docker?

Docker es una plataforma que te permite desarrollar, enviar y ejecutar aplicaciones dentro de contenedores ligeros. Los contenedores empaquetan todo lo que la aplicación necesita para ejecutarse, incluyendo el código, el entorno de ejecución, las bibliotecas y las herramientas del sistema. Esto asegura que la aplicación se ejecute de manera consistente en cualquier entorno, ya sea la máquina de un desarrollador, un servidor de prueba o un entorno de producción.

## Cómo construir una imagen Docker

Para construir una imagen Docker desde un `Dockerfile`, debes usar el comando `docker build`.

### Paso 1: Abrir la terminal

Asegúrate de estar en el directorio donde se encuentra tu `Dockerfile`.

### Paso 2: Construir la imagen Docker

Para construir la imagen Docker, ejecuta el siguiente comando:

```bash
docker build -t <nombre-de-tu-imagen> .
```

- **`docker build`**: Indica a Docker que debe construir una nueva imagen.
- **`-t <nombre-de-tu-imagen>`**: La opción `-t` se usa para etiquetar (nombrar) la imagen. Sustituye `<nombre-de-tu-imagen>` por el nombre que desees para tu imagen.
- **`.`**: El punto (`.`) especifica el directorio actual, que es donde Docker buscará el `Dockerfile`.

Por ejemplo, si estás construyendo una imagen para la aplicación de Python, ejecuta:

```bash
docker build -t my-python-app .
```

Esto creará una imagen llamada `my-python-app`.

### Paso 3: Espera a que la imagen se construya

Docker seguirá las instrucciones dentro del `Dockerfile` para construir la imagen. Verás los logs que indican el progreso, como la descarga de dependencias y la configuración del entorno. Cuando termine, puedes verificar que la imagen fue creada ejecutando:

```bash
docker images
```

Deberías ver tu imagen listada.

## Cómo correr una imagen Docker

Una vez que la imagen esté construida, puedes ejecutarla como un contenedor utilizando el comando `docker run`.

### Paso 1: Ejecutar el contenedor Docker con nombre

Para ejecutar el contenedor y asignarle un nombre, usa el siguiente comando:

```bash
docker run --name <nombre-del-contenedor> -p <puerto-host>:<puerto-contenedor> <nombre-de-tu-imagen>
```

- **`--name <nombre-del-contenedor>`**: Esta opción asigna un nombre personalizado al contenedor. Sustituye `<nombre-del-contenedor>` con el nombre que desees para tu contenedor.
- **`-p <puerto-host>:<puerto-contenedor>`**: La opción `-p` mapea un puerto en tu máquina host a un puerto dentro del contenedor. Sustituye `<puerto-host>` con el puerto que quieras usar en tu máquina local y `<puerto-contenedor>` con el puerto al que la aplicación dentro del contenedor está escuchando.
- **`<nombre-de-tu-imagen>`**: El nombre de la imagen que deseas ejecutar.

Por ejemplo, para ejecutar la aplicación de Python (que escucha en el puerto `80` dentro del contenedor) y mapearla al puerto `8000` en tu máquina local, además de asignarle el nombre `fastapi-container`, ejecuta:

```bash
docker run --name fastapi-container -p 8000:80 my-python-app
```

Esto iniciará el contenedor con el nombre `fastapi-container` y podrás acceder a la aplicación en `http://localhost:8000` desde tu navegador o cliente API.

### Paso 2: Verifica que la aplicación esté corriendo

Una vez que el contenedor esté corriendo, abre tu navegador y visita `http://localhost:8000`. Deberías ver la aplicación FastAPI en funcionamiento (o cualquier otra aplicación con la que estés trabajando).

### Paso 3: Ver los contenedores en ejecución

Para ver los contenedores que están actualmente en ejecución, puedes usar el comando:

```bash
docker ps
```

- **`docker ps`**: Muestra una lista de todos los contenedores que están actualmente en ejecución. Verás detalles como el ID del contenedor, su nombre, el puerto que está siendo utilizado, el tiempo de ejecución, etc.

### Paso 4: Ver todos los contenedores (en ejecución y detenidos)

Si deseas ver todos los contenedores, incluyendo aquellos que han sido detenidos, puedes usar el comando:

```bash
docker ps -a
```

- **`docker ps -a`**: Muestra todos los contenedores (en ejecución y detenidos) en tu máquina. Esto es útil para verificar contenedores que se han detenido, así como para obtener información sobre su estado y tiempo de ejecución.

## Cómo subir una imagen a Docker Hub

Para compartir tu imagen con otros o desplegarla en un servidor, puedes subirla a Docker Hub, un registro en la nube para imágenes Docker.

### Paso 1: Iniciar sesión en Docker Hub

Primero, inicia sesión en Docker Hub ejecutando:

```bash
docker login
```

Se te pedirá que ingreses tu nombre de usuario y contraseña de Docker Hub. Una vez autenticado, podrás subir imágenes a tu cuenta de Docker Hub.

### Paso 2: Etiquetar la Imagen

Antes de subir la imagen, necesitas etiquetarla con tu nombre de usuario y nombre del repositorio en Docker Hub. Por ejemplo, si tu nombre de usuario en Docker Hub es `miusuario` y la imagen se llama `my-python-app`, puedes etiquetar la imagen de la siguiente manera:

```bash
docker tag my-python-app miusuario/my-python-app:latest
```

Este comando etiquetará tu imagen local `my-python-app` con la etiqueta `miusuario/my-python-app:latest`, lo que indica que pertenece a tu cuenta de Docker Hub.

### Paso 3: Subir la Imagen a Docker Hub

Ahora, puedes subir la imagen etiquetada a Docker Hub utilizando el comando `docker push`:

```bash
docker push miusuario/my-python-app:latest
```

Docker subirá la imagen a tu cuenta de Docker Hub, y puedes verificarlo visitando [https://hub.docker.com/](https://hub.docker.com/) y entrando en tu cuenta.

## Conclusión

En esta guía, aprendiste cómo:

- Construir una imagen Docker a partir de un `Dockerfile`.
- Ejecutar un contenedor Docker a partir de una imagen.
- Subir una imagen Docker a Docker Hub.

Con este conocimiento, podrás contenerizar aplicaciones en múltiples lenguajes y desplegarlas de manera consistente en diferentes entornos.
