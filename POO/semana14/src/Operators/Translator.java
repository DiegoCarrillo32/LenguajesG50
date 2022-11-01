package Operators;

import java.util.ArrayList;
import java.util.Arrays;

public class Translator {

    ArrayList<Number> numberList = new ArrayList<>();
    ArrayList<Operator> operatorList = new ArrayList<>();

    public Translator(String[] regexSplit) {
        ArrayList<String> operators = new ArrayList<String>(
                Arrays.asList("+",
                        "-",
                        "/", "*"));
        for (int i = 0; i < regexSplit.length; i++) {
            if(operators.contains(regexSplit[i])){
                operatorList.add(new Operator(regexSplit[i]));
            } else {
                numberList.add(new Number(regexSplit[i]));
            }
        }

    }
    public Integer calculate(){
        Integer amount;
        while(true){

            amount = 0;
            switch (operatorList.get(0).operator){

                case "+":
                    amount += numberList.get(0).number + numberList.get( 1).number;

                    break;
                case "-":
                    amount += numberList.get(0).number - numberList.get( 1).number;
                    break;
                case "*":
                    amount += numberList.get(0).number * numberList.get( 1).number;
                    break;
                case "/":
                    amount += numberList.get(0).number / numberList.get( 1).number;
                    break;
            }

            operatorList.remove(0);
            numberList.remove(0);
            numberList.remove(0);
            if(numberList.isEmpty() || operatorList.isEmpty()){
                return amount;
            }
            numberList.add(0, new Number(amount.toString()));
        }

    }
}
