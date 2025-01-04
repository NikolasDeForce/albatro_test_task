package methods

type Excel interface {
	Header() []string
	Title() []string
	TitleRows() []interface{}
	HeaderRows() [][]interface{}
}

type DataXLS struct {
	DataXLS []struct {
		Name  string
		Price string
	}
}

// Имплементация методов интерфейса
func (d *DataXLS) Title() []string {
	return []string{"Rendez-Vous Parser"}
}

func (d *DataXLS) TitleRows() []interface{} {
	return []interface{}{""}
}

func (d *DataXLS) Header() []string {
	return []string{"Название", "Цена"}
}

func (d *DataXLS) HeaderRows() [][]interface{} {
	var row [][]interface{}
	for _, value := range d.DataXLS {
		r := []interface{}{
			value.Name,
			value.Price,
		}
		row = append(row, r)
	}
	return row
}
