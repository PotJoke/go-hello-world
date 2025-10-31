package tools

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func WaitForEnter() {
	fmt.Print("\nPress Enter...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}

func Input(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input[:len(input)-1]
}

var Version = "v1.0"
var Author = "potjoke"
