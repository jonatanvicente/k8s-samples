


### Configmaps

#### Creación de ConfigMaps por cmd

- Creamos en configmaps-examples file de config de nginx
- Creamos configmap en cmd con **kubectl create configmap nginx-config --from-file=configmaps-examples/nginx.conf**
    - **kubectl describe cm nginx-config** nos indica la key (**el nombre del file por default**) y todos los campos asociados
- Podemos crear un ConfigMap con varios files pasando la ruta de la folder(**kubectl create configmap nginx-config1 --from-file=configmaps-examples**)
    -  Verificamos el contenido con dos llaves (index.html y nginx.conf) con **kubectl describe cm nginx-config1**


#### Creación de ConfigMaps con manifests

- Ejemplo con **volumes** en cm-nginx-vol.yaml
    - Creamos deployment para consumirlo
- Ejemplo con **variables de entorno** y mixto en cm-nginx-env.yml
    -   Creamos otro configMap con script y env vars: las env vars serán consumidas del ConfigMap descrito
        -   Creamos script haciendo echo a folder root de nginx volcando a file. Esto será un script
    -   Indicamos el script en volume y montamos las env var haciendo ref a ConfigMap indicado
    -   Creamos pod, entramos y validamos env vars con **env**
    -   Revisamos script creado en /opt, lo invocamos
    -   **cat /usr/share/nginx/html/test.html** captura correctamente las env vars.