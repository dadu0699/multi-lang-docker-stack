# Node.js API con Docker

Este proyecto proporciona una API construida con Node.js y empaquetada en un contenedor Docker. A continuación, se detallan los pasos para construir y ejecutar la aplicación utilizando Docker.

## Instrucciones para correr el proyecto con Docker

### Paso 1: Construir la imagen de Docker

Primero, construye la imagen Docker de la aplicación Node.js ejecutando el siguiente comando:

```bash
docker build -t my-node-app .
```

Este comando creará una imagen de Docker a partir del `Dockerfile` en el directorio actual y la etiquetará como `my-node-app`.

### Paso 2: Ejecutar la aplicación en Docker

Una vez que la imagen se haya construido correctamente, puedes ejecutar el contenedor con el siguiente comando:

```bash
docker run --name node-container -p 3000:3000 my-node-app
```

Este comando hará lo siguiente:

- Ejecutará el contenedor con el nombre `node-container`.
- Mapeo de puertos: El puerto `3000` dentro del contenedor se expondrá en el puerto `3000` de tu máquina local.
- La aplicación Node.js será accesible a través de `http://localhost:3000` en tu navegador.

## Explicación del Dockerfile

A continuación, se explica línea por línea el contenido del `Dockerfile` de la aplicación Node.js:

```Dockerfile
FROM node:22-alpine
```

- **Descripción**: Se utiliza la imagen base `node:22-alpine`, que es una versión optimizada de Node.js 22 basada en Alpine Linux. Esta imagen es ligera y adecuada para la ejecución de aplicaciones en producción.

```Dockerfile
WORKDIR /api
```

- **Descripción**: Establece el directorio de trabajo dentro del contenedor a `/api`. A partir de este punto, todas las operaciones siguientes se realizarán dentro de este directorio.

```Dockerfile
COPY package.json package-lock.json /api/
```

- **Descripción**: Copia los archivos `package.json` y `package-lock.json` desde tu máquina local al contenedor, ubicándolos en `/api/`. Estos archivos contienen las dependencias de la aplicación.

```Dockerfile
RUN npm ci --omit=dev
```

- **Descripción**: Ejecuta `npm ci` para instalar las dependencias definidas en `package-lock.json`. La opción `--omit=dev` indica que solo se instalarán las dependencias de producción, omitiendo las dependencias de desarrollo. Esto ayuda a reducir el tamaño de la imagen y mejora el rendimiento en producción.

```Dockerfile
COPY . /api
```

- **Descripción**: Copia todo el código fuente de la aplicación desde tu máquina local al contenedor, específicamente al directorio `/api`.

```Dockerfile
EXPOSE 3000
```

- **Descripción**: Informa a Docker que el contenedor escuchará en el puerto `3000`. Esto es necesario para que la aplicación pueda ser accesible desde fuera del contenedor a través de ese puerto.

```Dockerfile
CMD ["node", "src/app.js"]
```

- **Descripción**: Especifica el comando por defecto que se ejecutará cuando el contenedor inicie. En este caso, ejecuta el archivo `src/app.js` con Node.js para iniciar la aplicación.
