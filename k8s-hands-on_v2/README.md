## Creación de backend

#### Creación de Service

* Relocated de files. Generamos descriptor backend.yml
* Regeneramos imagen con **docker build -t k8s-hands-on -f Dockerfile .**
* Deploy **kubectl apply -f backend.yml**
* Verificamos deploy con **kubectl get deployment.apps**. Error:
    - 0 available
    - Verificamos pods con **kubectl get pods**. Hay error ErrImagePull (no hay imagen en repo, está en local)
    - Cambiamos imagePullPolicy para que busque 1º en repo local. Regeneramos imagen y apply de nuevo.
* Al invocar en navegador, visualizamos el pod que resuelve el servicio (hostname)
* IP ClusterIP será punto de acceso interno, no accesible desde el exterior

#### Prueba del Service (IP + DNS)

* Creamos un pod de pruebas para ver si podemos acceder al ClusterIP (**kubectl run --rm -ti podtest3 --image=nginx:alpine -- sh**)
* Instalamos curl en el container (**apk add -U curl**).
* Ejecutamos curl [clusterIP] (disponible en Services en Lens). Es accesible (sale por 80).
    - Sucesivas invocaciones indican hostnames diferentes (un pod diferente atiende la petición)
* Ejecutamos **kubectl get svc** para obtener dns del cluster
    - Ejecutamos curl [name]:80. Accesible también.

