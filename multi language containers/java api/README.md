# Java API con Docker (Construcción en Dos Etapas)

Este proyecto utiliza Maven para construir una API en Java. La aplicación se empaqueta en un contenedor Docker utilizando una estrategia de construcción en dos etapas: una para la compilación y otra para la ejecución en tiempo de ejecución. A continuación se explican los pasos para construir y ejecutar la aplicación utilizando Docker.

## Instrucciones para correr el proyecto con Docker

### Paso 1: Construir la imagen de Docker

Primero, construye la imagen Docker de la aplicación Java ejecutando el siguiente comando:

```bash
docker build -t my-java-app .
```

Este comando creará una imagen de Docker a partir del `Dockerfile` en el directorio actual y la etiquetará como `my-java-app`.

### Paso 2: Ejecutar la aplicación en Docker

Una vez que la imagen se haya construido correctamente, puedes ejecutar el contenedor con el siguiente comando:

```bash
docker run --name java-container -p 8080:8080 my-java-app
```

Este comando hará lo siguiente:

- Ejecutará el contenedor con el nombre `java-container`.
- Mapeo de puertos: El puerto `8080` dentro del contenedor se expondrá en el puerto `8080` de tu máquina local.
- La aplicación Java será accesible a través de `http://localhost:8080` en tu navegador.

## Explicación del Dockerfile

Este `Dockerfile` utiliza una **estrategia de construcción en dos etapas**, lo que significa que se compila la aplicación en una imagen y luego se crea una imagen más pequeña y optimizada para la ejecución de la aplicación.

### Primera etapa: build de la aplicación

```Dockerfile
FROM maven:3.9-eclipse-temurin-24-alpine AS build
```

- **Descripción**: Usa la imagen base `maven:3.9-eclipse-temurin-24-alpine` para compilar la aplicación.  
  Esta imagen ya trae instalado **Maven** y **Eclipse Temurin 24 (JDK)** sobre **Alpine Linux**, lo que la hace ligera y adecuada para builds.  
  El alias `AS build` nombra esta etapa como `build` para poder referenciarla después en la segunda etapa.

```Dockerfile
WORKDIR /app
```

- **Descripción**: Establece el directorio de trabajo `/app` dentro del contenedor.  
  A partir de aquí, todos los comandos (`COPY`, `RUN`, etc.) se ejecutan relativos a este directorio.

```Dockerfile
COPY pom.xml mvnw ./
```

- **Descripción**: Copia el archivo `pom.xml` y el wrapper de Maven `mvnw` al directorio de trabajo (`/app`).  
  - `pom.xml` define las dependencias y configuración del proyecto.
  - `mvnw` permite ejecutar Maven usando el wrapper del proyecto, garantizando una versión consistente.

```Dockerfile
COPY .mvn .mvn
```

- **Descripción**: Copia la carpeta `.mvn`, que contiene archivos de configuración y scripts usados por el wrapper de Maven (`mvnw`).  
  Esto asegura que el wrapper funcione correctamente dentro del contenedor.

```Dockerfile
RUN mvn -B dependency:go-offline
```

- **Descripción**: Ejecuta Maven con el objetivo `dependency:go-offline`, que descarga todas las dependencias necesarias del proyecto.  
  - La opción `-B` activa el modo *batch* (sin interacción).  
  - Al descargar las dependencias en esta capa y basarla solo en `pom.xml`, Docker puede reutilizar esta capa en builds futuros mientras `pom.xml` no cambie, acelerando considerablemente las construcciones.

```Dockerfile
COPY src /app/src
```

- **Descripción**: Copia el código fuente de la aplicación (`src/`) al contenedor dentro de `/app/src`.  
  Esta instrucción se ejecuta después de haber cacheado las dependencias, de modo que los cambios en el código no obliguen a re-descargar todo desde cero.

```Dockerfile
RUN mvn -B clean package -DskipTests
```

