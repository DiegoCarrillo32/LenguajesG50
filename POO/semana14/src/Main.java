import Operators.InfixError;

public class Main {
    public static void main(String[] args) throws InfixError {
        TranStringNum tsn = new TranStringNum();
        tsn.translateStringNum("580 + 30");
    }

}