package ejercicio1contactos;
public class ContactoT1 extends Contacto {
//  CLASE DE PERSONAS EJECUTIVAS
    public String empresa;
    public String numero_empresa;

    public ContactoT1(Persona _p, Integer _numero, String _empresa, String _numero_empresa){
        super(_p, _numero);
        this.empresa = _empresa;
        this.numero_empresa = _numero_empresa;
    }

}