
### ReplicaSets
- **replicaset.yaml**
	- Crea 5 pods de 2 containers cada uno (con réplicas). OJO a:
		- Utilización de apiVersion y labels
		- Se crea con **kubectl apply -f [nombreFile].yaml**
		- Listamos con kubectl get replicaset o **kubectl get rs**
	- **Si eliminamos a mano un pod con kubectl delete pod [nombrnieCompletoPod] --> ReplicaSet lo vuelve a crear**
	- **Corregimos las propiedades de replicaset.yaml y fijamos a 2 el numero de replicas.**
	 	- Al ejecutar de nuevo **kubectl apply -f replicaset.yaml**, el replicaSet **elimina** los redundantes para dejar en funcionamiento los deseados (2).