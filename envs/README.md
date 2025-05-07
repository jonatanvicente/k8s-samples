

#### Environment variables


- **Asignar env vars** en env.yml.
    - Validamos entrando en el container y ejecutando **env** y echo $[varName] ($VAR1)
    - Vemos data que k8s maneja en **kubectl get pods envar-demo -o yaml**
- **Exponer información del pod a los containers** en ref.yml
    - **kubectl get pods envar-demo -o yaml** nos muestra muchos campos de config del pod
    - Podemos rescatar esta info y hacerla visible al container con atributo **fieldRef**
    - Entramos a container (**kubectl exec -it dapi-envars-fieldref -- /bin/sh**)
    - Revisamos env (**env**), valores MY_NODE_NAME ("minikube"), MY_POD_NAME("dapi-envars-fieldref"), MY_POD_NAMESPACE y MY_POD_IP