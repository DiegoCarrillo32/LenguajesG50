package ejercicio1contactos;
public class ContactoT2 extends Contacto {
//    Lista para miembros de familia
    public String tipo_familiar;
    public String apodo;

    public ContactoT2(Persona _p, Integer _numero, String _tipo, String _apodo){
        super(_p, _numero);
        this.tipo_familiar = _tipo;
        this.apodo = _apodo;
    }


}