package parse

import (
	"fmt"
	"log"
	"os"
	"parser/pkg/excelgenerator"
	"parser/pkg/methods"
	"parser/pkg/utils"
	"sort"
	"strconv"

	"github.com/gocolly/colly"
)

func ParseToXLS(fname string, sorting string) {
	datesXLS := methods.DataXLS{}

	datesXLS.DataXLS = make([]struct {
		Name  string
		Price string
	}, 0)

	//Создаём файл
	data, err := os.Create(fmt.Sprintf("./dates/%v.xls", fname))
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer data.Close()

	// Инициализируем коллектор
	c := colly.NewCollector()

	for i := 1; i < 6; i++ {
		// Вытаскиваем данные
		c.OnHTML("div.InternalProductCard__container___IPYI4", func(e *colly.HTMLElement) {
			data := DataCSV{
				e.ChildText("p.Title__title___ujUQX"),
				e.ChildText("div.Prices__wrapper___EeU1W"),
			}

			datesXLS.DataXLS = append(datesXLS.DataXLS, data)
		})

		c.Visit(fmt.Sprintf("https://www.tsum.ru/catalog/sumki-18438/?page=%v", i))

		log.Printf("Scraping page: %v", i)
	}
	//Проверяем переменную на вид сортировки и сортируем
	switch {
	case sorting == "asc":
		sort.SliceStable(datesXLS.DataXLS, func(i, j int) bool { return datesXLS.DataXLS[i].Name < datesXLS.DataXLS[j].Name })
	case sorting == "desc":
		sort.SliceStable(datesXLS.DataXLS, func(i, j int) bool { return datesXLS.DataXLS[i].Name > datesXLS.DataXLS[j].Name })
	}

	fileXLS, err := excelgenerator.ExcelGenerator(&datesXLS)

	if err != nil {
		log.Println(err)
		return
	}

	// Записываем содержимое в файл
	_, err = fileXLS.WriteTo(data)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	log.Println("Scraping finished, check the files!")

	// Считаем медианное значение и выводим в консоль
	var mediana int
	var sliceMedian []int

	for key := range datesXLS.DataXLS {
		price := datesXLS.DataXLS[key].Price
		priceWoutSpace := utils.RemoveSpace(price)
		priceWoutSpace = utils.TrimSuffix(priceWoutSpace, "₽")
		mediana, err = strconv.Atoi(priceWoutSpace)
		if err != nil {
			log.Fatalf("Median parse error %v\n", err)
		}
		sliceMedian = append(sliceMedian, mediana)
	}
	numMedian := utils.Median(sliceMedian)

	fmt.Printf("Медианное значение: %v\n", numMedian)
}
