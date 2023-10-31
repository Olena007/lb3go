package filemanager

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Структура для представлення інформації
type Information struct {
	Carrier string
	Capacity string
	Title string
	Author string
}

// Функція для створення файлу та запису структурованих даних у нього
func CreateFileWithData(filename string, data []Information) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, item := range data {
		_, err := fmt.Fprintf(writer, "Носій: %s\nОб'єм: %s\nНазва: %s\nАвтор: %s\n\n", item.Carrier, item.Capacity, item.Title, item.Author)
		if err != nil {
			return err
		}
	}

	return nil
}

// Функція для виведення вмісту файлу на екран
func PrintFileContents(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}

	return nil
}

// Функція для видалення першого елемента із заданим обсягом інформації
func DeleteFirstItemWithCapacity(filename string, capacityToDelete string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tempFile, err := os.Create("tempfile.txt")
	if err != nil {
		return err
	}
	defer tempFile.Close()

	writer := bufio.NewWriter(tempFile)
	defer writer.Flush()

	var found bool

	for scanner.Scan() {
		line := scanner.Text()
		if line == "Об'єм: " + capacityToDelete {
			found = true
		}
		if !found {
			_, err := fmt.Fprintln(writer, line)
			if err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	err = os.Remove(filename)
	if err != nil {
		return err
	}

	err = os.Rename("tempfile.txt", filename)
	if err != nil {
		return err
	}

	return nil
}

// Функція для додавання K елементів у кінець файлу
func AddKItemsToEndOfFile(filename string, dataToAdd []Information) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, item := range dataToAdd {
		_, err := fmt.Fprintf(writer, "Носій: %s\nОб'єм: %s\nНазва: %s\nАвтор: %s\n\n", item.Carrier, item.Capacity, item.Title, item.Author)
		if err != nil {
			return err
		}
	}

	return nil
}
