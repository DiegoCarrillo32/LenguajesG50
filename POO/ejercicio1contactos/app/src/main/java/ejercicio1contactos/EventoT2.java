package ejercicio1contactos;

import java.util.Date;

public class EventoT2 extends Evento {
    // Evento familiar
    public String familia;

    public EventoT2(String descripcion, String nombre, Date fecha, String familia) {
        super(descripcion, nombre, fecha);
        this.familia = familia;
    }
}
