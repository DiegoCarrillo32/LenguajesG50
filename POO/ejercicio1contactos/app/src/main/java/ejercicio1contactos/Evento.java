package ejercicio1contactos;

import java.util.Date;

public class Evento {
    public Evento(String descripcion, String nombre, Date fecha) {
        this.descripcion = descripcion;
        this.nombre = nombre;
        this.fecha = fecha;
    }

    public String descripcion;
    public String nombre;
    public Date fecha;

}
