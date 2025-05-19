
## Ingress Controller

- Creamos Ingress Controller a mano (**no como minikube addon**)

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

## Creación de service para enrutamiento a los pods

- Apply de ingress-example.yml
    - Creamos pods en **namespace default**. Modificamos el cmd del pod añadiendo instrucción específica de arranque
    - **Creamos Service ClusterIP** para acceso **interno** al Deployment.
        - **Probamos el service:** Creamos pod temporal, entramos en él (**kubectl exec -it podName -- sh**) y ejecutamos curl a **app/v1:8080**. Vemos salida configurada nginx en ingress-example.yml. OJO que aplicación todavía no está expuesta en el cluster, tan sólo llegamos a página de inicio de nginx. Necesitamos crear las reglas en Ingress

## Ingress: creación de reglas en punto de entrada 

**NOTA:**
- Usar **kubectl describe pods/pod-xxxxxx** para ver errores (con **-w para watch**). **kubectl get pods/pod-xxxx -o yaml** también.
