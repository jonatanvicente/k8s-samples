

## Volumes - Aprovisionamiento estático

#### emptyDir

- Creación de empty-dir en empty-dir.yml (**kubectl apply -f empty-dir.yml**).
- Directorio estará disponible tanto tiempo como viva el pod
- Si el contenedor se recrea, el directorio todavía estará disponible (k8s detecta si el container murió y lo recrea)
- Creamos un emptyDir dentro del directorio de logs de nginx (/var/logs/nginx).
- Entramos en container (**kubectl exec -it test-pd -- sh**) y creamos un file vacío dentro del emptyDir (**touch /var/log/nginx/test.log**)
- Matamos el proceso nginx (**pkill nginx**). El container muere.
- k8s recrea el container
- Volvemos a entrar al pod y vemos el file creado que sigue estando (está ligado al Pod, que no murió)
- Este mecanismo suele usarse con caché.

#### PV/PVC

Los pasos son siempre los mismos:

1.- Creación de un PersistentVolume (PV). Lugar donde se almacena la data físicamente
    - Esto puede apuntar a una localización física, a una en cloud, etc.
    - (Para efectos de prácticas, crearemos un hostPath (directorio) en el cluster local en una folder, y ésa folder actuará como un PV)
2.- Creación del Claim (PersistentVolumeClaim, PVC). Quien reclama o apunta al PV.
3.- Creación del volume dentro del container, que apuntará al PVC



#### Creación sin selectors

-
- Después reclamaremos esto desde el PVC (el storage) y lo montaremos en un pod
- Aplicamos **pv-pvc.yml**
- Revisamos creación de volume:
    - **kubectl get pv --show-labels**
    - **kubectl describe pv task-pv-volume**
- Ahora reclamamos contenido del PV con el PVC
    - **kubectl get pvc --show-labels**. Revisamos estado, está BOUND al volume correspondiente


#### Creación con selectors

- Usamos pv-pvc-selectors.yml para forzar la utilización de un PV por un Claim
- Aplicamos pv-pvc-selectors.yml
- Al revisar los PVs, vemos un pv available pero otro bound (**kubectl get pv --show-labels**)

#### Vinculación de un PV/PVC a un pod

- Aplicamos **pod-pvc.yml**
- Aparte del PV/PVC añadimos deployment con pod que refleja el volume
- Al revisar el pod creado, podemos ver el Mount (**kubectl describe pod xxxx**)
- Entramos en el pod y por cmd accedemos a mysql
- Creamos database
- Salimos y eliminamos pod
- ReplicaSet recrea el pod
- Al entrar de nuevo en mysql, se puede ver la db creada en el pod anterior :)


## Volumes - Aprovisionamiento dinámico

k8s creará el PV si un PVC no encuentra uno disponible. Se hace con storageClass, empleamos **storage-class.yml**
- Creamos pvc y revisamos (**kubectl get pvc**). Aparece vinculado a un pv
- **kubectl get pv** aparece un storage pv standard de tipo hostpath (el que permite minikube) que ha creado k8s
    - Este storageClass podría ser un almacenamiento dinámico AWS, GoogleCloud, Azure, etc  (provisto por instalación k8s prod).
    - OJO que el ReclaimPolicy default es Delete, lo cual es peligroso (los PVs dinámicos serán eliminados si eliminamos el PVC)
    - Al eliminar el PVC (**kubectl delete -f storage-class.yml**) se elimina también el PV.
- **kubectl describe pv  [pvname]** nos permite ver la ruta donde minikube lo instaló (dentro del container minikube estará accesible)
- **Este recurso es dinámico. Al crear el Claim, automáticamente nos crea el PV**