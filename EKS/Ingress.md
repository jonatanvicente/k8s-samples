####  4.- Adición de Ingress

![Diseño](EKS_schema.png)

- Ya tenemos el cluster en EKS funcionando (máster + nodes)
- Proceso:
    -  Crearemos el IngressController como Deployment
    -  Crearemos las IngressRules
    -  El IngressController creará automáticamente (dentro de la infraestructura AWS) un Load Balancer accesible desde el exterior
    -