package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings" // убираем символ
)

/*************  ✨ Codeium Command ⭐  *************/
// getFloat reads a line of input from the standard input, trims any
// whitespace, and attempts to parse it as a float64. It returns the
// parsed float64 value and an error if any occurred during reading
// or parsing.

/******  ff4d77b7-6e8f-4af3-9e1e-cba06a6f4d5a  *******/
func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)   // чтение с клавиатуры
	input, err := reader.ReadString('\n') // возвращает текст и возврат ошибки. Если nill - ошибки нет
	if err != nil {
		return 0, err // сообщить об ошибке и остановить
	}

	input = strings.TrimSpace(input)             // убираем переход строки
	number, err := strconv.ParseFloat(input, 64) // преобразуем сроку в float64
	if err != nil {
		return 0, err
	}

	return number, nil

}

/*************  ✨ Codeium Command ⭐  *************/
// main is the entry point of the program that prompts the user to enter a grade,
// determines whether the grade is passing or failing, and prints the result.

/******  7be0d6a1-90dc-4212-afa6-2b54870ee57c  *******/
func main() {
	fmt.Print("Enter a grade: ")
	grade, err := getFloat() // вызывает функцию getFloat
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing" // замена := На =
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status) // вывод текста

}
