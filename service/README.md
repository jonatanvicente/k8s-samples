
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
- File nodeport.yaml: Utilización de **nodeport** en lugar de <b>ClusterIP**.
	- OJO que cambian los nombres del deployment y del service
	- Hacemos apply del NodePort con svc.yaml también desplegado.
	- get pods muestra los 6 pods (filtramos con <b>kubectl get pods -l app=front** o <b>app=backend**)
	- <b>kubectl get svc** arroja ambos servicios (my-service y my-service1, además de kubernetes)
		- Bajo PORT(S) vemos el port que expone NodePort (p.ej. 8080:31078/TCP)

Para accesos desde fuera del cluster k8s a los servicios expuestos, véanse las siguientes secciones **Port Forwarding** y **minikube tunnel**

#### Port Forwarding

- Para exponer cualquier puerto del cluster (con propósitos de **desarrollo**), es necesario hacer port-forwarding. Para acceder a la salida del servicio:
    -  Desde localhost **kubectl port-forward service/[myService] -n [namespace] [external_port]:[pod_port]**
    - Para acceso desde cualquier IP **kubectl port-forward service/[myService] -n [namespace] [external_port]:[pod_port] --address 0.0.0.0**


#### minikube tunnel

- Para ejecutar un tunnel y hacer accesible cualquier servicio podemos hacer un tunnel con
    - **minikube service ingress-nginx-controller -n ingress-nginx  --url [url]**