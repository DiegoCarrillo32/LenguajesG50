package Operators;


public class Operator {

    public String operator;
    public Operator(String operator) {
        try {
            this.operator = operator;
        } catch (Exception e) {
            System.out.println(e);
        }

    }
}
