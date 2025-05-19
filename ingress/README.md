
## Ingress Controller

Pasos a seguirr:

- Creamos pods
- Creamos ClusterIP para acceso a los pods y verificamos acceso desde un pod temporal
- Creamos Ingress Controller a mano (**no como minikube addon**)
- Creamos Ingress Rules manejadas por Ingress Controller

### Instalación como minikube addon

- **minikube addons enable ingress**. Habilitamos addons con nginx default
- **kubectl get ns**. Verificamos existencia de namespace ''ingress-nginx''
- **kubectl get pods -n ingress-nginx**. Verificamos el pod de nginx en ejecución.

### Instalación de Ingress Controller como si fuese cluster completo K8s (no minikube)

- IngressController **se comporta como un Servicio expuesto más**.
- En producción, una de las opciones es poner un LoadBalancer (external provider) delante, que enrute a su NodePort como puerta de entrada al cluster k8s.
- Es necesario ejecutar kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.12.2/deploy/static/provider/baremetal/deploy.yaml.
    - En su lugar, descargamos el manifest  (**nginx-controller.yml**) y ejecutamos apply
    - Verificamos existencia de pods con **kubectl get pods -n ingress-nginx** (2 completed y 1 running)
    - Verificamos existencia de services con **kubectl get svc  -n ingress-nginx**. Tenemos un ClusterIP y un NodePort
    - Este pod será el controlador. Ejecutamos **kubectl get svc -n ingress-nginx** y verificamos dos servicios nginx construidos.
    - Hacemos **ip a** para conocer ip real del equipo (usando wi-fi).
    - **Hacemos minikube tunnel para mapear los puertos de los services** al exterior del cluster (o, si no, port forward).
    - Al acceder vemos pagina de nginx (del Ingress Controller ejecutándose, ninguna página de momento).

### Mapeo de puertos para ver punto de entrada

#### Port Forwarding

- Para exponer cualquier puerto del cluster (con propósitos de **desarrollo**), es necesario hacer port-forwarding. Para acceder a la salida del servicio:
    - Desde localhost **kubectl port-forward service/[myService] -n [namespace] [external_port]:[pod_port]**
    - Para acceso desde cualquier IP **kubectl port-forward service/[myService] -n [namespace] [external_port]:[pod_port] --address 0.0.0.0**


#### minikube tunnel

- Para ejecutar un tunnel y hacer accesible cualquier servicio podemos hacer un tunnel con
    - **minikube service ingress-nginx-controller -n ingress-nginx  --url [url]**. (ponemos '/' como url)
        - Al ejecutar la instrucción anterior, el tunnel abre dos ip **externas al cluster k8s** (una mapeando el puerto 80 y otra para el 443). La web de nginx es accesible desde el navegador. La url al puerto 80 será el puerto de entrada.
            - http://192.168.49.2:30784       # port 80
            - http://192.168.49.2:30679       # 443


## Creación de service para enrutamiento a los pods

- Apply de ingress-example.yml
    - Creamos pods en **namespace default**. Modificamos el cmd del pod añadiendo instrucción específica de arranque
    - **Creamos Service ClusterIP** para acceso **interno** al Deployment.
        - **Probamos el service accesible:** Creamos pod temporal, entramos en él (**kubectl exec -it podName -- sh**) y ejecutamos **curl app-v1:8080**.
        - El servicio está accesible: vemos salida configurada nginx en ingress-example.yml.
        - OJO que aplicación todavía no está expuesta desde fuera del cluster, tan sólo llegamos a página de inicio de nginx (que hemos modificado). Necesitamos crear las reglas en Ingress

## Ingress: creación de reglas en punto de entrada 

Véase ingress-rules.yml
    - **Manejo de atributo ingressClassName:** The ingressClassName attribute in an Ingress resource tells Kubernetes which Ingress Controller should handle this Ingress rule.
        -**It must match the ingressClass resource that was registered by your Ingress Controller**.
        - Para saber cuál es el nombre **kubectl get ingressclass** nos indica el controller creado (en el ejemplo, a partir de nginx-controller.yml) y **su nombre**.
            - Este nombre es el que debemos asignar a ingressClassName.
- Ahora, al invocar http://192.168.49.2:30784/appv1 vemos la salida correcta (IP asignada para red fuera del cluster:port asignado / path asignado en rules)

#### Manejo de Errores

- **Errores en despliegue de pods**: Usar **kubectl describe pods/pod-xxxxxx** para ver errores (con **-w para watch**). **kubectl get pods/pod-xxxx -o yaml** también.
- **Errores en enrutamientos desde Controller/Rules a Service**: Ver pods en namespace ingress-nginx (el controller es uno de ellos)
    - Ver logs de arranque del Controller con **kubectl -n ingress-nginx logs -f ingress-nginx-controller-xxxxxxx**
        - Cuando los enrutamientos son correctos, la traza arroja los bindings de ruta y también peticiones que van llegando.

