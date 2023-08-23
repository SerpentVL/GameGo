package main

import (
	"fmt"
	"os"
	"strconv"
)

func Start() string {
	fmt.Println("Я буду задавать вопросы, а ты будешь мне отвечать на клавиатуре. ")
	fmt.Println("По-английски. ")
	fmt.Println("Если твой ответ - да,  то набирай на клавиатуре yes ")
	fmt.Println("Если твой ответ - нет, то набирай на клавиатуре no ")
	fmt.Println("И потом нажми Enter")
	fmt.Println("Ты поняла ?")
	fmt.Println("Ответь yes (поняла и хочешь поиграть)")
	fmt.Println("Ответь no  (играть не хочешь)")

	var answer string
	fmt.Scan(&answer)
	return answer
}

// -- двоичный поиск задуманного числа --
func Work(number int) int {
	var first int = 1
	var last int = number
	var middle int
	var answer string
	for first < last {
		middle = first + (last-first)/2
		//fmt.Println(first, middle, last)
		fmt.Println("Ты задумала число БОЛЬШЕ ", middle, "?")
		fmt.Scan(&answer)
		if answer == "yes" {
			first = middle + 1
		} else if answer == "no" {
			last = middle
		}
	}
	return first
}

func main() {
	var number int = 30

	if len(os.Args) > 1 {
		//fmt.Println(os.Args[1])
		n, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err == nil {
			if n > 0 {
				number = int(n)
			}
		}
	}
	//fmt.Println(number)

	fmt.Println("--------------")
	fmt.Println("Привет, Анюта!")
	fmt.Println("--------------")
	fmt.Println("Задумай число от 1 до", number)
	fmt.Println("Я его угадаю. ")
	fmt.Println("--------------")
	var answer string
	answer = Start()

	if answer == "yes" {
		fmt.Println("--------------")
		fmt.Println("Погнали!")
		Number := Work(number)
		fmt.Println("Ты задумала число", Number)
	} else if answer == "no" {
		fmt.Println("--------------")
		fmt.Println("До свидания!")
	} else {
		fmt.Println("--------------")
		fmt.Println("Ты ошиблась. Давай попробуем еще раз.")
		answer = Start()
		if answer == "yes" {
			fmt.Println("--------------")
			fmt.Println("Погнали!")
			Number := Work(number)
			fmt.Println("Ты задумала число", Number)
		} else if answer == "no" {
			fmt.Println("--------------")
			fmt.Println("До свидания!")
		} else {
			fmt.Println("Ты устала. Отдохни. Потом попробуй еще раз.")
			fmt.Println("До свидания!")
		}
	}
	fmt.Println("--------------")
	fmt.Println("----- КОНЕЦ -----")
}
