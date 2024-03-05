package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanMap = []struct {
	decVal int
	symbol string
}{
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

func isRomanian(x string) bool {
	rom := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for i := 0; i < len(rom); i++ {
		if rom[i] == x {
			return true
		}
	}
	return false
}

func isArabian(x string) bool {
	arab := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for i := 0; i < len(arab); i++ {
		if arab[i] == x {
			return true
		}
	}
	return false
}

func intToRomanRecursive(x int) string {
	if x == 0 {
		return ""
	}
	for _, pair := range romanMap {
		if x >= pair.decVal {
			return pair.symbol + intToRomanRecursive(x-pair.decVal)
		}
	}
	return ""
}

func intToStr(x int, flag int) string {
	var st string = ""
	if flag == 0 {
		st = strconv.Itoa(x)
	} else if flag == 1 {
		if x < 1 {
			panic("В римской системе нет отрицательных чисел.")
		} else {
			st = intToRomanRecursive(x)
		}
	}
	return st
}

func strToInt(x string, flag int) int {
	var k int = 0
	rom := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	arab := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	if flag == 0 {
		for i := 0; i < len(arab); i++ {
			if arab[i] == x {
				k = i + 1
				break
			}
		}
	} else if flag == 1 {
		for i := 0; i < len(rom); i++ {
			if rom[i] == x {
				k = i + 1
				break
			}
		}
	}
	return k
}

func alert(x []string) int {
	var flag int = 0
	var flag1 int = 0
	if len(x) > 3 {
		panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	} else if len(x) < 3 {
		panic("Cтрока не является математической операцией.")
	} else {
		if isRomanian(x[0]) && isRomanian(x[2]) {
			flag = 1
		} else if isArabian(x[0]) && isArabian(x[2]) {
			flag = 0
		} else if (isArabian(x[0]) && isRomanian(x[2])) || (isRomanian(x[0]) && isArabian(x[2])) {
			panic("Используются одновременно разные системы счисления.")
		} else {
			panic("Числа выходят из диапозона допустимых.")
		}
		var oper []string = []string{"+", "-", "/", "*"}
		for i := 0; i < len(oper); i++ {
			if oper[i] == x[1] {
				flag1 = 1
				break
			}
		}
		if flag1 == 0 {
			panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		}
	}
	return flag
}

func operation(x []string, flag int) string {
	switch x[1] {
	case "+":
		return intToStr(strToInt(x[0], flag)+strToInt(x[2], flag), flag)
	case "-":
		return intToStr(strToInt(x[0], flag)-strToInt(x[2], flag), flag)
	case "*":
		return intToStr(strToInt(x[0], flag)*strToInt(x[2], flag), flag)
	case "/":
		return intToStr(strToInt(x[0], flag)/strToInt(x[2], flag), flag)
	default:
		return ""
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var flag int = 0
	var result string = ""
	var arrtext []string
	for {
		fmt.Println("Введите выражение:")
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")
		arrtext = strings.Split(text, " ")
		flag = alert(arrtext)
		result = operation(arrtext, flag)
		fmt.Println(result)
	}
}
