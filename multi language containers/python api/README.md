# Python API con Docker

Este proyecto proporciona una API construida con FastAPI y empaquetada en un contenedor Docker. A continuación, se detalla cómo construir y ejecutar la aplicación utilizando Docker.

## Instrucciones para correr el proyecto con Docker

### Paso 1: Construir la imagen de Docker

Primero, construye la imagen Docker de la aplicación Python ejecutando el siguiente comando:

```bash
docker build -t my-python-app .
```

Este comando creará una imagen de Docker a partir del `Dockerfile` en el directorio actual y la etiquetará como `my-python-app`.

### Paso 2: Ejecutar la aplicación en Docker

Una vez que la imagen se haya construido correctamente, puedes ejecutar el contenedor con el siguiente comando:

```bash
docker run --name fastapi-container -p 8000:80 my-python-app
```

Este comando hará lo siguiente:

- Ejecutará el contenedor con el nombre `fastapi-container`.
- Mapeo de puertos: El puerto `80` dentro del contenedor se expondrá en el puerto `8000` de tu máquina local.
- La aplicación FastAPI será accesible a través de `http://localhost:8000` en tu navegador.

## Explicación del Dockerfile

A continuación, se explica línea por línea el contenido del `Dockerfile` de la aplicación:

```Dockerfile
FROM python:3.13-alpine
```

- **Descripción**: Se utiliza la imagen base `python:3.13-alpine`, que es una versión optimizada y ligera de Python 3.13 basada en Alpine Linux.

```Dockerfile
WORKDIR /code
```

- **Descripción**: Establece el directorio de trabajo dentro del contenedor a `/code`. A partir de este punto, todas las operaciones se realizarán dentro de este directorio.

```Dockerfile
COPY ./requirements.txt requirements.txt
```

- **Descripción**: Copia el archivo `requirements.txt` desde el directorio local al contenedor, ubicándolo en `/code/requirements.txt`.

```Dockerfile
RUN pip install --no-cache-dir -r requirements.txt
```

- **Descripción**: Instala las dependencias de Python listadas en `requirements.txt` dentro del contenedor. La opción `--no-cache-dir` evita almacenar caché de paquetes innecesarios, manteniendo la imagen más ligera.

```Dockerfile
COPY ./app ./app
```

- **Descripción**: Copia todo el código fuente de la carpeta `app` desde tu máquina local al contenedor, específicamente en el directorio `/code/app`.

```Dockerfile
EXPOSE 80
```

- **Descripción**: Informa a Docker que el contenedor escuchará en el puerto `80`. Esto es importante para la exposición del servicio web.

```Dockerfile
CMD ["fastapi", "run", "app/main.py", "--port", "80"]
```

- **Descripción**: Especifica el comando por defecto que se ejecutará cuando el contenedor inicie. En este caso, ejecuta FastAPI y lanza la aplicación en el puerto `80`.
