# Kubernetes

Este proyecto demuestra cómo usar Kubernetes para orquestar múltiples contenedores. Incluye ejemplos de cómo definir y ejecutar recursos en Kubernetes, como Deployments, Services, ConfigMaps y Secrets.

## Tabla de Contenidos

1. [¿Qué es Kubernetes?](#qué-es-kubernetes)
2. [Cómo aplicar configuraciones por primera vez](#cómo-aplicar-configuraciones-por-primera-vez)
3. [Cómo actualizar recursos existentes](#cómo-actualizar-recursos-existentes)
4. [Cómo crear un namespace](#cómo-crear-un-namespace)
5. [Cómo crear secrets](#cómo-crear-secrets)
6. [Explicación de los archivos YAML](#explicación-de-los-archivos-yaml)
7. [Consideraciones adicionales](#consideraciones-adicionales)

## ¿Qué es Kubernetes?

Kubernetes es una plataforma de orquestación de contenedores que automatiza el despliegue, la gestión y la escalabilidad de aplicaciones en contenedores. Con Kubernetes, puedes definir y gestionar aplicaciones complejas que constan de múltiples contenedores.

## Cómo aplicar configuraciones por primera vez

Si estás aplicando configuraciones por primera vez, generalmente usarás el comando `kubectl apply` para crear los recursos en el clúster.

```bash
kubectl apply -f <deployment | service | configmap>.yaml
```

## Cómo actualizar recursos existentes

Para actualizar recursos ya existentes, puedes seguir utilizando el comando `kubectl apply`. Este comando se encarga de hacer las actualizaciones necesarias en los recursos que ya están en el clúster según los cambios que hayas hecho en tus archivos de configuración.

```bash
kubectl apply -f <deployment | service | configmap>.yaml
```

Cuando actualizas un `ConfigMap` en Kubernetes, los cambios no se aplican automáticamente a los pods que ya están en ejecución. Esto es porque los pods leen el `ConfigMap` solo al momento de iniciar. Si deseas que los pods reflejen los cambios realizados en el `ConfigMap`, debes reiniciar los pods o hacer que se actualicen utilizando `Rolling Updates`.

Aquí te muestro cómo hacerlo:

- **Comando `kubectl rollout restart`:** Si tus pods están gestionados por un `Deployment`, puedes reiniciar los pods usando el comando `kubectl rollout restart`. Esto actualizará el `Deployment`, lo que a su vez reiniciará todos los pods asociados.

```bash
kubectl rollout restart deployment <nombre-del-deployment> -n <namespace>
```

- **Eliminar los Pods Manualmente:** Otra opción es eliminar los pods manualmente. Kubernetes recreará los pods automáticamente si están gestionados por un Deployment, ReplicaSet, o StatefulSet.

```bash
kubectl delete pod <nombre-del-pod> -n <namespace>
```

## Cómo crear un namespace

Para crear un namespace en Kubernetes, puedes usar el siguiente archivo YAML:

```yaml
kind: Namespace
apiVersion: v1
metadata:
  name: tutorial
  labels:
    name: tutorial
```

Aplica el archivo con el siguiente comando:

```bash
kubectl apply -f namespace.yml
```

## Cómo crear secrets

Para crear secrets en Kubernetes, primero necesitas convertir los valores a base64. Puedes usar el siguiente comando en tu terminal:

```bash
echo -n 'valor' | base64
```

Por ejemplo, para crear un secret con el usuario y la contraseña de Redis:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: db-user-secretkey
  namespace: tutorial
type: Opaque
data:
  REDIS_USER: "base64-encoded-user"

apiVersion: v1
kind: Secret
metadata:
  name: db-pass-secretkey
  namespace: tutorial
type: Opaque
data:
  REDIS_PASSWORD: "base64-encoded-password"
```

Aplica los archivos con el siguiente comando:

```bash
kubectl apply -f db.user.secretkey.yaml
kubectl apply -f db.pass.secretkey.yaml
```

## Explicación de los archivos YAML

El archivo `namespace.yml` define un namespace en Kubernetes. Aquí hay una explicación de las secciones más comunes:

- **kind**: Especifica el tipo de recurso (Namespace, Deployment, Service, etc.).
- **apiVersion**: Especifica la versión de la API de Kubernetes.
- **metadata**: Contiene información sobre el recurso, como el nombre y las etiquetas.
- **spec**: Define la especificación del recurso, como los contenedores en un Deployment o los puertos en un Service.

Ejemplo de un archivo `db.deployment.yml`:

```yaml
kind: Deployment
apiVersion: apps/v1
metadata:
  name: redis-deployment
  namespace: tutorial
  labels:
    app: redis
  annotations:
    config-update: "v1.0.0"
spec:
  replicas: 1
  selector:
    matchLabels:
      app: redis
  minReadySeconds: 10
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: redis
    spec:
      containers:
        - name: redis
          image: myUsername/myRedis:1.0.0
          imagePullPolicy: Always
          ports:
            - containerPort: 6379
          env:
            - name: REDIS_USER
              valueFrom:
                secretKeyRef:
                  name: db-user-secretkey
                  key: REDIS_USER
            - name: REDIS_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: db-pass-secretkey
                  key: REDIS_PASSWORD
          resources:
            limits:
              memory: "256Mi"
            requests:
              memory: "64Mi"
      imagePullSecrets:
        - name: regcred
```

## Consideraciones adicionales

- **Visualizar Cambios:** Si quieres ver qué cambios se harán antes de aplicar, puedes usar el comando `kubectl diff` (requiere que esté habilitada la funcionalidad de `kubectl diff` en tu clúster).

```bash
kubectl diff -f <archivo>.yaml
```

- **Revisar Estado:** Después de aplicar configuraciones, es útil verificar el estado de los recursos para asegurarte de que todo se ha aplicado correctamente. Puedes usar comandos como `kubectl get` y `kubectl describe`.

```bash
kubectl get deployments -n <namespace>
kubectl get services -n <namespace>
```

```bash
kubectl describe deployment <nombre-deployment> -n <namespace>
kubectl describe service <nombre-service> -n <namespace>
```

- **Eliminar Recursos:** Si necesitas eliminar recursos, puedes usar el comando `kubectl delete`. Esto es útil si estás reemplazando un archivo con una nueva configuración que requiere la eliminación de recursos existentes.

```bash
kubectl delete -f <archivo>.yaml
```

- **Rollbacks:** Kubernetes mantiene el historial de revisiones para los Deployments, por lo que puedes realizar rollbacks a versiones anteriores si es necesario:

```bash
kubectl rollout undo deployment/<nombre-deployment>
```

- **Rolling Update:** Para realizar un `Rolling Update`, solo necesitas actualizar la imagen del contenedor en tu archivo de configuración del `Deployment` y aplicar los cambios. Por ejemplo, si quieres actualizar la imagen de `my-app:1.0` a `my-app:2.0`, deberías cambiar la línea image en tu archivo YAML y después, aplica los cambios con `kubectl apply`. Puedes verificar el progreso de tu `Rolling Update` con el siguiente comando:

```bash
kubectl rollout status deployment/my-app
```

## Puertos para NodePort

En Kubernetes, cuando se configura un `Service` con el tipo `NodePort`, es crucial seleccionar un puerto que esté disponible y no esté en uso por otro servicio. Aunque Kubernetes no proporciona una herramienta integrada para listar los puertos libres en los nodos, puedes seguir estos pasos para gestionar los puertos de NodePort de manera efectiva:

### 1. Rango de Puertos

Kubernetes asigna puertos de `NodePort` dentro del rango de puertos de **30000** a **32767** por defecto. Esto significa que puedes elegir cualquier puerto dentro de este rango.

Sin embargo, hay algunas consideraciones:

- Evitar Colisiones: Es recomendable evitar puertos que ya estén en uso por otros servicios en el clúster o en el nodo.
- Documentación: Lleva un registro de los puertos que estás utilizando para evitar colisiones futuras.

### 2. Consultar Puertos Usados en el Nodo

Aunque Kubernetes no ofrece una forma directa de consultar puertos libres en los nodos, puedes inspeccionar los puertos utilizados en los nodos y evitar esos puertos para `NodePort`. Aquí hay algunos métodos:

- **Usar Comandos del Sistema Operativo:** Puedes usar comandos del sistema operativo en los nodos para listar puertos en uso. Dependiendo del sistema operativo, los comandos pueden variar:
  - **Linux:**

    ```bash
    sudo netstat -tuln | grep LISTEN
    ```

    o

    ```bash
    sudo ss -tuln
    ```

  - **Windows:**

    ```bash
    netstat -aon | findstr LISTEN
    ```

- **Revisar Servicios Kubernetes:** Verifica qué servicios `NodePort` ya están en uso en tu clúster Kubernetes. Puedes usar el siguiente comando para listar todos los servicios y sus puertos

```bash
kubectl get services --all-namespaces -o wide
```
