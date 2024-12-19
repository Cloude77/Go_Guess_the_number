package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() // текущее время дата в фолрмате целого числа
	rand.Seed(seconds)           // случайные числа
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number beetwen 1 and 100. ")
	fmt.Println("Can you guess it?")
	//fmt.Println(target) чтобы не сообщать заранее число

	reader := bufio.NewReader(os.Stdin) // чтение с клавиатуры

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.") // оставшиеся попытки dвычитается из 10

		fmt.Println("Make a guess:")
		input, err := reader.ReadString('\n') // Прочитать данные введеные пользователем возвращает текст и возврат ошибки. Если nill - ошибки нет
		if err != nil {
			log.Fatal(err) // сообщить об ошибке и остановить
		}

		input = strings.TrimSpace(input)  // убираем переход строки
		guess, err := strconv.Atoi(input) // преобразуем сроку в int
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. Your guess is too LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess is too HIGH.")
		} else {
			success = true // если угадал - success = true
			fmt.Println("Good job! You guessed it!")
			break // выход из цикла , если угадал
		}

	}

	if !success { // если не угадал
		fmt.Println("Sorry, you didn't guess my number. It was:", target) // рандомное число
	}
}
