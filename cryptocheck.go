package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)
	checkSymbols := checkCmd.String("symbols", "", "Comma-separated list of symbols")
	checkFiat:= checkCmd.String("fiat", "USD", "Fiat currency")

	topCmd := flag.NewFlagSet("top", flag.ExitOnError)
	topAmount := topCmd.Int("amount", 10, "Amount of top currencies")
	topFiat := topCmd.String("fiat", "USD", "Fiat currency")

	if len(os.Args) < 2{
		fmt.Println("provide a command please")
		os.Exit(1)
	}

	switch os.Args[1]{
	case "check":
		checkCmd.Parse(os.Args[2:])
		fmt.Println("Check command")
		fmt.Println("Symbols: ", *checkSymbols)
		fmt.Println("Fiat: ", *checkFiat)
		os.Exit(0)
	case "top":
		topCmd.Parse(os.Args[2:])
		fmt.Println("Top command")
		fmt.Println("Amount: ", *topAmount)
		fmt.Println("Fiat: ", *topFiat)
		os.Exit(0)
	default:
		fmt.Println("expected some command")
		os.Exit(1)
	}
}