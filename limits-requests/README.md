

#### Uso de limits & requests

##### RAM

* limits-ram.yml muestra ejemplo de manejo de pod entre ambos límites
* Estresamos el pod usando limits-ram2.yml:
    *  Revisamos estado con **kubectl describe pod [podName]**. Revisamos logs
    *  Revisamos también **kubectl get pod [podName] -o yaml** para obtener más info
        *  Verificamos estado y causa
* Introducimos cantidad de memoria inalcanzable en limits-ram3.yml.
    *  Revisamos **kubectl describe pod [name]**
    *  Indica en los eventos warning: no hay nodos con la memoria que el pod solicita ('Insufficient memory').

##### CPUs

* limit-cpu.yml. Ejemplo de exceder valor de cpu's
* limit-cpu2.yml.
    * Vemos numero de cpu's en maquina con **kubectl describe node minikube**.