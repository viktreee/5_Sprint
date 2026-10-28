package actioninfo

import (
	"fmt"
	"log"
)

// DataParser объединяет разбор данных и формирование отчета об активности.
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// Info выводит отчеты об активности, логирует ошибки и пропускает неверные записи.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Println(err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(info)
	}
}
