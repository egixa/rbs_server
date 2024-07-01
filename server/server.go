package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	Filesistem "github.com/filesistem"
	//"github.com/filesistem"
)

const asc = "asc"
const desc = "desc"

func handleRequest(w http.ResponseWriter, r *http.Request) (string, string, error) {
	// Разбирает URL-адрес
	query := r.URL.Query()
	/*if err != nil {
	 http.Error(w, "Ошибка разбора URL-адреса", http.StatusBadRequest)
	 return
	}*/

	// Извлекает флаги из строки запроса
	rootFolder := query.Get("root")
	sortOption := query.Get("sort")

	// Валидация флагов
	if rootFolder == "" {
		fmt.Println(time.Now().Format("01-02-2006 15:04:05"), "Отсутствуют данные о местоположении директории.")
		return "", "", fmt.Errorf(fmt.Sprint("Отсутствуют данные о местоположении директории.\nОжидаемые параметры вызова программы:", rootFolder, sortOption))
	}
	if sortOption != asc && sortOption != desc {
		sortOption = asc
		fmt.Println("Введен некорректный параметр сортировки. По умолчанию будет использована сортировка по возрастанию.")
	}

	// Проверка существования директории
	_, err := os.Stat(rootFolder)
	if err != nil {
		if os.IsNotExist(err) {
			return "", "", fmt.Errorf("Ошибка при обнаружении директории:", err)
		}
	}
	return rootFolder, sortOption, nil
}

func main() {
	rootFolder, sortOption, err := handleRequest(w http.ResponseWriter, r *http.Request)
	
	http.HandleFunc("/", Filesistem.GetFolders(rootFolder, sortOption))      // Устанавливаем роутер
	err := http.ListenAndServe(":3000", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
