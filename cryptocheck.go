package main

import (
	"flag"
	"fmt"
	"os"

	"example.com/cryptocheck/commands"
)

func main() {
	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)
	helpCheck := helpCmd.Bool("check", false, "Show help for the 'check' command")
	helpTop := helpCmd.Bool("top", false, "Show help for the 'top' command")

	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)
	checkSymbols := checkCmd.String("symbols", "", "Comma-separated list of symbols")
	checkFiat := checkCmd.String("fiat", "USD", "Fiat currency")

	topCmd := flag.NewFlagSet("top", flag.ExitOnError)
	topAmount := topCmd.Int("amount", 10, "Amount of top currencies")
	topFiat := topCmd.String("fiat", "USD", "Fiat currency")

	switch os.Args[1] {
	case "help":
		helpCmd.Parse(os.Args[2:])
		commands.Help(*helpCheck, *helpTop)
	case "check":
		checkCmd.Parse(os.Args[2:])
		err := commands.Check(checkSymbols, checkFiat)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	case "top":
		topCmd.Parse(os.Args[2:])
		err := commands.Top(topAmount, topFiat)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	default:
		fmt.Println("expected some command")
		os.Exit(1)
	}
}
