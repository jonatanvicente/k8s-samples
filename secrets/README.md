

#### Creación de un secret desde un file

- Creación de file test.txt con secret
- Secrets se codifican en Base64
- Comandos empleados:
    - **kubectl create secret generic mysecret --from-file=./secret-files/test.txt**
    - **kubectl get secrets -o yaml**
    - **kubectl describe secrets mysecret**