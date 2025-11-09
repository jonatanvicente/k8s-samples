

### Horizontal Pod Autoscaler

* Requiere siempre la instalación del server de métricas
* Para instalar en AWS, véase https://docs.aws.amazon.com/eks/latest/userguide/horizontal-pod-autoscaler.html
* Escala por CPU, y necesitaremos definir los límites


#### Ejecución de pruebas del HPA

- Ejecutamos deployment **kubectl apply -f https://k8s.io/examples/application/php-apache.yaml** (anexo también por si deseamos hacerlo localmente)
- Creamos HPA para el despliegue **kubectl autoscale deployment php-apache --cpu-percent=50 --min=1 --max=10**
    -   50% CPU, con min y max pods
    -   Cuando la carga promedio de la CPU es < 50%, HPA intenta reducir el número de pods al minimo (1)
    -   Si es superior, aumentará a un máximo de 10
- Vemos HPA **kubectl get hpa**, **kubectl get hpa php-apache**, **kubectl get svc**
- Creamos una carga para el server

```
            kubectl run -i \
                --tty load-generator \
                --rm --image=busybox \
                --restart=Never \
                -- /bin/sh -c "while sleep 0.01; do wget -q -O- http://php-apache; done"
```

- Observamos escalado **watch kubectl get pods** (Linux command, execute/display a program periodically) o bien **kubectl get hpa php-apache -w**
    - Al subir el consumo de CPU, HPA crea nuevos pods que puedan manejar las solicitudes, autoescalando nuestros pods cuando la CPU pasa de un determinado porcentaje.
    - Para escalar hacia abajo, HPA tarda más (intentando prever si hubiese otra subida)
- Borramos el Deployment: **kubectl delete deployment.apps/php-apache service/php-apache horizontalpodautoscaler.autoscaling/php-apache**


#### ¿Qué sucede si no hay más nodos disponibles para los pods que HPA está creando?

- No importa si colocamos un máximo de pods exagerado: si no hay nodos suficientes, k8s no podrá crear los últimos pods que HPA demande.
- La última réplica del pod no llegará a crearse (CPU insuficiente).
- Para que un nodo sea automáticamente creado (o eliminado en escalada hacia abajo), necesitaremos user el **Cluster Autoscaler**.

#### Parametrización

- Si estamos en AWS, es necesario revisar cuánto de CPU consumirá cada pod (según la instancia seleccionada, t3.medium, etc)
    - 1 cpu = 1000 millicores 