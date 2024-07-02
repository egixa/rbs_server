package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	Filesistem "svr/filesistem"
)

const asc = "asc"
const desc = "desc"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Разбирает URL-адрес
	query := r.URL.Query()

	// Извлекает флаги из строки запроса
	rootFolder := query.Get("root")
	sortOption := query.Get("sort")

	// Валидация флагов
	if rootFolder == "" {
		panic("Отсутствуют данные о местоположении директории")
	}
	if sortOption != asc && sortOption != desc {
		sortOption = asc
		fmt.Println("Введен некорректный параметр сортировки. По умолчанию будет использована сортировка по возрастанию.")
	}

	// Проверка существования директории
	_, err := os.Stat(rootFolder)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "Такой директории не существует", http.StatusNotFound)
		} else {
			panic("Ошибка при обнаружении директории")
		}
	}

	files := Filesistem.GetFolders(rootFolder, sortOption)
	jsonBytes, err := json.Marshal(files)
	if err != nil {
		panic("Ошибка при преобразовании в json")
	}
	fmt.Fprintf(w, string(jsonBytes))
}

func main() {
	http.HandleFunc("/", handleRequest)      // Устанавливаем роутер
	err := http.ListenAndServe(":3000", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
