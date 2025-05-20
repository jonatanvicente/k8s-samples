

### Elastic Kubernetes Service (AWS)

Pasos a ejecutar:

#### 1.- Creación del master

Pasos a ejecutar:
- Crear cuenta y user IAM con permisos administrador
- Instalar aws cli
- Instalar eksctl
- Seleccionar zonas deseadas (Virginia, us-east-1, p.ej.)
- Crear el **master en AWS** sin crear ningún grupo de nodes (lo añadiremos más tarde)
    - Ejecutar **eksctl create cluster --name test-cluster --region us-east-1 --zones us-east-1a,us-east-1b --version 1.32 --without-nodegroup**
        - (Opcional) Version hace referencia a k8s version, ejecutar **kubectl version** para conocer.
        - Zonas de disponibilidad **zones**
- Entramos en Servicio CloudFormation de AWS y revisamos creación (tarda unos 15 min)
- Revisamos config creada en **"/home/[user]/.kube/config"**
    -  Si lo perdemos, ejecutaríamos **aws eks --region us-east-1 update-kubeconfig --name [clusterName]**
- En AWS EKS revisamos cluster
- Una vez creado cluster, **kubectl get svc** nos da info de servicios disponibles. Con **kubectl cluster-info** vemos info del cluster al que estamos conectados.

#### 2.- Creación de nodes

- **kubectl get nodes -o wide** indica que aún no existen nodos, no podemos crear pods aún. Seleccionaremos **Managed node group** (para personalizar)
- Creamos node:
    - **eksctl create nodegroup --cluster=test-cluster --name=test-workers --node-type=t3.medium --nodes=1 --nodes-min=1 --nodes-max=3 --asg-access --spot**
        - **name**: nombre del cluster
        - **node-type**: tipo de la instancia de máquina
        - **asg-access**: tipo de política de acceso
- Al crear el nodo, lo vincula a instancias EC2 de AWS. Es posible administrar también desde AWS Console.


####  Creación de pods

- **kubectl run nginx-pod --image=nginx:alpine --restart=Never**