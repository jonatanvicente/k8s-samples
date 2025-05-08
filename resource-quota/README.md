

### ResourceQuota

* Creamos ResourceQuota junto con namespace
* Añadimos un Deployment
* Le fijamos límites
* Desplegamos (**kubectl apply -f res-quota.yml**)
* Revisamos correcta creación, pero hay errores (**kubectl describe deployments.apps -n uat deployment-test**)
* Revisamos el ReplicaSet (**kubectl get rs -n uat**). Ya que es el ReplicaSet quien crea los pods directamente, buscamos errores (**kubectl describe rs -n uat**)
    * **OJO: siempre que creamos un ResourceQuota, debemos fijar los Resources y los Limits** . Error porque no estaban fijados los limits. Añadimos.
* Verficamos que hemos llegado al límite (**kubectl describe ns uat**).
* Podemos revisar errores con **kubectl get deployments.apps -n uat deployment-test -o yaml**

##### Limite del total de objetos que deseamos crear

* Ejemplo en pod-quota.yml
* Verificamos creación máxima de objetos con **kubectl describe namespaces qa** (no se sobrepasa).
* Verificamos pods creados con **kubectl get pods -n qa*
* Si cambiamos el numero de replicas en pod-quota y hacemos apply de nuevo, los que exceden del numero de pods en las especificaciones del ResourceQuota no son creados.