# Go API con Docker (Construcción en Dos Etapas)

Este proyecto utiliza Go y el framework Gin para construir una API. La aplicación se empaqueta en un contenedor Docker utilizando una estrategia de construcción en dos etapas: una para la compilación y otra para la ejecución en tiempo de ejecución. A continuación se explican los pasos para construir y ejecutar la aplicación utilizando Docker.

## Instrucciones para correr el proyecto con Docker

### Paso 1: Construir la imagen de Docker

Primero, construye la imagen Docker de la aplicación Go ejecutando el siguiente comando:

```bash
docker build -t my-go-api .
```

Este comando creará una imagen de Docker a partir del `Dockerfile` en el directorio actual y la etiquetará como `my-go-api`.

### Paso 2: Ejecutar la aplicación en Docker

Una vez que la imagen se haya construido correctamente, puedes ejecutar el contenedor con el siguiente comando:

```bash
docker run --name go-api-container -p 8080:8080 my-go-api
```

Este comando hará lo siguiente:

- Ejecutará el contenedor con el nombre `go-api-container`.
- Mapeo de puertos: El puerto `8080` dentro del contenedor se expondrá en el puerto `8080` de tu máquina local.
- La API Go será accesible a través de `http://localhost:8080` en tu navegador.

## Explicación del Dockerfile

Este `Dockerfile` utiliza una **estrategia de construcción en dos etapas**, lo que significa que se compila la aplicación en una imagen y luego se crea una imagen más pequeña y optimizada para la ejecución de la aplicación.

### Primera etapa: Construcción de la aplicación

```Dockerfile
FROM golang:1.23.4-alpine AS build
```

- **Descripción**: Utiliza la imagen `golang:1.23.4-alpine` como base para compilar la aplicación Go. Esta imagen contiene Go y está basada en Alpine Linux, lo que la hace ligera y adecuada para el entorno de desarrollo.

```Dockerfile
RUN apk add --no-cache build-base git
```

- **Descripción**: Instala las dependencias necesarias para la compilación, como `build-base` (para herramientas de construcción en Alpine) y `git` (para descargar dependencias de Go que podrían necesitarse desde repositorios).

```Dockerfile
WORKDIR /app
```

- **Descripción**: Establece el directorio de trabajo dentro del contenedor a `/app`. A partir de este punto, todas las operaciones siguientes se realizarán dentro de este directorio.

```Dockerfile
COPY go.mod go.sum ./
RUN go mod tidy && go mod download
```

- **Descripción**: Copia los archivos `go.mod` y `go.sum`, que definen las dependencias de la aplicación, y luego ejecuta los comandos `go mod tidy` y `go mod download` para descargar y asegurar que las dependencias estén disponibles.

```Dockerfile
COPY . .
```

- **Descripción**: Copia todo el código fuente de la aplicación al contenedor en el directorio `/app`.

```Dockerfile
RUN GOOS=linux GOARCH=amd64 go build -o main cmd/server/main.go
```

- **Descripción**: Compila la aplicación Go para la plataforma Linux y la arquitectura `amd64`, generando un binario llamado `main` en el directorio de trabajo.

### Segunda etapa: Creación de la imagen para ejecución en tiempo de ejecución

```Dockerfile
FROM alpine:3.21
```

- **Descripción**: Utiliza una imagen base de Alpine Linux para la segunda etapa. Esta imagen es más ligera y adecuada para ejecutar aplicaciones en producción.

```Dockerfile
RUN apk add --no-cache ca-certificates
```

- **Descripción**: Instala solo los certificados de autoridad necesarios para la ejecución de la aplicación, lo cual es esencial si la aplicación necesita hacer peticiones HTTPS.

```Dockerfile
WORKDIR /app
```

- **Descripción**: Establece el directorio de trabajo dentro del contenedor a `/app`, donde se ejecutará el binario de la aplicación.

```Dockerfile
COPY --from=build /app/main .
```

- **Descripción**: Copia el binario compilado `main` desde la fase de construcción al contenedor de ejecución.

```Dockerfile
ENV GIN_MODE=release
```

- **Descripción**: Establece la variable de entorno `GIN_MODE` en `release`, lo que configura a Gin para que ejecute la aplicación en modo de producción.

```Dockerfile
EXPOSE 8080
```

- **Descripción**: Expone el puerto `8080` en el contenedor. Esto permite que la API sea accesible a través de este puerto desde el exterior del contenedor.

```Dockerfile
CMD ["./main"]
```

- **Descripción**: Especifica el comando que se ejecutará cuando el contenedor inicie. En este caso, ejecuta el binario `main` que fue copiado desde la fase de construcción.

## ¿Por qué usar construcción en dos etapas?

La construcción en dos etapas ofrece los siguientes beneficios:

- **Optimización de la imagen**: La imagen final será más pequeña porque no incluirá herramientas de construcción como el compilador de Go ni los archivos temporales generados durante el proceso de compilación.
- **Mejor seguridad**: Minimiza la superficie de ataque, ya que el contenedor de producción solo contiene el binario de la aplicación y las dependencias necesarias para ejecutarlo.
- **Desempeño mejorado**: La imagen de producción es más ligera, lo que mejora el tiempo de inicio y reduce el consumo de recursos.
