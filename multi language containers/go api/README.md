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

Este `Dockerfile` utiliza una **estrategia multi-stage moderna**, lo que permite compilar la aplicación en una imagen base con Go y luego copiar solo el binario final a una imagen mínima optimizada para producción.

### Primera etapa: Build del binario

```Dockerfile
FROM golang:1.23-alpine AS builder
```

- **Descripción**: Usa la imagen oficial `golang:1.23-alpine` para compilar la aplicación. Está basada en Alpine, lo que la hace ligera y adecuada para tiempos de build rápidos.

```Dockerfile
RUN apk add --no-cache ca-certificates git
```

- **Descripción**: Instala dependencias mínimas necesarias para compilar el proyecto, incluyendo `git` para resolver módulos Go y `ca-certificates` para realizar descargas seguras.

```Dockerfile
WORKDIR /app
```

- **Descripción**: Define el directorio de trabajo para todas las operaciones siguientes dentro del contenedor.

```Dockerfile
COPY go.mod go.sum ./
```

- **Descripción**: Copia los archivos que definen las dependencias de Go. Esto permite que Docker aproveche el caching del módulo Go en builds posteriores.

```Dockerfile
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
```

- **Descripción**: Descarga todas las dependencias usando un caché persistente, acelerando builds repetidas.

```Dockerfile
COPY . .
```

- **Descripción**: Copia el resto del código fuente del proyecto dentro del contenedor.

```Dockerfile
ARG APP_PATH=./cmd/server
```

- **Descripción**: Define una variable de build para controlar el camino del archivo principal (`main.go`) sin modificar el Dockerfile.

```Dockerfile
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/app ${APP_PATH}
```

- **Descripción**: Compila un binario estático optimizado para producción (`-s -w` reduce tamaño). Usa Go Build cache para builds más rápidos.

### Segunda etapa: Imagen final de producción

```Dockerfile
FROM gcr.io/distroless/static-debian12:nonroot
```

- **Descripción**: Usa una imagen **Distroless** ultra mínima, sin shell ni paquetes innecesarios, ideal para producción y alta seguridad. La variante `nonroot` evita ejecutar como root.

```Dockerfile
WORKDIR /srv
```

- **Descripción**: Define el directorio desde donde se ejecutará el binario.

```Dockerfile
COPY --from=builder /out/app /srv/app
```

- **Descripción**: Copia únicamente el binario final compilado desde la fase anterior. Esto mantiene la imagen extremadamente ligera.

```Dockerfile
ENV GIN_MODE=release
```

- **Descripción**: Configura Gin para ejecutarse en modo producción.

```Dockerfile
ENV PORT=8080
```

- **Descripción**: Define el puerto por defecto que utilizará la aplicación.

```Dockerfile
EXPOSE 8080
```

- **Descripción**: Documenta que el contenedor expone el puerto `8080`.

```Dockerfile
USER nonroot:nonroot
```

- **Descripción**: Ejecuta la aplicación con un usuario no root, fortaleciendo la seguridad.

```Dockerfile
ENTRYPOINT ["/srv/app"]
```

- **Descripción**: Define el binario como punto de entrada del contenedor, permitiendo pasar argumentos adicionales con `docker run`.

## ¿Por qué usar construcción en dos etapas?

La construcción multi-stage ofrece importantes ventajas:

- **Optimización de la imagen**: La imagen final es extremadamente ligera y no incluye herramientas de compilación.
- **Mejor seguridad**: Se reduce la superficie de ataque al no incluir librerías ni paquetes innecesarios.
- **Mejor rendimiento**: Menor tamaño implica despliegues más rápidos, menor uso de almacenamiento y un bootstrap más eficiente.
