package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const asc = "asc"
const desc = "desc"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Разбираем URL-адрес
	query := r.URL.Query()

	// Извлекаем флаги из строки запроса
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
			panic("Такой директории не существует")
		} else {
			panic("Ошибка при обнаружении директории")
		}
	}
	return

	GetFolders()
}

func main() {
	http.HandleFunc("/", handleRequest)      // Устанавливаем роутер
	err := http.ListenAndServe(":3000", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
