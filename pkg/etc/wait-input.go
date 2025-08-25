package etc

import (
	"fmt"
	"os"
	"os/exec"
)

func WaitInput() {
	fmt.Println("Для продолжения нажмите Enter...")
	fmt.Scanln()
	ClearConsole()
}

func ClearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
