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

### Primera etapa: Construcción de la aplicación

```Dockerfile
FROM maven:3.9.9-eclipse-temurin-21-alpine AS build
```

- **Descripción**: Utiliza la imagen `maven:3.9.9-eclipse-temurin-21-alpine` como base para compilar la aplicación. Esta imagen tiene Maven y OpenJDK 21 preinstalados y está basada en Alpine Linux, lo que la hace ligera.

```Dockerfile
WORKDIR /app
```

- **Descripción**: Establece el directorio de trabajo dentro del contenedor a `/app`. A partir de este punto, todas las operaciones siguientes se realizarán dentro de este directorio.

```Dockerfile
COPY pom.xml /app/pom.xml
RUN mvn dependency:go-offline
```

- **Descripción**: Copia el archivo `pom.xml`, que define las dependencias de la aplicación, y luego ejecuta el comando `mvn dependency:go-offline` para descargar las dependencias de Maven, asegurando que estén disponibles para el siguiente paso. Esto también ayuda a evitar tener que descargar las dependencias cada vez que se construya la imagen si no ha habido cambios en `pom.xml`.

```Dockerfile
COPY src /app/src
```

- **Descripción**: Copia el código fuente de la aplicación al contenedor en el directorio `/app/src`.

```Dockerfile
RUN mvn clean package -DskipTests
```

- **Descripción**: Ejecuta el comando `mvn clean package`, que limpia el proyecto y genera el archivo `.jar` de la aplicación. La opción `-DskipTests` se utiliza para omitir las pruebas y acelerar el proceso de construcción.

### Segunda etapa: Creación de la imagen para ejecución en tiempo de ejecución

```Dockerfile
FROM eclipse-temurin:21.0.1_12-jre-alpine
```

- **Descripción**: Utiliza la imagen base `eclipse-temurin:21.0.1_12-jre-alpine`, que es una versión optimizada de OpenJDK 21 basada en Alpine Linux, ideal para ejecutar aplicaciones Java en producción.

```Dockerfile
COPY --from=build /app/target/api-0.0.1-SNAPSHOT.jar /home/api-0.0.1-SNAPSHOT.jar
```

- **Descripción**: Copia el archivo `.jar` generado en la primera etapa desde el contenedor `build` al contenedor de ejecución. El archivo `.jar` se encuentra en `/app/target/api-0.0.1-SNAPSHOT.jar` y se coloca en el directorio `/home/` del contenedor de ejecución.

```Dockerfile
EXPOSE 8080
```

- **Descripción**: Informa a Docker que el contenedor escuchará en el puerto `8080`. Esto es necesario para que la aplicación pueda ser accesible desde fuera del contenedor.

```Dockerfile
CMD ["java", "-jar", "/home/api-0.0.1-SNAPSHOT.jar"]
```

- **Descripción**: Especifica el comando que se ejecutará cuando el contenedor inicie. En este caso, ejecuta el archivo `.jar` con el comando `java -jar`.

## ¿Por qué usar construcción en dos etapas?

La construcción en dos etapas ofrece los siguientes beneficios:

- **Optimización de la imagen**: La imagen final será más pequeña porque no incluirá herramientas de construcción como Maven ni archivos temporales generados durante el proceso de compilación.
- **Mejor seguridad**: Minimiza la superficie de ataque, ya que el contenedor de producción solo contiene el archivo `.jar` y lo necesario para ejecutar la aplicación.
- **Desempeño mejorado**: La imagen de producción es más ligera, lo que mejora el tiempo de inicio y reduce el consumo de recursos.
