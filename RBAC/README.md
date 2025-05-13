
## RBAC

- **Creamos el user** (seguir steps.md).
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


### Creación de usuario del Cluster (**ClusterRoles**) que puede visualizar todos los pods

- Uso de ClusterRoles. Apply de clusterrole-jonatan.yml con user admin
- Con user jonatan ahora podemos visualizar pods en cualquier namespace:
    - **kubectl get pods -n test**
    - **kubectl get pods -n dev**
    - **kubectl get pods -n kube-system****. Incluso puede visualizar pods internos de k8s (como admin user)

### Creación de usuario del Cluster (**ClusterRoles**) que puede visualizar todos los pods

Para crear un user administrador, tan sólo hay que crear un RoleBinding que enlace con roles admnistrador predefinidos en plataforma.
- Revisamos ClusterRoles disponibles (**kubectl get clusterroles**).
    - **minikube** está utilizando el clusterrole **cluster-admin**. Usaremos este rol predefinido para asignarlo a user jonatan.
    - Como el Role ya está creado, sólo será necesario añadirlo a un manifest, pero no crear un rol nuevo (véase **cluster-admin.yml**).
    - En RoleRef referenciamos el rol ya creado (**cluster-admin**) y ejecutamos como admin (**kubectl config use-context minikube**).
    - Hacemos switch a user jonatan (**kubectl config use-context jonatan**)
    - Ahora user jonatan puede acceder a todos los recursos del sistema (**kubectl get ns**, **kubectl get endpoints**, **kubectl get svc**, etc).



### Aplicación de roles sobre grupos

-  Generamos un nuevo usuario 'juan' (seguir pasos 1 y 3 de steps.txt), creado en grupo 'dev'. También creamos nuevo contexto
-  **kubectl config get-contexts** ahora arroja a minikube, jonatan y juan
    -  Con user admin, eliminamos los permisos admin de jonatan (**kubectls delete -f cluster-admin.yml**)
    -  Creamos role de desarrollo para jonatan y juan
    -  Aplicamos como admin **group-role.yml**
    -  Ahora tanto jonatan como juan pueden ver los services (**kubectl get svc**, **kubectl get svc -n kube-system**, etc)