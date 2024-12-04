# Corriendo el proyecto

## Paso 1: Construir la Aplicación Python

```bash
docker build -t my-python-app .
```

## Paso 2: Correr la Aplicación Python

Una vez que la imagen esté construida, puedes ejecutar el contenedor con el siguiente comando:

```bash
docker run --name fastapi-container -p 8000:80 my-python-app
```

Esto iniciará la aplicación FastAPI dentro del contenedor y mapeará el puerto `80` del contenedor al puerto `8000` en tu máquina local. Podrás acceder a la aplicación visitando `http://localhost:8000`.

## Explicación del Dockerfile

Vamos a repasar el `Dockerfile` de la aplicación Python. A continuación se explica cada línea:

```Dockerfile
FROM python:3.13-alpine
```

- Esta línea especifica la imagen base que usará el contenedor. Usamos la imagen `python:3.13-alpine`, que es una versión ligera de Python 3.13 basada en Alpine Linux.

```Dockerfile
WORKDIR /code
```

- Establece el directorio de trabajo dentro del contenedor como `/code`. Todas las instrucciones siguientes se ejecutarán dentro de este directorio.

```Dockerfile
COPY ./requirements.txt /code/requirements.txt
```

- Copia el archivo `requirements.txt` desde tu máquina local (el directorio actual) al directorio `/code/` dentro del contenedor.

```Dockerfile
RUN pip install --no-cache-dir --upgrade -r /code/requirements.txt
```

- Ejecuta el comando `pip install` para instalar las dependencias listadas en `requirements.txt` dentro del contenedor.

```Dockerfile
COPY ./app /code/app
```

- Copia el código fuente desde el directorio `app` en tu máquina local al directorio `/code/app` dentro del contenedor.

```Dockerfile
EXPOSE 80
```

- Indica a Docker que el contenedor escuchará en el puerto `80`. Esto es útil para cuando se ejecuta el contenedor y se expone el puerto.

```Dockerfile
CMD ["fastapi", "run", "app/main.py", "--port", "80"]
```

- La instrucción `CMD` especifica el comando predeterminado que se ejecutará cuando el contenedor inicie.
