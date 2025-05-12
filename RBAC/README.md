
## RBAC

- Creamos el user (seguir steps.md o bien steps.txt).
- Creamos el role:
    - **kubectl config use-context minikube**
    - **kubectl apply -f jonatan-pods.yml**
- Enlazamos el role con el user (RoleBinding) siendo admin
     - **kubectl apply -f jonatan-pods.yml** (con el elemento rolebinding)
- Con user jonatan intentamos ver pods, ya visibles (en namespace default)
    - No podremos ver services (**kubectl get svc**) ni endpoints, p.ej. (**kubectl get endpoints**)
- Añadimos más reglas a nuestro rol personalizado en jonatan-pods_2.yml
    - Para añadir deployments, buscamos en kubectl api-resources a que APIGROUP pertenecen ("apps/v1")
- Ejecutamos apply jonatan-pods_2.yml siendo admin (minikube)
    - Verificamos que podemos ver deployments una vez asignada la segunda regla. Si volvemos a ejecutar jonatan-pods.yml, eliminamos permisos y (de nuevo) los deployments no son visibles.
- El usuario **jonatan** ni siquiera puede crear pods.
    -  **kubectl config use-context jonatan**
    -  **kubectl apply -f ../pod/pod.yaml** no se puede ejecutar (p.ej.)
