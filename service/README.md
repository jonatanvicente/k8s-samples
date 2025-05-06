

### Services

- Ver file svc.yaml
- Podemos añadir en el file distintos tipos de resources que se crearán de manera secuencial
- Mapeo de puertos y secuencia de entrada (file svc.yaml):
	- IP del service: Kubernetes garantizará la IP del service (que no cambia)
	- El port del service --> spec > ports > port
	- Ports de los pods adonde redirige el Service --> spec > ports > targetPort
- File svc2.yaml: Inclusión de tipo de cluster en el Service. Delete service anterior y aplicación del nuevo:
	- kubectl delete -f svc.yaml
	- kubectl apply -f svc2.yaml
- File nodeport.yaml: Utilización de <b>nodeport</b> en lugar de <b>ClusterIP</b>.
	- OJO que cambian los nombres del deployment y del service
	- Hacemos apply del NodePort con svc.yaml también desplegado.
	- get pods muestra los 6 pods (filtramos con <b>kubectl get pods -l app=front</b> o <b>app=backend</b>)
	- <b>kubectl get svc</b> arroja ambos servicios (my-service y my-service1, además de kubernetes)
		- Bajo PORT(S) vemos el port que expone NodePort (p.ej. 8080:31078/TCP)