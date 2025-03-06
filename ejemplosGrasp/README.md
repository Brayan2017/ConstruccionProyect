# patterns_examples

### Ejemplos de aplicación de grasp

Se implementaron los patrones de controller y de polimorfismo

- Ubicación: folder ejemplosGrasp
- Se encuentra desarrollado en Golang.
- Se implementa una api sencilla con dos endpoints para explicarlos.
- Para representar el patrón controlador se tiene una carpeta controller que se encarga de recibir y validar las solicitudes para direccinar al servicio luego de algunas validaciones que se realizan.
- El patrón de polimorfismo se encuentra representado en el endpoint '/polymorphism/vehicles/{model}/starts', con el objetivo de simular el encendido de un auto de diferentes modelos.

**Endpoints**

- /controller/vehicles/{vehicle_id}/find: Retorna infoirmacion de un vehiculo como id, marca, modelo, color.
- /polymorphism/vehicles/{model}/starts: retorna si el modelo del vehiculo se enciende con llave o botón.
