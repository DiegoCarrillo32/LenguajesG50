package Operators;

public class Number {
    public Integer number;
    public Number(String _number) {
        try {
            this.number = Integer.parseInt(_number);
        } catch (Exception e) {
            System.out.println(e);
        }

    }
}
