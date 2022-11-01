package ejercicio1contactos;
import java.util.ArrayList;
public class Persona {
  public String nombre;
  public ArrayList<String> apellidos = new ArrayList<>();
  public Integer edad;

  public Persona(String _nombre, String _apellido,Integer _edad){
    this.nombre = _nombre;
    this.apellidos.add(_apellido);
    this.edad = _edad;
  }
}