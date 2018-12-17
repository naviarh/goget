package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Список импортируемых пакетов
	packages := make(map[string]bool)
	// Список проверяемых файлов
	files := []string{}
	// Получение аргументов утилиты
	args := os.Args[1:]
	// Если указана точка
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
				// ----------- Циклический поиск в составном импорте
				if words[1] == "(" {
					for i := n; i < len(text); i++ {
						text[i] = strings.Trim(string(text[i]), " 	") // удаление пробелов и табов
						// Посимвольный перебор
						for s := 0; s < len(text[i]); s++ {
							switch string(text[i][s]) {
							case "\"": // начало имени импортируемого пакета
								packages[strings.Trim(string(text[i][s:]), "\" 	")] = false
								s = len(text[i]) - 1
							case "/": // комментарий
								s = len(text[i]) - 1
							case ")": // конец составного импорта
								i = len(text) - 1
							}
						}
					}
					// ------- Поиск в одиночном импорте
				} else {
					// Посимвольный перебор
					for s := 0; s < len(str); s++ {
						if string(str[s]) == "\"" {
							packages[strings.Trim(string(str[s:]), "\"")] = false
							s = len(str)
						}
					}
				}
			}
		}
	}
	// Обновление пакетов
	println()
	println("Found packages: ", len(packages))
	println()
	for pack, _ := range packages {
		print(pack, " ..   ")
		out, err := exec.Command("go", "get", "-u", "-v", pack).Output()
		if err != nil {
			println(err.Error())
			packages[pack] = true // метка ошибки для пакета
		} else {
			if string(out) != "" {
				println(string(out))
			} else {
				println("Ok")
			}
		}
	}
	println()
	// Вывод пакетов с ошибками обновления
	out := "Errors:\n\n"
	for pack, iserr := range packages {
		if iserr {
			out += pack + "\n"
		}
	}
	if out != "Errors:\n\n" {
		println(out)
	}
}
