

#### Creación de un secret desde un file txt

- Creación de file test.txt con secret
- Secrets se codifican en Base64
- Comandos empleados:
    - **kubectl create secret generic mysecret --from-file=./secret-files/test.txt**
    - **kubectl get secrets -o yaml**
    - **kubectl describe secrets mysecret**


#### Creación de un secret desde manifest

- Creación fijando data sensible como base64 en secret-data.yml
- StringData codifica directamente la data sensible (secret-stringdata.yml).

#### Ocultación de credenciales / sustitución de valores

- Creación de file con secrets ocultos (secret-secure.yml). k8s no hace la sustitución correctamente de valores cuando damos valores a env vars
- Ocultación y fijación de valores
    -   Creamos file con referencia a env vars (secret-secure_2.yml)
    -   Utilizamos tool envsubst para sustituir el valor del file por las env vars previamente creadas
    -   Volcamos a un nuevo file NO guardado en CVS (**envsubst < secret-secure_2.yml > tmp.yml**)
    -   Aplicamos el nuevo file (**kubectl apply -f tmp.yml**)
