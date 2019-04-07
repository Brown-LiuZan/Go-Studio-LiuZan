package stdpkgusage

import (
	"flag"
	"fmt"
	"testing"
	"time"
)

func printFlag(f *flag.Flag) {
	fmt.Println("Name:", f.Name)
	fmt.Println("Usage:", f.Usage)
	fmt.Println("Value:", f.Value)
	fmt.Println("DefVal:", f.DefValue)
}

func TestFlag(t *testing.T) {
	var bool_flag bool
	var int_flag int
	var int64_flag int64
	var uint_flag uint
	var uint64_flag uint64
	var float64_flag float64
	var string_flag string
	var duration_flag time.Duration
	flag.BoolVar(&bool_flag, "bool-flag-name", false, "`bool` flag: usage message")
	flag.IntVar(&int_flag, "int-flag-name", -2, "`int` flag: usage message")
	flag.Int64Var(&int64_flag, "int64-flag-name", -22, "`int64` flag: usage message")
	flag.UintVar(&uint_flag, "uint-flag-name", 2, "`uint` flag: usage message")
	flag.Uint64Var(&uint64_flag, "uint64-flag-name", 22, "`uint64` flag: usage message")
	flag.Float64Var(&float64_flag, "float64-flag-name", 2.2, "`float64` flag: ussage message")
	flag.StringVar(&string_flag, "string-flag-name", "flag usage", "`string` flag: usage message")
	flag.DurationVar(&duration_flag, "duration-flag-name", 1*time.Second, "`time.Duration` flag: usage message")
	flag.Parse()

	if flag.Parsed() {
		fmt.Println("Command line flags parsed successfully.")
		fmt.Printf("Number of flags: %d\n", flag.NFlag())
		fmt.Printf("Number of non-flag arguments: %d\n", flag.NArg())
		fmt.Println("Name of top-level flagset:", flag.CommandLine.Name())
		fmt.Printf("Output of top-level flagset: %T\n", flag.CommandLine.Output())
		fmt.Println("ErrorHandling of top-level flagset:", flag.CommandLine.ErrorHandling())
		flag.Visit(printFlag)
		flag.VisitAll(printFlag)
	} else {
		panic("Command line flags unparsed yet.")
	}
}
