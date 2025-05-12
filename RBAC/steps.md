## RBAC

## Steps

* Arranque de minikube indicando el modo de auth (**minikube start --vm-driver=none --extra-config=apiserver.authorization-mode=RBAC)**

## Creación de certificado & firma

* Create keys and sign (**openssl genrsa -out jonatan.key 2048**)
    * Aparecerá en el dir jonatan.key
* **Creamos CSR** (file) con la key generada previamente. Asignamos CommonName y Organization que k8s usará después (O ficticia)
    * Habitualmente el developer creará su CSR para solicitar acceso y enviarlo al admin de k8s
    * **openssl req -new -key jonatan.key -out jonatan.csr -subj "/CN=jonatan/O=dev"**
* Revisamos la CA location (**kubectl config view nos permite ver CA location**)
* **Firma del certificado con la CA** (revisando location de ca.crt y **ca.key que es la privada, ojo**). En este caso, jonatan.crt será el certificado de user firmado por la CA.
    * **sudo openssl x509 -req -in jonatan.csr -CA /root/.minikube/ca.crt -CAkey /root/.minikube/ca.key -CAcreateserial -out jonatan.crt -days 500**
* **Validamos** que escribimos bien el CN (**openssl x509 -in jonatan.crt -noout -text**)


## Firma del certificado en entorno aislado (Docker ephimeral)

Utilizaremos una maquina distinta para firmar el crt del nuevo user y para intentar acceder a cluster minikube remoto
* Usuario actual (habitual) en minikube es minikube (administrador, **kubectl config current-context**)
* Revisamos server de minikube (**kubectl config view | grep server**)
* Para firmarlo en una location completamente aislada (Docker container) -> **docker run --rm -ti -v $PWD:/test -w /test -v /root/.minikube/ca.crt -v /usr/bin/kubectl:/usr/bin/kubectl alpine sh**
* Creamos una folder compartida donde pasamos el crt **público** de minikube (**no la key**) y también la utilidad kubectl para firmar
* **Será necesario configurar el cluster dentro de Docker (kubectl dentro del container no apunta a ningún sitio)
    *  Revisamos donde está el cluster (**kubectl cluster-info**) y apuntamos a él (**kubectl config set-cluster minikube --server=https://192.168.1.140:8443 --certificate-authority=/ca.crt**)
    * Agregamos credenciales (**kubectl config set-credentials jonatan --client-certificate=jonatan.crt --client-key=jonatan.key**).
    * Seteamos contexto (**kubectl config set-context jonatan --cluster=minikube --user=jonatan**)
    * Cambiamos a context creado (**kubectl config use-context jonatan**)
* Ya tenemos un usuario **dentro del container** que está llamando al cluster minikube. Da forbidden ya que no se ha

### Secuencia de firma de crt de user en la misma maquina

* **kubectl config set-cluster minikube --server=https://192.168.49.2:8443 --certificate-authority=/home/Jonatan/.minikube/ca.crt**
* **kubectl config set-credentials jonatan --client-certificate=jonatan.crt --client-key=jonatan.key**
* **kubectl config set-context jonatan --cluster=minikube --user=jonatan**
* **kubectl config use-context jonatan**
* **kubectl config get-contexts** nos arroja el nuevo contexto con nuevo usuario


### Switch entre contexts / users

* **kubectl config get-contexts**
* **kubectl config use-context jonatan**. 



