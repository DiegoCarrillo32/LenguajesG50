package ejercicio1contactos;

import java.util.Date;

public class EventoT1 extends  Evento {
    // Evento ejecutivo
    public String empresa;
    public String organizador;

    public EventoT1(String descripcion, String nombre, Date fecha, String empresa, String organizador) {
        super(descripcion, nombre, fecha);
        this.empresa = empresa;
        this.organizador = organizador;
    }
}
