# 2 парсера интернет-магазинов через CLI-приложение с параметрами

# Тестовое задание от компании Albatro на языке Golang.

Перед началом нужно сделать go build и go install

Если после этого при попытке запустить команду parser выдает ошибку: command not found, сделать: 

export GOPATH=$(go env GOPATH)
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN

Приложение запускается командой parser с указанием нужного парсера. Их 2:

- rendez - парсер сайта интернет-магазина Rendez-Vous
- tsum - парсер сайта интернет-магазина ЦУМ

Далее нужно указать параметры. Их 3:

--sort - Выбор варианта сортировки. asc - алфавитный порядок, desc - обратный порядок, without - без сортировки.

--file - Выбор сохранения файла. csv либо xls

--name - Название файла.

Например:

parser rendez --sort asc --file csv --name rendezAsc - Парсинг Rendez-Vous с сортировкой по алфавиту и сохранением названия файла rendezAsc в формате .csv

parser tsum --sort desc --file xls --name tsumDesc - Парсинг ЦУМ с сортировкой в обратном порядке и сохранением названия файла tsumDesc в формате .xls

