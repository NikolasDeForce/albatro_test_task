package excelgenerator

import (
	"bytes"
	"parser/pkg/methods"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func ExcelGenerator(data methods.Excel) (*bytes.Buffer, error) {
	//Создаём новый файл
	file := excelize.NewFile()

	sheetName := "Sheet1"

	// Получение заголовков
	header := data.Header()

	// Получение строк заголовков
	rows := data.HeaderRows()

	// Получение заголовков подзаголовков
	titles := data.Title()

	// Получение строк подзаголовков
	titleRows := data.TitleRows()

	// Создание заголовков
	for i, h := range titles {
		colName, _ := excelize.ColumnNumberToName(1)
		// Установка значения заголовка в ячейку
		file.SetCellValue(sheetName, colName+strconv.Itoa(i+1), h)
	}

	// Заполнение заголовков
	for i, h := range titleRows {
		colName, _ := excelize.ColumnNumberToName(2)
		// Установка значения заголовка в ячейку
		file.SetCellValue(sheetName, colName+strconv.Itoa(i+1), h)
	}

	// Создание заголовков колонок
	for i, h := range header {
		colName, _ := excelize.ColumnNumberToName(i + 1)
		// Установка значения заголовка колонки в ячейку
		file.SetCellValue(sheetName, colName+"7", h)
	}

	// Заполнение колонок
	for r, row := range rows {
		for c, val := range row {
			colName, _ := excelize.ColumnNumberToName(c + 1)
			// Установка значения в соответствующую ячейку
			file.SetCellValue(sheetName, colName+strconv.Itoa(r+8), val)
		}
	}

	f, _ := file.WriteToBuffer()

	return f, nil
}
