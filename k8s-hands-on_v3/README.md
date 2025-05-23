## Creación de frontend

#### Pruebas con pod efímero

* Creamos pod temporal (**kubectl run --rm -ti  podtest3 --image=nginx:alpine -- sh**) y le añadimos curl (**apk add curl**) para simular el front.
* Pod accede correctamente a IP del ClusterIP de backend
* Copiamos el index de front (adjunto) en lugar de pg de inicio de nginx **del pod** en /usr/share/nginx/html/index.html (vi viene por defecto con Alpine)
* Arrancamos nginx (comando **nginx** en Alpine) y verificamos con **ps aux** que está corriendo.
* Revisamos nuestro pod creado con **kubectl get pods -o wide** (aparte del cluster)
    - Habilitamos visualización de containers k8s para la session (**eval $(minikube -p minikube docker-env)**)
* Acceder a la IP del pod temporal desde el navegador no es permitido (necesitaríamos hacer port-forwarding o crear un service para acceso).

#### Creación del manifest

* Creamos frontend.yml (similar a backend.yml)
* Le cambiamos el tipo de service a NodePort
* Construimos image (OJO que tag debe coincidir con el que pusimos en el manifest frontend.yml) **docker build -t frontend-k8s-hands-on:v1 -f Dockerfile .**
* Desplegamos **kubectl apply -f frontend.yml**
* Verificamos **kubectl get svc**. Frontend service tiene NodePort
* Error CORS:
    - Intentamos allow-origin en backend, o accediendo desde index por ip/dns.
    - Pods de front resuelven bien y obtienen respuesta accediendo por ip/dns (entrando en cada pod)
    - Sin embargo, la pagina html no resuelve backend con la ip interna o el dns (aunque muestra correctamente código estático).
    - Deberemos establecer un NodePort en lugar de un ClusterIP para backend, o bien un proxy en front.