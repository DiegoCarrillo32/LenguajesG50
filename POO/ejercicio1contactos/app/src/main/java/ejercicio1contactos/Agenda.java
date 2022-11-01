package ejercicio1contactos;
import java.util.ArrayList;
public class Agenda {
  ArrayList<Contacto> lista_contacto = new ArrayList<>();
  ArrayList<Evento> lista_evento = new ArrayList<>();
  public Agenda(){

  };

  public ArrayList<Contacto> getLista_contacto() {
    return lista_contacto;
  };
  public void addContacto(Contacto _contacto){
      lista_contacto.add(_contacto);
  }

  public ArrayList<Evento> getLista_evento(){ return lista_evento; };

  public void addEvento(Evento _evento) { lista_evento.add(_evento); };
}