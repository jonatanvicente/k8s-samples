

## Volumes

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