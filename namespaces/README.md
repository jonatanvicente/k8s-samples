

### Namespaces

* Creación de 2 namespaces (dev y pro) con 1 y 5 replicas respectivamente
```shell
kubectl apply -f ns.yml
kubectl get deploy -n dev  # Verificamos
kubectl get pods -n dev
```