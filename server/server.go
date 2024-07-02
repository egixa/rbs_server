package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	Filesistem "svr/filesistem"
)

// validateArgs проверяет валидные или невалидные флаги
func validateArgs(rootFolder string, sortOption string) (string, error) {
	// Валидация флагов
	if rootFolder == "" {
		return "", fmt.Errorf("Отсутствуют данные о местоположении директории")
	}
	if sortOption != Filesistem.Asc && sortOption != Filesistem.Desc {
		sortOption = Filesistem.Asc
		fmt.Println("Введен некорректный параметр сортировки. По умолчанию будет использована сортировка по возрастанию.")
		return sortOption, nil
	}

	// Проверка существования директории
	_, err := os.Stat(rootFolder)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("Такой директории не существует", err)
		} else {
			return "", fmt.Errorf("Ошибка при обнаружении директории", err)
		}
	}
	return sortOption, nil
}

// handleRequest принимает ответ от сервера и отправляет отсортированный массив с информацией о содержимом
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Разбираем URL-адрес и извлекаем флаги
	query := r.URL.Query()

	rootFolder := query.Get("root")
	sortOption := query.Get("sort")

	fmt.Println("Host:", query.Get("localhost"))

	// Проверяем валидность флагов
	sortOption, err := validateArgs(rootFolder, sortOption)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Получаем отсортированное содержимое директории
	files, err := Filesistem.GetFolder(rootFolder, sortOption)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Сериализуем данные
	jsonBytes, err := json.Marshal(files)
	if err != nil {
		http.Error(w, "Ошибка при преобразовании в json", http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	/*// Отправляем данные на сервер
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
	 log.Printf("error while writing: %v", err)
	 return
	}*/

	//fmt.Fprintf(w, string(jsonBytes))

}

func main() {

	// Устанавливаем роутер
	http.HandleFunc("/", handleRequest)

	// устанавливаем порт веб-сервера
	err := http.ListenAndServe(":3007", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
