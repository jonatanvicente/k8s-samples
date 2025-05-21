
### Deployment

- **dep.yaml**
	- Creación de Deployment. **kubectl apply -f dep.yaml**
	- Deployment visible con **kubectl get deployment**
		- Labels visibles con **kubectl get deployment --show-labels**
- **dep2.yaml**
	- Actualización de Deployment sobre dep.yaml
- **IMPORTANTE:**
	- Podemos ejecutar la actualización con **kubectl apply -f dep2.yaml**
	- Vemos ejecución de cambios en tiempo real con **kubectl rollout status deployment [nombreDeploy]**
	- La actualización se efectúa en base a los parámetros fijados (o default, si no) **MaxUnavailable** y **MaxSearch**
- **dep3.yaml**
	- Nuevo cambio para ver histórico de cambios en Deployments. Podemos verificar con **kubectl rollout history deployment [nombreDeploy]**
	- Volvemos a colocar versión de imagen de dep.yaml.
	- Ejecutamos con **kubectl apply -f dep3.yaml**
- **dep4.yaml**
	- Inclusión de parámetro revisionHistoryLimit a 1 para modificación del máximo de registros de histórico (por default es 10).
	- Eliminamos Deployment con **kubectl apply -f dep4.yaml**


##### Change Cause (dep5.yaml)

- Elemento que aparece al mostrar el history en Deployments con **kubectl rollout history deployment [deploymentName]**
- Contiene metadata para añadir CHANGE-CAUSE (kubernetes.io/change-cause)


#### Rollback de un despliegue (dep6.yaml)

- File incorrecto (que apunta a una imagen que no existe) para forzar un rollback.
- Ejecutamos **kubectl apply -f dep6.yaml ** (forzado error)
- **kubectl get pods** muestra el error
- **kubectl rollout history deployment deployment-test** nos indica el historial (para hacer el rollback)