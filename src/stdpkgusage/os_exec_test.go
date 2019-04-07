package stdpkgusage

import (
	"fmt"
	"log"
	"os"
	osexec "os/exec"
	"testing"
)

func TestOsexecDummy(t *testing.T) {
	log.Println("Package log implements a simple logging package. It defines a type, Logger, with methods for formatting output. It also has a predefined 'standard' Logger accessible through helper functions Print[f|ln], Fatal[f|ln], and Panic[f|ln], which are easier to use than creating a Logger manually. That logger writes to standard error and prints the date and time of each logged message. Every log message is output on a separate line: if the message being printed does not end in a newline, the logger will add one. The Fatal functions call os.Exit(1) after writing the log message. The Panic functions call panic after writing the log message.")
	fmt.Println("Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.")
	cwd, _ := os.Getwd()
	fmt.Printf("Current working directory: %s\n", cwd)
	cmd := osexec.Command("bash", "-c", "echo", "'hi'")
	cmd.Run()
}
