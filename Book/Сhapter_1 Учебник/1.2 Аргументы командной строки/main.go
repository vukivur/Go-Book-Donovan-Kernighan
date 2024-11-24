// В данном пакете выведем несколькими способами аргументы командной строки
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(command_arguments_1())
	fmt.Println(command_arguments_2())
	fmt.Println(command_arguments_3())
	fmt.Println(command_arguments_4())
}

// Вариант 1. функция выводит аргументы командной строки через цикл for
func command_arguments_1() string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] // сперва переменная sep инициализируется пустой строкой "" по умолчанию
		sep = " "
	}
	return s
}

// Вариант 2. Используя range
func command_arguments_2() string {
	s_1, sep_1 := "", ""
	for _, arg := range os.Args[1:] { // используется пустой идентификатор с именем: _
		s_1 += sep_1 + arg
		sep_1 = " "
	}
	return s_1
}

// Вариант 3. Используя функцию из пакета strings
func command_arguments_3() string {
	return strings.Join(os.Args[1:], " ")
}

// Вариант 4. Если нам не нужно беспокоиться о формате и нужно увидеть только значения
func command_arguments_4() []string {
	s_2 := os.Args[1:]
	return s_2
}

/*
Традиционный цикл while: for condition {...
Традиционный бесконечный цикл: for {...
Завершить бесконечный цикл можно при помощи: break или return
*/
