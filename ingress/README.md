

### Instalación como minikube addon

- **minikube addons enable ingress**. Habilitamos addons con nginx default
- **kubectl get ns**. Verificamos existencia de namespace ''ingress-nginx''
- **kubectl get pods -n ingress-nginx**. Verificamos el pod de nginx en ejecución.

### Instalación de Ingress Controller como si fuese cluster completo K8s (no minikube)

- Es necesario ejecutar kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.12.2/deploy/static/provider/baremetal/deploy.yaml.
    - Descargamos el manifest  (**nginx-controller.yml**)
    - Ejecutamos apply
    - Verificamos existencia con **kubectl get pods -n ingress-nginx**
    - Este pod será el controlador. Ejecutamos **kubectl get svc -n ingress-nginx** y verificamos dos servicios nginx construidos.
