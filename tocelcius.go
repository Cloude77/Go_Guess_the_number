package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func main() {

	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := getFloat() // вызывает функцию getFloat для ввода температуры
	if err != nil {
		log.Fatal(err)
	}

	celsius := (fahrenheit - 32) * 5 / 9 // температура преобразуется к градусах Цельсия
	fmt.Printf("%0.2f degrees Celsius\n", celsius)

}
