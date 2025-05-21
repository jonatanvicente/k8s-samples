## Pods

### Creación de recursos con descriptores yaml (folder pods)

- **pod.yaml**
	- Creación simple de un pod (u otro recurso) con **kubectl apply -f pod.yaml**
	- **Borramos** con **kubectl delete -f pod.yaml**
- **pod2.yaml**
	- Descriptor con dos pods en el mismo yml.
- **containers.yaml**
	- **Creación de un pod con dos contenedores**. Demo de qué ocurre cuando colisionamos puertos. Instalación de comando en Alpine.
		- **OJO** a: no podremos acceder desde un container al otro, pero **sí** comparten la red (no podrán compartir el mismo puerto).
		- Los containers se ven como localhost
	- Probamos **en Docker** imagen de python (que tiene server http de ejecución simple)
		- Lo hacemos con **docker run --net host -ti python:3.6-alpine sh**. Esto ejecuta lo siguiente:
			- Utiliza la imagen indicada (la baja si no está en registry local)
			- La ejecuta en red del host
			- Abre terminal interactivo (ti)...
			- ... con la ejecución del comando sh (shell)
	- Ejecutamos en shell del container python ejecutándose **python -m http.server 8081**
	- Comprobamos correcta ejecución en localhost:8081
	- Añadimos al yml los command para que cree un file con echo diferente
	- **Si ponemos el mismo puerto en ambos containers, da error. kubectl describe doscont** no arroja explicación del error
	- **kubectl logs doscont -c cont1** no da error.
	- **kubectl logs doscont -c cont2** >> Address in use (si pusimos el mismo puerto)
	- **Entramos en container con kubectl exec -ti doscont -c cont1 -- sh**
		- El SO Alpine no trae curl. Lo instalamos con **apk add -U curl** para hacer las pruebas
	- Si, para rehacer los containers (corrigiendo puerto), ejecutamos de nuevo **kubectl apply** da error: **un pod no puede actualizarse a sí mismo** (hacen falta objetos de más alto nivel).
	- Una vez la definición es correcta, **desde <i>dentro</i> de cada container hay acceso a localhost:8082 y 8083** con **curl localhost:[puerto]**. Comparten host, cada container se ve a sí mismo y al otro.
	- **Desde fuera, podemos hacer port-forward a los puertos del pod y ver la salida de cada container del pod.**
- **labels.yaml**. Metadata para distinguir pods.
	- Distinguimos con **kubectl get pods -l app=backend**, p.ej.
	- Se añade el parámetro -l con la label y el valor
	- Siempre deberemos añadir por lo menos el de app para distinguir qué tiene cada pod.