package parse

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"parser/pkg/utils"
	"sort"
	"strconv"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
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

	defer writer.Flush()

	datesCSV := make([]DataCSV, 0, 500)

	// Описываем конфигурации селениума
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
	}})

	driver, err := selenium.NewRemote(caps, "")

	if err != nil {
		log.Fatal("Error:", err)
	}

	err = driver.MaximizeWindow("")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Начинаем сборку данных в цикле
	for i := 1; i <= 6; i++ {
		// Переходим на каждую страницу
		err := driver.Get(fmt.Sprintf("https://global.wildberries.ru/catalog/muzhchinam/odezhda/verhnyaya-odezhda?page=%v", i))
		if err != nil {
			log.Print("Error:", err)
		}

		log.Printf("Scraping page: %v", i)

		utils.RandomDelay(1000, 5000) // Первая задержка

		products, err := driver.FindElements(selenium.ByCSSSelector, "div.product-card")
		if err != nil {
			log.Fatal("Error finds products:", err)
		}

		for _, product := range products {
			nameElem, err := product.FindElement(selenium.ByCSSSelector, "div.product-card__title-text")
			if err != nil {
				log.Printf("Error nameElem: %v", err)
			}

			priceElem, err := product.FindElement(selenium.ByCSSSelector, "span.price__lower")
			if err != nil {
				log.Printf("Error priceElem: %v", err)
			}

			name, err := nameElem.Text()
			if err != nil {
				log.Printf("Error name: %v", err)
			}

			price, err := priceElem.Text()
			if err != nil {
				log.Printf("Error price: %v", err)
			}

			data := DataCSV{}
			data.Name = name
			data.Price = price

			datesCSV = append(datesCSV, data)
		}

		utils.RandomDelay(2000, 4000) // Ждем немного перед закрытием

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
