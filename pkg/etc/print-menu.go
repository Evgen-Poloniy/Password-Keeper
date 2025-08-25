package etc

import "fmt"

func PrintMenu(menu []string) {
	fmt.Println("Выберите действие:")
	fmt.Println()

	menu[0] = fmt.Sprintf("Текущая учетная запись: \"%s\"", Settings.CurrentUsername)

	for _, action := range menu {
		fmt.Println(action)
	}
}
