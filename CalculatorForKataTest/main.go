package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func execute(a, b int, operator string) (rez int, err error) {

	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("неправильный оператор: %s", operator)
	}
}

func romanToInt(str string) (num int, err error) {

	var roman = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}

	for k, v := range roman {
		if str == k {
			return v, nil
		}
	}
	return 0, fmt.Errorf("неверный ввод римских цифр: %s", str)
}

func intToRoman(num int) (roman string) {

	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	sym := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			roman += sym[i]
		}
	}
	return roman
}

func stringToInt(str string) (result int, err error) {

	for i := 1; i <= 10; i++ {
		if str == strconv.Itoa(i) {
			return i, nil
		}
	}
	return 0, fmt.Errorf("неверный ввод арабских цифр: %s", str)
}

func calculate(input string) (result string, err error) {

	isArabic := strings.ContainsAny(input, "0123456789")
	isRoman := strings.ContainsAny(input, "IVX")

	if isArabic && isRoman {
		return "", fmt.Errorf("неверный формат выражения, римские и арабские вперемешку")
	}

	tokens := strings.Fields(input)
	if len(tokens) != 3 {
		return "", fmt.Errorf("неверный формат выражения")
	}

	if isRoman {
		a, err := romanToInt(tokens[0])
		if err != nil {
			return "", err
		}
		b, err := romanToInt(tokens[2])
		if err != nil {
			return "", err
		}
		rez, err := execute(a, b, tokens[1])
		if err != nil {
			return "", err
		}
		return intToRoman(rez), nil
	}

	if isArabic {
		a, err := stringToInt(tokens[0])
		if err != nil {
			return "", err
		}
		b, err := stringToInt(tokens[2])
		if err != nil {
			return "", err
		}
		rez, err := execute(a, b, tokens[1])
		if err != nil {
			return "", err
		}
		return strconv.Itoa(rez), nil
	}

	return "", fmt.Errorf("неверный формат выражения")
}

func main() {
	fmt.Println("Калькулятор умеет выполнять операции сложения, вычитания,  ")
	fmt.Println("умножения и деления с двумя числами: a + b, a - b, a * b, a / b.")
	fmt.Println("Калькулятор умеет работать как с арабскими (1, 2, 3, 4, 5...), ")
	fmt.Println("так и с римскими (I, II, III, IV, V...) числами от 1 до 10")
	fmt.Println("Калькулятор умеет работать только с арабскими или римскими цифрами одновременно")
	fmt.Println("Данные передаются в одну строку через пробел")
	fmt.Println("Введите выражение: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	result, err := calculate(text)
	if err != nil {
		panic(err)
	}

	fmt.Println("Результат : ", result)
}
