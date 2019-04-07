package stdpkgusage

import (
	"fmt"
	"testing"
)

func TestFmtPrint(t *testing.T) {
	fmt.Print("Hi,", "I'm", "fmt.Print(...).")
	fmt.Print("Default word delimiter is nothing. And no auto new line.\n")
	fmt.Println("Hi,", "I'm", "fmt.Println(...).")
	fmt.Println("Default word delimiter is space. And auto new line.")
	fmt.Printf("Hi, I'm %s.\nI have formatting ability.\n",
		"fmt.Printf(fmtstr, ...)")
}
