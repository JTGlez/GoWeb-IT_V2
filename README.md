# GoWeb-IT_V2

Esta segunda iteración del proyecto web de Go involucra un refactor completo del proyecto siguiendo la filosofía *Domain-Driven Design* que divide la arquitectura de la aplicación en 3 dominios principales:
+ **cmd**: dominio que abarca los entrypoints de la aplicación y los handlers del servidor web. 
+ **internal**: dominio que incluye todos aquellos elementos internos a la aplicación que no se desean exponer al exterior y que representan el core de la misma.
+ **pkg**: dominio que incluye aquellos elementos reusables de la aplicación y que pueden ser usados de forma aislada a la aplicación.

A su vez, para estructurar la arquitectura del servidor se emplea un diseño en capas, donde el flujo de las peticiones según la entidad involucrada emplea el siguiente esquema:

+ **Controller**: capa de entrada que recibe peticiones del cliente, valida que los datos de entrada cumplan criterios iniciales de aceptación y devuelve respuestas.
+ **Service**: capa de negocio que procesa datos, genera nuevas estructuras y maneja recursos y/o llamadas externas, como APIs o microservicios dentro de la red interna.
+ **Repository**: capa de persistencia que abstrae el acceso a los datos y se encarga de obtenerlos y/o manipularlos de una fuente de datos (archivos, bases de datos, etc).

El esquema de comunicaciones entre capas se implementa por medio de interfaces, de modo que las llamadas no se realizan de forma directa, si no por medio de un contrato que controla cómo deben comunicarse cada una de las capas y qué metodos deben implementar para ello. Esto se ve reflejado con el siguiente ejemplo, que muestra gráficamente cómo se realiza la comunicación entre capas:

![image](https://github.com/user-attachments/assets/6a48fe1a-980e-44dc-9ddc-e4357f9c5df2)

+ Los servicios se crean mediante factories, que devuelven interfaces a partir de structs que implementen la firma de la interfaz a devolver. 
+ + En este esquema, el router atrapa la petición GET y llama al servicio svcPong, que es una InterfacePong, e invoca el método GetPong. Se resalta que en ningún momento se está llamando directamente a un struct, si no que al crear el servicio se está devolviendo un struct wrappeado sobre la interfaz.
+ Finalmente, por medio de la llamada a la firma de la interfaz, se invoca a la implementación propia de GetPong por parte del struct pongService que, a su vez, es una InterfacePong.
  
