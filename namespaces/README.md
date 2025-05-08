

## Namespaces

* Creación de 2 namespaces (dev y pro) con 1 y 5 replicas respectivamente en ns.yml
```shell
kubectl apply -f ns.yml
kubectl get deploy -n dev  # Verificamos
kubectl get pods -n dev
```

#### Uso de DNSs entre namespaces

* Creamos pod efímero (**kubectl run --rm -ti podtest3 --image=nginx:alpine -- sh**)
* curl a servicios en otro namespace no funciona, salvo que pongamos el descriptor completo [ServiceName].[NamespaceName].svc.cluster.local
```shell
curl backend-k8s-hands-on   # error
# respuesta correcta del backend
curl backend-k8s-hands-on.ci.svc.cluster.local    # [ServiceName].[NamespaceName].svc.cluster.local
```

#### Contexts

* Revisamos opciones actuales de config con **kubectl config view**
* Creamos nuevo context asociándolo a un namespace determinado.
```shell
kubectl config set-context ci-context --namespace ci --cluster minikube --user=minikube
```
* Cambiamos de contexto (**kubectl config use-context ci-context**).
    -  Al estar apuntando a otro namespace, vemos sus recursos por default (no es necesario fijar el namespace al hacer get pods, get services, get deployments, etc).