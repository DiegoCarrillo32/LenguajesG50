import Operators.InfixError;
import Operators.Translator;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class TranStringNum {

    public void translateStringNum(String number) throws InfixError {

        try {
            Pattern pattern = Pattern.compile("(([0-9]+) ([-+/*]) )+\\w+", Pattern.CASE_INSENSITIVE);

            Matcher matcher = pattern.matcher(number);
            boolean matchFound = matcher.find();
            if(matchFound) {
                Pattern patternSpace = Pattern.compile(" ");

                String[] patron = patternSpace.split(number);
                Translator trans = new Translator(patron);
                System.out.println(trans.calculate());
                System.out.println("Match found");
            } else {
                System.out.println("Match not found");
            }
        } catch (Exception e ){
            throw new InfixError("Expresion regular invalida "+number);
        }

    }

}
