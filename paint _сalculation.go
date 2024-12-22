package main

import "fmt"

/*************  ✨ Codeium Command ⭐  *************/

// Paint calculator program
//
// Calculates how much paint is needed to paint a
// rectangular room with user-provided dimensions.
//
// The calculation is simply the area of the room
// divided by 10 (an estimate of how much paint is
// needed to paint the room).
//
/******  6ea3acd6-2ad5-4879-8e8b-92fbf3ffd075  *******/

func paintNeed(width, height float64) (float64, error) { // возникли ли ошибки
	if width < 0 { // если ширина меньше 0
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 { // если высота меньше 0
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil // отсутствие ошибки
}

/*************  ✨ Codeium Command ⭐  *************/
// main is the entry point of the paint calculator program.
// It calculates and prints the amount of paint needed for
// two rectangular rooms with given dimensions, as well as
// the total amount of paint needed for both rooms.

/******  81048b7c-5a7d-4c93-b360-f704a51917c9  *******/
func main() {
	var amount, total float64
	amount, err := paintNeed(4.2, 3.0)
	fmt.Println(err) // выводим ошибку (или nil если ошибок нет)
	fmt.Printf("Amount of paint needed: %.2f\n", amount)

	total += amount

	amount, err = paintNeed(3.0, 2.5) // -2/5 ошибка
	fmt.Println(err)
	fmt.Printf("%0.2f liters needed\n", amount)

	total += amount

	fmt.Printf("Total amount of paint needed: %.2f\n", total)

}