- **Descripción**: Compila el proyecto y genera el archivo `.jar` dentro del directorio `target/`.  
  - `clean`: limpia compilaciones previas.
  - `package`: compila el código y empaqueta la aplicación (normalmente en un `.jar` ejecutable).
  - `-DskipTests`: omite la ejecución de tests durante el build para hacerlo más rápido (útil para entornos de ejemplo o desarrollo).

### Segunda etapa: imagen de runtime (ejecución)

```Dockerfile
FROM eclipse-temurin:24-jre-alpine
```

- **Descripción**: Define la imagen base de runtime usando `eclipse-temurin:24-jre-alpine`, que contiene solo el **JRE 24** sobre Alpine.  
  Es mucho más ligera que la imagen con Maven y JDK, reduciendo el tamaño de la imagen final.

```Dockerfile
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
```

- **Descripción**: Crea un grupo (`appgroup`) y un usuario (`appuser`) en el sistema dentro del contenedor.  
  - `-S` crea usuarios/grupos del sistema (sin home completo, etc.).  
  - `adduser -S appuser -G appgroup` agrega el usuario `appuser` al grupo `appgroup`.  
  Esto permite ejecutar la aplicación con un **usuario no root**, mejorando la seguridad.

```Dockerfile
WORKDIR /home/app
```

- **Descripción**: Establece el directorio de trabajo de la imagen de runtime en `/home/app`.  
  Aquí se ubicará el `.jar` que se va a ejecutar.

```Dockerfile
COPY --from=build /app/target/api-0.0.1-SNAPSHOT.jar app.jar
```

- **Descripción**: Copia el archivo `.jar` generado en la primera etapa (`build`) desde `/app/target/api-0.0.1-SNAPSHOT.jar` al directorio de trabajo de esta etapa, renombrándolo como `app.jar`.  
  - `--from=build` indica que el archivo se toma de la etapa llamada `build`.  
  - Este es el corazón del **multi-stage build**: todo lo necesario se construye en la primera etapa, pero solo el artefacto final (`app.jar`) se copia a la imagen ligera de runtime.

> 🔁 **Importante**: si cambias el `artifactId` o la versión en el `pom.xml`, el nombre del `.jar` cambiará y deberás actualizar también esta ruta en el Dockerfile.

```Dockerfile
EXPOSE 8080
```

- **Descripción**: Documenta que el contenedor usará el puerto `8080` para aceptar conexiones.  
  No abre el puerto por sí mismo, pero sirve como indicación para quien ejecute el contenedor y para herramientas como Docker o Kubernetes.

```Dockerfile
USER appuser
```

- **Descripción**: Indica que, a partir de este punto, todos los procesos dentro del contenedor se ejecutarán como el usuario `appuser` (no root).  
  Esta es una **buena práctica de seguridad**, evitando que la aplicación tenga privilegios de superusuario dentro del contenedor.

```Dockerfile
ENTRYPOINT ["java", "-XX:+UseContainerSupport", "-XX:MaxRAMPercentage=75.0", "-jar", "app.jar"]
```

- **Descripción**: Define el comando por defecto que se ejecutará cuando se inicie el contenedor.  
  - `java`: ejecuta la JVM.
  - `-XX:+UseContainerSupport`: hace que la JVM sea consciente de los límites de CPU y memoria del contenedor.
  - `-XX:MaxRAMPercentage=75.0`: limita la memoria máxima de la JVM al 75% de la memoria disponible para el contenedor.
  - `-jar app.jar`: indica que debe ejecutar el archivo `app.jar` copiado previamente.

## ¿Por qué usar construcción en dos etapas?

La construcción en dos etapas ofrece los siguientes beneficios:

- **Optimización de la imagen**: La imagen final será más pequeña porque no incluirá herramientas de construcción como Maven ni archivos temporales generados durante el proceso de compilación.
- **Mejor seguridad**: Minimiza la superficie de ataque, ya que el contenedor de producción solo contiene el archivo `.jar` y lo necesario para ejecutar la aplicación.
- **Desempeño mejorado**: La imagen de producción es más ligera, lo que mejora el tiempo de inicio y reduce el consumo de recursos.
