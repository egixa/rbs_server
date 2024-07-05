package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

const Asc = "asc"
const Desc = "desc"

// sortDirectory сортирует директории по входному параметру
func sortDirectory(directoryContent []File, sortOption string) []File {

	switch sortOption {
	case Desc:
		sort.Slice(directoryContent, func(i, j int) (less bool) {
			return directoryContent[i].Size > directoryContent[j].Size
		})
	case Asc:
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
	Type       string `json:"type"`
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	SizeFormat string `json:"formatingSize"`
}

// dirSize определяет размер директории
func dirSize(path string) (int64, error) {
	var sizeSumm int64
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("Ошибка при обходе директории:%w", err)
		}
		sizeSumm += info.Size()
		return nil
	})
	return sizeSumm, err
}

// GetFolder сортирует директорию и выводит информацию о размере содержимого
func GetFolder(rootFolder string, sortOption string) ([]File, error) {

	// Открываем директорию
	dir, err := os.Open(rootFolder)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при открытии директории%w ", err)
	}
	defer dir.Close()

	// Получаем список файлов и директорий
	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при прочтении директории%w ", err)
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

				directoryContent = append(directoryContent, File{"Директория", fileName, dirSize, formatSize(dirSize)})
			}()
		} else {
			directoryContent = append(directoryContent, File{"Файл", fileName, fileSize, formatSize(fileSize)})
			continue
		}
	}
	wg.Wait()

	// Сортируем содержимое директории по указанному параметру
	directoryContent = sortDirectory(directoryContent, sortOption)

	return directoryContent, nil
}
