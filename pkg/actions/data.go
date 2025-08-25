package actions

import (
	"Password-Keeper/pkg/etc"
	"fmt"
	"strings"
	"unicode/utf8"
)

type offset struct {
	number         string
	password       string
	description    string
	date           string
	tableLenght    int
	maxNumber      int
	maxPassword    int
	maxDescription int
}

func (act *Action) printAllData(data []etc.Data) {
	oft := NewOffset(data)

	fmt.Println(oft.fillSymbols(oft.tableLenght, "-"))
	fmt.Printf("| №%s | Название пароля%s | Описание%s | Дата создания%s |\n", oft.number, oft.password, oft.description, oft.date)
	fmt.Println(oft.fillSymbols(oft.tableLenght, "-"))
	for i, dt := range data {
		fmt.Printf("| %d%s | %s%s | %s%s | %s%s |\n",
			i+1, oft.fillSymbols(oft.maxNumber-oft.maxDigit(i+1), " "),
			dt.PasswordName, oft.fillSymbols(oft.maxPassword-utf8.RuneCountInString(dt.PasswordName), " "),
			dt.Description, oft.fillSymbols(oft.maxDescription-utf8.RuneCountInString(dt.Description), " "),
			dt.DateOfCreation, oft.fillSymbols(19-utf8.RuneCountInString(dt.DateOfCreation), " "),
		)
		fmt.Println(oft.fillSymbols(oft.tableLenght, "-"))
	}
	fmt.Println()
}

func (o *offset) maxDigit(number int) int {
	if number == 0 {
		return 1
	}

	var digitOfNumber int = 0

	for number > 0 {
		number /= 10
		digitOfNumber++
	}

	return digitOfNumber
}

func NewOffset(data []etc.Data) *offset {
	var oft offset

	oft.tableLenght = 13

	oft.maxNumber = oft.maxDigit(len(data))
	oft.tableLenght += oft.maxNumber
	oft.number = oft.fillSymbols(oft.maxNumber-1, " ")

	oft.maxPassword = 15
	for _, dt := range data {
		lenght := utf8.RuneCountInString(dt.PasswordName)
		if lenght > oft.maxPassword {
			oft.maxPassword = lenght
		}
	}
	oft.tableLenght += oft.maxPassword
	oft.password = oft.fillSymbols(oft.maxPassword-15, " ")

	oft.maxDescription = 8
	for _, dt := range data {
		lenght := utf8.RuneCountInString(dt.Description)
		if lenght > oft.maxDescription {
			oft.maxDescription = lenght
		}
	}
	oft.tableLenght += oft.maxDescription
	oft.description = oft.fillSymbols(oft.maxDescription-8, " ")

	oft.tableLenght += 19
	oft.date = oft.fillSymbols(6, " ")

	return &oft
}

func (o *offset) fillSymbols(offset int, symbol string) string {
	if offset < 1 {
		return ""
	}
	return strings.Repeat(symbol, offset)
}
