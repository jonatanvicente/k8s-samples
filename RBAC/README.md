
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


#### Uso con ConfigMaps

- Creamos cm-role.yml con namespace incluido
    - Damos permisos sobre configmaps en namespace dev
    - Ejecutamos con admin (**kubectl config use-context minikube**)
    - Verificamos creación de ConfigMap (**kubectl get cm -n dev**) y visualización con user jonatan en namespace dev (**kubectl get cm -n dev**). En default jonatan no tiene permisos
    -  **kubectl edit cm** no puede editar. Tendriamos que aplicar "patch" a los verbos. Aplicamos.
    -  Le damos permisos de borrado (verb "delete") con admin y, con user jonatan, borramos CM (**kubectl delete cm vars -n dev**) del namespace dev