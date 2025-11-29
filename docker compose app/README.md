# Aplicación Docker Compose

Este proyecto demuestra cómo usar Docker Compose para orquestar múltiples contenedores Docker. Incluye ejemplos de cómo definir y ejecutar servicios en un archivo `docker-compose.yml`.

## Tabla de Contenidos

1. [¿Qué es Docker Compose?](#qué-es-docker-compose)
2. [Cómo definir servicios en Docker Compose](#cómo-definir-servicios-en-docker-compose)
3. [Cómo correr una aplicación con Docker Compose](#cómo-correr-una-aplicación-con-docker-compose)
4. [Cómo detener y eliminar contenedores](#cómo-detener-y-eliminar-contenedores)
5. [Explicación del archivo docker-compose.yml](#explicación-del-archivo-docker-composeyml)

## ¿Qué es Docker Compose?

Docker Compose es una herramienta para definir y ejecutar aplicaciones Docker multi-contenedor. Con Docker Compose, puedes usar un archivo YAML para configurar los servicios de tu aplicación. Luego, con un solo comando, puedes crear e iniciar todos los servicios desde tu configuración.

## Cómo definir servicios en Docker Compose

Para definir servicios en Docker Compose, necesitas crear un archivo `docker-compose.yml` en el directorio raíz de tu proyecto. Aquí hay un ejemplo básico:

```yaml
name: web_app_example
services:
  web:
    image: nginx:latest
    container_name: web_server
    ports:
      - "8080:80"
  redis:
    image: redis:latest
```

Este archivo define dos servicios: `web` y `redis`. El servicio `web` usa la imagen `nginx:latest` y mapea el puerto `8080` en el host al puerto `80` en el contenedor. El servicio `redis` usa la imagen `redis:latest`.

## Cómo correr una aplicación con Docker Compose

Una vez que hayas definido tus servicios en el archivo `docker-compose.yml`, puedes ejecutar la aplicación con Docker Compose.

### Paso 1: Abrir la terminal

Asegúrate de estar en el directorio donde se encuentra tu archivo `docker-compose.yml`.

### Paso 2: Ejecutar Docker Compose

Para iniciar los servicios definidos en tu archivo `docker-compose.yml`, ejecuta el siguiente comando:

```bash
docker-compose up
```

Este comando descargará las imágenes necesarias (si no están ya presentes en tu máquina) y luego creará y arrancará los contenedores.

### Paso 3: Verificar que los servicios estén corriendo

Una vez que los servicios estén en ejecución, abre tu navegador y visita `http://localhost:8080`. Deberías ver la página de bienvenida de Nginx.

### Paso 4: Ejecutar en segundo plano

Si deseas ejecutar los servicios en segundo plano (detached mode), puedes usar la opción `-d`:

```bash
docker-compose up -d
```

Esto iniciará los servicios en segundo plano y liberará tu terminal.

## Cómo detener y eliminar contenedores

Para detener los servicios en ejecución, puedes usar el siguiente comando:

```bash
docker-compose down
```

Este comando detendrá y eliminará los contenedores, redes y volúmenes creados por `docker-compose up`.

## Explicación del archivo docker-compose.yml

El archivo `docker-compose.yml` es donde defines los servicios que componen tu aplicación. Aquí hay una explicación de las secciones más comunes:

- **version**: Especifica la versión del formato de archivo de Docker Compose.
- **services**: Define una lista de servicios que componen tu aplicación.
  - **image**: Especifica la imagen Docker que se usará para el servicio.
  - **ports**: Mapea puertos del host a puertos del contenedor.
  - **depends_on**: Define las dependencias entre servicios. Por ejemplo, si el servicio `api` depende de que el servicio `redis` esté en ejecución, se usa `depends_on` para especificar esta relación.
  - **volumes**: Define volúmenes para persistir datos. Los volúmenes permiten que los datos persistan incluso si el contenedor se detiene o se elimina. En el ejemplo, el volumen `vol` se usa para persistir los datos de Redis.

Ejemplo de un archivo `docker-compose.yml` con estas secciones:

```yaml
name: tutorial
services:
  redis:
    image: redis:7.4-alpine
    container_name: 'redis'
    ports:
      - 6379:6379
    environment:
      - REDIS_PORT=6379
    volumes:
      - vol:/data
  api:
    build: .
    container_name: 'api'
    ports:
      - '80:80'
    restart: on-failure
    environment:
      - REDIS_HOST=redis
      - REDIS_PORT=6379
    depends_on:
      - redis

volumes:
  vol:
    driver: local
```

## Conclusión

En esta guía, aprendiste cómo:

- Definir servicios en un archivo `docker-compose.yml`.
- Ejecutar una aplicación multi-contenedor con Docker Compose.
- Detener y eliminar contenedores usando Docker Compose.

Con este conocimiento, podrás orquestar aplicaciones complejas que constan de múltiples servicios, todo con un solo comando.
