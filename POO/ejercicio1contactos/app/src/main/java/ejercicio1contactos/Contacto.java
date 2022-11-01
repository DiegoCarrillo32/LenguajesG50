package ejercicio1contactos;
public abstract class Contacto {
    public Persona contacto_informacion;
    public Integer numero;

    public Contacto(Persona _p, Integer _numero){
        this.contacto_informacion = _p;
        this.numero = _numero;
    }
}