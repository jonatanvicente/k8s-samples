
## Creación de API token asociados a ServiceAccount

- Trabajaremos en Namespace por default.
- Creamos un nuevo ServiceAccount (**kubectl create serviceaccount serviceaccounttest**).
- Revisamos en **kubectl get sa** (hay dos) y **kubectl get serviceaccounts/serviceaccounttest -o yaml**


### Token Temporal (preferible)

- Creamos pod (**kubectl apply -f pod.yaml**) vinculado al ServiceAccount con **pod.yaml**. Usaremos el output (podemos usarlo para acceso a la API, contiene flag --duration si deseamos, o si no **será eliminado cuando el ServiceAccount sea borrado**)
- Creamos token enlazándolo al pod (**kubectl create token serviceaccounttest --bound-object-kind="Pod" --bound-object-name="podtest2"**)
- Copiamos el output y lo colocamos en token-review.yml
- Ejecutamos **kubectl create -o yaml -f token-review.yml**
- El JWT resultante podemos decodificarlo, etc

###  Token Permanente

- Tiene mayor riesgo que el temporal. Aplicamos **long-lived-API-token.yml**
- Si borramos el ServiceAccount que tiene asociado el secret, k8s automaticamente borra el token
- Revisamos token en **kubectl get secret/[tokenName] -o yaml**. Son almacenados como secrets
    - **kubectl get secret/long-lived-token -o yaml** vemos el resultado:
        -  ca.crt. Certificado público
        -  namespace codificado en base64
        -  token: jwt codificado en base64
- Al ejecutar **kubectl get secrets**, vemos  el service-account-token permanente creado
- Al ejecutar **kubectl describe sa** vemos los ServiceAccount del namespace por default. El namespace serviceaccounttest creado contiene el token enlazado **long-lived-token**

