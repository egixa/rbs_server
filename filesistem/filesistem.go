package filesistem

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

const asc = "asc"
const desc = "desc"

// flagParse считывает флаги из строки терминала
func flagParse() (string, string, error) {
	rootFolder := flag.String("root", "", "Путь до директории для вывода структуры\n")
	sortOption := flag.String("sort", "", "Параметр сортировки:\n по убыванию -\n по возрастанию -")
	flag.Parse()

	// Проверка флагов на корректность ввода
	if *rootFolder == "" {
		fmt.Println(time.Now().Format("01-02-2006 15:04:05"), "Отсутствуют данные о местоположении директории.")
		fmt.Println("Ожидаемые параметры вызова программы:")
		flag.PrintDefaults()
		return "", "", fmt.Errorf(fmt.Sprint("Отсутствуют данные о местоположении директории.\nОжидаемые параметры вызова программы:", rootFolder, sortOption))
	}
	if *sortOption != asc && *sortOption != desc {
		*sortOption = asc
		fmt.Println("Введен некорректный параметр сортировки. По умолчанию будет использована сортировка по возрастанию.")
	}

	// Проверка существования директории
	_, err := os.Stat(*rootFolder)
	if err != nil {
		if os.IsNotExist(err) {
			return "", "", fmt.Errorf("Ошибка при обнаружении директории:", err)
		}
	}
	return *rootFolder, *sortOption, nil
}

// sortDirectory сортирует директории по входному параметру
func sortDirectory(directoryContent []File, sortOption string) []File {
	switch sortOption {
	case desc:
		sort.Slice(directoryContent, func(i, j int) (less bool) {
			return directoryContent[i].Size > directoryContent[j].Size
		})
	case asc:
		sort.Slice(directoryContent, func(i, j int) (less bool) {
			return directoryContent[i].Size < directoryContent[j].Size
		})
	}
	return directoryContent
}

// formatSize форматирует размер файлов из байтов в килобайты, мегабайты и гигабайты
func formatSize(size int64) string {
	const gigabyte = 1000 * 1000 * 1000
	const megabyte = 1000 * 1000
	const kilobyte = 1000

	if size > gigabyte {
		return fmt.Sprintf("%.2f гб", float64(size)/(gigabyte))
	} else if size > megabyte {
		return fmt.Sprintf("%.2f мб", float64(size)/(megabyte))
	} else if size > kilobyte {
		return fmt.Sprintf("%.2f кб", float64(size)/(kilobyte))
	}
	return fmt.Sprintf("%d б", size)
}

// Создание массива информации о файлах
type File struct {
	Type string
	Name string
	Size int64
}

// dirSize определяет размер директории
func dirSize(path string) (int64, error) {

	var size int64
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Ошибка при обходе директории:", err)
			return err
		}
		size += info.Size()
		return nil
	})
	return size, err
}

func GetFolders() {
	start := time.Now()

	//Считываем флаги
	rootFolder, sortOption, err := flagParse()
	if err != nil {
		panic(err)
	}

	// Открываем директорию
	dir, err := os.Open(rootFolder)
	if err != nil {
		panic("Ошибка при открытии директории")
	}
	defer dir.Close()

	// Получаем список файлов и директорий
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Ошибка при прочтении директории", err)
		return
	}

	// Создаем массив структур с информацией о содержании директории
	directoryContent := []File{}

	var wg sync.WaitGroup
	// Записываем имена, размеры файлов и директорий в массив структур
	for _, file := range files {
		fileName := file.Name()
		fileSize := file.Size()
		filePath := filepath.Join(rootFolder, fileName)

		if file.IsDir() {
			wg.Add(1)

			go func() {
				defer wg.Done()

				dirSize, err := dirSize(filePath)

				if err != nil {
					return
				}

				directoryContent = append(directoryContent, File{"Директория", fileName, dirSize})
			}()

		} else {
			directoryContent = append(directoryContent, File{"Файл", fileName, fileSize})
			continue
		}
	}
	wg.Wait()

	// Сортируем содержимое директории по указанному параметру
	directoryContent = sortDirectory(directoryContent, sortOption)

	// Выводим содержимое в консоль
	for _, file := range directoryContent {
		fmt.Printf("%s %s Размер: %s\n", file.Type, file.Name, formatSize(file.Size))
	}

	// Время выполнения
	duration := time.Since(start)
	fmt.Println("Программа завершена. Время выполнения:", duration)
}
