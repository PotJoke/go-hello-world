package main

import (
	"firstmodule/tools"
	"fmt"
)

func calc(sword int, shield int) bool {
	return sword > shield
}

func main() {
	tools.ClearScreen()
	fmt.Println("Приветствую в моей первой программе на golang")
	tools.WaitForEnter()

	tools.ClearScreen()
	fmt.Print(`Пока она умеет немного, я в основном разбирался \n
	с дурацкими импортами итд`)
	tools.WaitForEnter()

	tools.ClearScreen()
	fmt.Println("Давай попробуем команду Input")
	fmt.Println("Вы ввели: " + tools.Input("Введи что нибудь: "))
	tools.WaitForEnter()

	tools.ClearScreen()
	fmt.Println("Что ж, жду тебя в следующих релизах")
	tools.WaitForEnter()
}
