
### Deployment

- <b>dep.yaml</b>
	- Creación de Deployment. <b>kubectl apply -f dep.yaml</b>
	- Deployment visible con <b>kubectl get deployment</b>
		- Labels visibles con <b>kubectl get deployment --show-labels</b>
- <b>dep2.yaml</b>
	- Actualización de Deployment sobre dep.yaml
- <b>IMPORTANTE:</b>
	- Podemos ejecutar la actualización con <b>kubectl apply -f dep2.yaml</b>
	- Vemos ejecución de cambios en tiempo real con <b>kubectl rollout status deployment [nombreDeploy]</b>
	- La actualización se efectúa en base a los parámetros fijados (o default, si no) <b>MaxUnavailable</b> y <b>MaxSearch</b>
- <b>dep3.yaml</b>
	- Nuevo cambio para ver histórico de cambios en Deployments. Podemos verificar con <b>kubectl rollout history deployment [nombreDeploy]</b>
	- Volvemos a colocar versión de imagen de dep.yaml.
	- Ejecutamos con <b>kubectl apply -f dep3.yaml</b>
- <b>dep4.yaml</b>
	- Inclusión de parámetro revisionHistoryLimit a 1 para modificación del máximo de registros de histórico (por default es 10).
	- Eliminamos Deployment con <b>kubectl apply -f dep4.yaml</b>


##### Change Cause (dep5.yaml)

- Elemento que aparece al mostrar el history en Deployments con <b>kubectl rollout history deployment [deploymentName]</b>
- Contiene metadata para añadir CHANGE-CAUSE (kubernetes.io/change-cause)


#### Rollback de un despliegue (dep6.yaml)

- File incorrecto (que apunta a una imagen que no existe) para forzar un rollback.
- Ejecutamos <b>kubectl apply -f dep6.yaml </b> (forzado error)
- <b>kubectl get pods<b> muestra el error
- <b>kubectl rollout history deployment deployment-test</b> nos indica el historial (para hacer el rollback)