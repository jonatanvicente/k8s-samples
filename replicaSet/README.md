
### ReplicaSets
- <b>replicaset.yaml</b>
	- Crea 5 pods de 2 containers cada uno (con réplicas). OJO a:
		- Utilización de apiVersion y labels
		- Se crea con <b>kubectl apply -f [nombreFile].yaml</b>
		- Listamos con kubectl get replicaset o <b>kubectl get rs</b>
	- <b>Si eliminamos a mano un pod con kubectl delete pod [nombrnieCompletoPod] --> ReplicaSet lo vuelve a crear</b>
	- <b>Corregimos las propiedades de replicaset.yaml y fijamos a 2 el numero de replicas.</b>
	 	- Al ejecutar de nuevo <b>kubectl apply -f replicaset.yaml</b>, el replicaSet <b>elimina</b> los redundantes para dejar en funcionamiento los deseados (2).