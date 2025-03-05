public class JugadorAmerica implements Jugador {

    private String nombre;
    private String numero;
    private String posicion;

    public JugadorAmerica(String nombre, String numero, String posicion) {
        this.nombre = nombre;
        this.numero = numero;
        this.posicion = posicion;
    }

    public String getNombre() {
        return nombre;
    }

    public void setNombre(String nombre) {
        this.nombre = nombre;
    }

    public String getNumero() {
        return numero;
    }

    public void setNumero(String numero) {
        this.numero = numero;
    }

    public String getPosicion() {
        return posicion;
    }

    public void setPosicion(String posicion) {
        this.posicion = posicion;
    }

    @Override
    public void jugarPartido() {

        System.out.println("El jugador "+nombre +" esta jugando el partido con el numero "+ numero);

    }
}
