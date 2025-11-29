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
FROM node:24-alpine
```

- **Descripción**: Se utiliza la imagen base `node:24-alpine`, que es una versión optimizada de Node.js 24 basada en Alpine Linux. Esta imagen es ligera y adecuada para la ejecución de aplicaciones en producción.

```Dockerfile
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
```

- **Descripción**:
  - Define la variable de entorno `PNPM_HOME` donde se almacenarán los binarios de pnpm.
  - Agrega este directorio al `PATH` del sistema.
  - Habilita Corepack, que viene integrado en Node.js 16+ y permite gestionar gestores de paquetes como pnpm.
  - Esto garantiza que dentro del contenedor se utilice la versión correcta de pnpm sin necesidad de instalarlo manualmente.

```Dockerfile
WORKDIR /api
```

- **Descripción**: Establece el directorio de trabajo dentro del contenedor a `/api`. A partir de este punto, todas las operaciones siguientes se realizarán dentro de este directorio.

```Dockerfile
COPY package.json pnpm-lock.yaml ./
```

- **Descripción**: Copia los archivos `package.json` y `pnpm-lock.yaml` desde tu máquina local al contenedor, ubicándolos en `/api/`. Estos archivos contienen las dependencias de la aplicación.

```Dockerfile
RUN pnpm install --prod --frozen-lockfile
```

- **Descripción**: Ejecuta `pnpm install` para instalar las dependencias definidas en `pnpm-lock.yaml`. La opción `--prod` indica que solo se instalarán las dependencias de producción, omitiendo las dependencias de desarrollo. La opción `--frozen-lockfile` asegura que el archivo de bloqueo no se modifique durante la instalación. Esto ayuda a reducir el tamaño de la imagen y mejora el rendimiento en producción.

```Dockerfile
COPY . /api
```

- **Descripción**: Copia todo el código fuente de la aplicación desde tu máquina local al contenedor, específicamente al directorio `/api`.

```Dockerfile
RUN adduser -D myuser
RUN chown -R myuser:myuser /api
USER myuser
```

- **Descripción**:
  - Crea un usuario no root llamado `myuser` para mejorar la seguridad del contenedor.
  - Cambia la propiedad de todos los archivos en el directorio `/api` al usuario `myuser`.
  - Cambia el usuario actual a `myuser`, asegurando que la aplicación se ejecute con privilegios limitados.

```Dockerfile
EXPOSE 3000
```

- **Descripción**: Informa a Docker que el contenedor escuchará en el puerto `3000`. Esto es necesario para que la aplicación pueda ser accesible desde fuera del contenedor a través de ese puerto.

```Dockerfile
CMD ["node", "src/app.js"]
```

- **Descripción**: Especifica el comando por defecto que se ejecutará cuando el contenedor inicie. En este caso, ejecuta el archivo `src/app.js` con Node.js para iniciar la aplicación.
