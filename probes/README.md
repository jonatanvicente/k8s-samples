
#### Probes

#### Liveness

En el ejemplo liveness-tcp.yml, ejecutamos los siguientes pasos:
* Creación de un container que crea un file, duerme 30 segundos, elimina el file y duerme 600 segundos
* **Secuencia del liveness:**
    * Cada 5 segundos hace un cat del file. Cuando llegue al intervalo de 35 segundos, dará un fail
    * k8s reiniciará el contenedor dentro del pod. Pasados 35 segundos, volverá a dar fail
    * La secuencia se repetirá varias veces hasta que caiga en **CrashLoopBackOff**, que significa que k8s intentó varias veces pero un bug hace crush del container, y dejará de reiniciarlo
* **Secuencia de instrucciones:**
    *  **kubectl get pods**
    *  **kubectl get pods -w**. Para observar eventos que se producen en los pods
    *  **kubectl describe pod liveness-exec**. Examinar log completo de eventos


#### Readiness

Véase ejemplo en liveness-tcp.yml
* **Readiness** desregistra los puertos del container del servicio si hay error, pero no reinicia el container.