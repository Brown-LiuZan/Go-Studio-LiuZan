package stdpkgusage

import (
	"log"
	"os"
	"testing"
)

var myLogger = log.New(os.Stderr, "[MyLogger]",
	log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile|log.LUTC)

func TestLogPrints(t *testing.T) {
	log.Print("Hi,", "I'm", "log.Print(...).",
		"Word delimiter is nothing.")
	log.Println("Hi,", "I'm", "log.Println(...).",
		"Word delimiter is space.")
	log.Printf("Hi, I'm %s.", "log.Printf(fmtstr, ...)")
	log.Println("Always auto new line if no one there.",
		"Always prepend date and time automatically.",
		"This is different from fmt package.")
	myLogger.Print("Hi,", " I'm myLogger with log.Ldate|log.Ltime|",
		"log.Lmicroseconds|log.Llongfile|log.LUTC set.")
}

/*
func TestLogFatal(t *testing.T) {
	log.Fatal("Hi,", "I'm", "log.Fatal(...).",
		"Word delimiter is nothing.",
		"Will Exit(1)")
}

func TestLogFatalln(t *testing.T) {
	log.Fatalln("Hi,", "I'm", "log.Fatal(...).",
		"Word delimiter is whitespace.",
		"Will Exit(1)")
}

func TestLogFatalf(t *testing.T) {
	log.Fatalf("Hi, I'm %s. I have formatting ability. Will Exit(1).",
		"log.Fatal(...)")
}
*/

func TestLogPanic(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			log.Println("Error:", p)
		}
	}()

	log.Panic("Hi,", "I'm", "log.Panic(...).",
		"Word delimiter is nothing.",
		"Will panic() by defualt.")
}

func TestLogPanicln(t *testing.T) {
	defer func() {
		if pe := recover(); pe != nil {
			log.Println("pe := recover() invoked. Error:", pe)
		}
	}()

	log.Panicln("Hi,", "I'm", "log.Panicln(...).",
		"Word delimiter is whitespace.",
		"Will panic() by defualt.")
}

func TestLogPanicf(t *testing.T) {
	defer func() {
		if pe := recover(); pe != nil {
			log.Println("pe := recover() invoked. Error:", pe)
		}
	}()

	log.Panicf("Hi, I'm %s. Will panic() by defualt.",
		"log.Panicf(fmtstr, ...)")
}
