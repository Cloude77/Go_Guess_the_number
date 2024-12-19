package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings" // убираем символ
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)   // чтение с клавиатуры
	input, err := reader.ReadString('\n') // возвращает текст и возврат ошибки. Если nill - ошибки нет
	if err != nil {
		log.Fatal(err) // сообщить об ошибке и остановить
	}

	input = strings.TrimSpace(input)            // убираем переход строки
	grade, err := strconv.ParseFloat(input, 64) // преобразуем сроку в float64
	{
		if err != nil {
			log.Fatal(err)
		}

	}

	var status string
	if grade >= 60 {
		status = "passing" // замена := На =
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status) // вывод текста

}
