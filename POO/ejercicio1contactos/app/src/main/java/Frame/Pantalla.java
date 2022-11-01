package Frame;

import ejercicio1contactos.Agenda;
import ejercicio1contactos.ContactoT1;
import ejercicio1contactos.ContactoT2;

import javax.swing.*;
import java.awt.*;

public class Pantalla extends JPanel {
    Agenda ag;
    GridBagConstraints c = new GridBagConstraints();
    JPanel showContactosPanel = new JPanel();

    boolean OffOnnContactos = true;

    public Pantalla(Agenda ag) {
        this.ag = ag;
        this.setLayout(new GridBagLayout());
        this.setSize(500, 500);
        JButton button = new JButton("Contactos");
        JButton button2 = new JButton("Eventos");

        button.addActionListener(e -> {
            // AÑADIR SU FUNCION AQUI JOSE, PUEDE TOAMR EN CUENTA LA FUNCION QUE SE LLAMA SHOWCONTACTOS;
            // TAMBIEN PONGA LAS VALIDACIONES QUE HICE EN EL BUTTON2, para que se borren datos anteriores
        });
        button2.addActionListener(e -> {
            if (OffOnnContactos){

                showContactos();
                this.OffOnnContactos = false;
            } else {
                showContactosPanel.removeAll();
                this.remove(showContactosPanel);
                this.OffOnnContactos = true;
                this.repaint();
            }

        });
        c.fill = GridBagConstraints.HORIZONTAL;
        c.weightx = 0.5;
        c.gridx = 1;
        c.gridy = 0;

        this.add(button, c);
        c.fill = GridBagConstraints.HORIZONTAL;
        c.weightx = 0.5;
        c.gridx = 2;
        c.gridy = 0;
        this.add(button2, c);
    }

    public void showContactos(){

        showContactosPanel.setLayout(new GridBagLayout());



        for (int i = 0; i < this.ag.getLista_contacto().size(); i++) {
            c.fill = GridBagConstraints.HORIZONTAL;
            c.ipady = 20;      //make this component tall
            c.weightx = 0.0;
            c.gridwidth = 3;
            c.gridx = 0;
            c.gridy = i;
            if(this.ag.getLista_contacto().get(i) instanceof ContactoT1){
                ContactoT1 T1 = (ContactoT1) this.ag.getLista_contacto().get(i);
                showContactosPanel.add(new JLabel(T1.empresa+" "+ T1.numero_empresa+ " "+T1.contacto_informacion.nombre+ " "+T1.contacto_informacion.edad), c);
            } else {
                ContactoT2 T2 = (ContactoT2) this.ag.getLista_contacto().get(i);
                showContactosPanel.add(new JLabel(T2.apodo+ " "+T2.numero+ " "+T2.tipo_familiar), c);
            }
        }
        showContactosPanel.repaint();
        c.fill = GridBagConstraints.HORIZONTAL;
        c.ipady = 40;      //make this component tall
        c.weightx = 0.0;
        c.gridwidth = 3;
        c.gridx = 0;
        c.gridy = 1;

        this.add(showContactosPanel, c);
        this.updateUI();


    }







}
