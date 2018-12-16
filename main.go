package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Список импортируемых пакетов
	packages := []string{}
	// Список проверяемых файлов
	files := []string{}
	// Получение аргументов утилиты
	args := os.Args[1:]
	// Еси указана точка
	if args[0] == "." {
		// Получениие списка файлов с расширением .go
		directory, err := ioutil.ReadDir(".")
		if err != nil {
			panic(err)
		}
		// Составление списка файлов *.go
		for _, file := range directory {
			if !file.IsDir() {
				name := strings.Split(file.Name(), ".")
				if name[len(name)-1] == "go" {
					files = append(files, file.Name())
				}
			}
		}
		// Если указаны файлы
	} else {
		// Составление списка файлов *.go
		for _, arg := range args {
			name := strings.Split(arg, ".")
			// Проверка аргумента на файл .go
			if name[len(name)-1] == "go" {
				// Проверка на наличие файла в каталоге
				if _, err := os.Stat(arg); err == nil {
					files = append(files, arg)
				}
			}
		}
	}
	// Если список файлов пуст - выход
	if len(files) == 0 {
		println("No files!")
		os.Exit(0)
	}
	for _, file := range files {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			println("I can not read the file", file)
			continue
		}
		// Получеине текста по строкам
		text := strings.Split(string(bytes), "\n")
		// Построчный поиск команды импортирования
		for n, str := range text {
			// Проверка строки на команду импорта
			words := strings.Split(str, " ")
			if words[0] == "import" {
				// Циклический поиск в составном импорте
				if words[1] == "(" {
					for i := n; i < len(text); i++ {
						// Посимвольный поиск
						for s := 0; s < len(text[i]); s++ {
							switch string(text[i][s]) {
							case ")": // конец составного импорта
								i = len(text) - 1
							case "\"": // начало имени импортируемого пакета
								packages = append(packages, "")
								// Получение имени импортируемого пакета
								for j := s + 1; j < len(text[i]); j++ {
									s++ // инкремент счётчика символов текущей строки
									// Проверка конца имени пакета
									if string(text[i][j]) == "\"" {
										break
									}
									// Очередная буква имени пакета
									packages[len(packages)-1] += string(text[i][j])
								}
							}
						}
					}
					// Поиск в одиночном импорте
				} else {
					println(str)
				}
			}
		}
	}
	fmt.Println(packages)
}
