import re


class Translate:
    availableOperators = ["*", "/", "+", "-"]
    numberList = []
    operatorList = []

    def __init__(self):
        self.geek = "Test"

    def translate_regex(self, regex):

        for i in regex:
            if i in self.availableOperators:
                self.operatorList.append(i)
            else:
                self.numberList.append(int(i))

        while True:
            amount = 0
            match self.operatorList[0]:
                case "+":
                    amount += self.numberList[0] + self.numberList[1]

                case "-":
                    amount += self.numberList[0] - self.numberList[1]

                case "*":
                    amount += self.numberList[0] * self.numberList[1]

                case "/":
                    amount += self.numberList[0] / self.numberList[1]

            self.operatorList.pop(0)
            self.numberList.pop(0)
            self.numberList.pop(0)

            if len(self.numberList) == 0 or len(self.operatorList) == 0:
                return amount;
            self.numberList.insert(0, amount);


    def translate_string_num(self, number):
        result = re.match("(([0-9]+) ([-+/*]) )+\\w+", number)

        if result:
            print("Match found");
            res = re.split(" ", number)
            print(self.translate_regex(res));

        else:
            print("Match not found")
