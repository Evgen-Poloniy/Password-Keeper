package etc

var (
	RedColor    string = "\033[31m"
	GreenColor  string = "\033[32m"
	YellowColor string = "\033[33m"
	BlueColor   string = "\033[34m"
	WhiteColor  string = "\033[0m"
)

func ResetColors() {
	RedColor = ""
	GreenColor = ""
	YellowColor = ""
	BlueColor = ""
	WhiteColor = ""
}
