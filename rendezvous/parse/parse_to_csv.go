package parse

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"parser/pkg/utils"
	"sort"
	"strconv"

	"github.com/gocolly/colly"
)

type DataCSV struct {
	Name  string
	Price string
}

func ParseToCSV(fname string, sorting string) {
	// Создаем файл и записываем
	file, err := os.Create(fmt.Sprintf("./dates/%v.csv", fname))
	if err != nil {
		log.Fatalf("Cannot create file %v: %s\n", file, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"name",
		"price",
	}

	writer.Write(headers)

	// Инициализируем коллектор
	c := colly.NewCollector()

	datesCSV := make([]DataCSV, 0, 500)

	for i := 1; i < 6; i++ {
		// Вытаскиваем данные
		c.OnHTML("li.item", func(e *colly.HTMLElement) {
			data := DataCSV{
				e.ChildText("div.item-name"),
				e.ChildText("span.item-price-new"),
			}

			datesCSV = append(datesCSV, data)
		})

		c.Visit(fmt.Sprintf("https://rendez-vous.ru/catalog/female/page/%v/", i))

		log.Printf("Scraping page: %v", i)
	}

	//Проверяем переменную на вид сортировки и сортируем
	switch {
	case sorting == "asc":
		sort.SliceStable(datesCSV, func(i, j int) bool { return datesCSV[i].Name < datesCSV[j].Name })

		for key := range datesCSV {
			r := make([]string, 0, len(headers))
			r = append(
				r,
				datesCSV[key].Name,
				datesCSV[key].Price,
			)

			writer.Write(r)
		}
	case sorting == "desc":
		sort.SliceStable(datesCSV, func(i, j int) bool { return datesCSV[i].Name > datesCSV[j].Name })

		for key := range datesCSV {
			r := make([]string, 0, len(headers))
			r = append(
				r,
				datesCSV[key].Name,
				datesCSV[key].Price,
			)

			writer.Write(r)
		}
	default:
		for key := range datesCSV {
			r := make([]string, 0, len(headers))
			r = append(
				r,
				datesCSV[key].Name,
				datesCSV[key].Price,
			)

			writer.Write(r)
		}
	}

	log.Println("Scraping finished, check the files!")

	//Считаем медианное значение и выводим в консоль
	var mediana int
	var sliceMedian []int

	for key := range datesCSV {
		price := datesCSV[key].Price
		priceWoutSpace := utils.RemoveSpace(price)
		mediana, err = strconv.Atoi(priceWoutSpace)
		if err != nil {
			log.Fatalf("Median parse error %v\n", err)
		}
		sliceMedian = append(sliceMedian, mediana)
	}
	numMedian := utils.Median(sliceMedian)

	fmt.Printf("Медианное значение: %v\n", numMedian)
}
