package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToArabic(roman string) int {
	romanMap := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	return romanMap[roman]
}

func arabicToRoman(arabic int) string {
	arabicSlice := []int{arabic / 10 * 10, arabic % 10}
	romanMap := map[int]string{
		0:  "",
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
		20: "XX",
		30: "XXX",
		40: "XL",
		50: "L",
		60: "LX",
		70: "LXX",
		80: "LXXX",
		90: "XC",
	}
	return romanMap[arabicSlice[0]] + romanMap[arabicSlice[1]]
}

func calculate(a, b int, operator rune) int {
	switch operator {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		if b == 0 {
			panic("Деление на ноль.")
		}
		return a / b
	default:
		panic("Неподдерживаемая операция.")
	}
}

func main() {
	var input string
	fmt.Println("Добро пожаловать в калькулятор на Go!")
	fmt.Print("Введите выражение (например, 2 + 3): ")
	input, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	inputSlice := strings.Fields(input)
	if len(inputSlice) != 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию - два операнда и один оператор (+, -, /, *).")
	}
	var a, b int
	var useRoman bool

	if num, err := strconv.Atoi(inputSlice[0]); err == nil {
		a = num
		useRoman = false
	} else {
		a = romanToArabic(inputSlice[0])
		useRoman = true
	}

	if num, err := strconv.Atoi(inputSlice[2]); err == nil {
		b = num
		if useRoman {
			panic("Выдача паники, так как используются одновременно разные системы счисления.")
		}
	} else {
		b = romanToArabic(inputSlice[2])
		if !useRoman {
			panic("Выдача паники, так как используются одновременно разные системы счисления.")
		}
	}

	if a < 1 || b < 1 || a > 10 || b > 10 {
		panic("Для расчета принимаются числа от 1 до 10.")
	}

	result := calculate(a, b, []rune(inputSlice[1])[0])

	if useRoman {
		if result < 1 {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
		fmt.Println(arabicToRoman(result))
	} else {
		fmt.Println(result)
	}
}
