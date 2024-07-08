package main

import (
 "context"
 "encoding/json"
 "fmt"
 "log"
 "net/http"
 "os"
 "os/signal"
 Filesistem "srv/server/filesystem"
 "syscall"
 "time"
)

// Config - структура для хранения конфигурации
type Config struct {
 Port int `json:"port"`
}

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
 w.Header().Set("Access-Control-Allow-Origin", "*")

 // Разбираем URL-адрес и извлекаем флаги
 query := r.URL.Query()

 rootFolder := query.Get("root")
 sortOption := query.Get("sort")

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
 
 w.Write(jsonBytes)
 w.Header().Set("Content-Type", "application/json")
}

func main() {

 // Читаем конфигурацию из файла
 file, err := os.Open("./config/config.json")
 if err != nil {
  log.Fatalf("Ошибка открытия конфигурации: %v", err)
 }
 defer file.Close()

 var Config Config
 decoder := json.NewDecoder(file)
 if err := decoder.Decode(&Config); err != nil {
  log.Fatalf("Ошибка разбора JSON: %v", err)
 }

 // Устанавливаем роутер
 http.HandleFunc("/", handleRequest)

 // Создаем сервер
 srv := &http.Server{
  Addr:    fmt.Sprintf(":%d", Config.Port),
  Handler: http.HandlerFunc(handleRequest),
 }

 // Запускаем сервер в отдельной горутине
 go func() {
  if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
   log.Fatalf("ListenAndServe: %v", err)
  }
 }()

 // Создаем канал для сигналов
 interrupt := make(chan os.Signal, 1)
 signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

 // Ожидаем сигнал прерывания
 <-interrupt

 // Даем серверу 5 секунд на завершение работы
 ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
 defer cancel()

 // Завершаем работу сервера
 srv.Shutdown(ctx)
 fmt.Println("Сервер остановлен")

}