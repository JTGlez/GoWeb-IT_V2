# GoWeb-IT_V2

Esta segunda iteración del proyecto web de **Go** implica un refactor completo siguiendo la filosofía de **Domain-Driven Design (DDD)**, que divide la arquitectura de la aplicación en tres dominios principales:

## Ejecución

Para ejecutar el proyecto, se utiliza un **Makefile** con diversos comandos para gestionar y manipular la aplicación. Para iniciar la aplicación, simplemente usa:

```bash
make run
```

## Dependencias
+ **Go** 1.23.1

### Estructura de Dominios

- **cmd**: contiene los puntos de entrada de la aplicación y los handlers del servidor web.
- **internal**: abarca los elementos internos de la aplicación que no deben exponerse, representando el core del sistema.
- **pkg**: incluye los elementos reutilizables de la aplicación que pueden usarse de forma independiente.

### Diseño en Capas

Para estructurar la arquitectura del servidor, se utiliza un **diseño en capas** que facilita el flujo de las peticiones según la entidad involucrada, siguiendo el siguiente esquema:

- **Controller**: capa de entrada que recibe las peticiones del cliente, valida los datos de entrada para que cumplan los criterios de aceptación y devuelve respuestas.
- **Service**: capa de negocio que procesa datos, genera nuevas estructuras y gestiona recursos y llamadas externas, como APIs o microservicios.
- **Repository**: capa de persistencia que abstrae el acceso a los datos, encargándose de su obtención y manipulación desde fuentes como archivos o bases de datos.

### Esquema de Comunicación entre Capas

La comunicación entre capas se implementa mediante **interfaces**, de modo que las llamadas no se realizan de forma directa, sino mediante un contrato que especifica cómo deben comunicarse las capas y qué métodos deben implementar. A continuación, se muestra un ejemplo gráfico de este esquema de comunicación entre capas:

![image](https://github.com/user-attachments/assets/6a48fe1a-980e-44dc-9ddc-e4357f9c5df2)

### Creación de Servicios mediante Factories

- Los servicios se generan a través de **factories**, que devuelven interfaces a partir de **structs** que implementan la firma de la interfaz.
    - En este esquema, el **router** captura la petición `GET` y llama al servicio `svcPong`, que es una instancia de **InterfacePong**. Este servicio invoca el método `GetPong`.
    - Cabe destacar que, en ningún momento, se llama directamente a un `struct`; en su lugar, se devuelve un `struct` envuelto en la interfaz al crear el servicio.

Finalmente, mediante la firma de la interfaz, se invoca la implementación específica de `GetPong` en el **struct pongService**, cumpliendo así el contrato de la **InterfacePong**.

