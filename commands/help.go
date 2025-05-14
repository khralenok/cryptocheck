package commands

import "fmt"

func Help(isForCheck bool, isForTop bool) {
	if isForCheck {
		fmt.Println("Help for check:\nUsage: check -symbols=\"btc,usd\" -fiat=\"usd\"")
	} else if isForTop {
		fmt.Println("Help for top:\nUsage: top -amount=\"5\" -fiat=\"usd\"")
	} else {
		fmt.Println("Available commands:\n1. help(Optional: -check or -top)\n2. check(Required: -symbol, Optional: -fiat)\n3. top(Optional: -amount, -fiat)")
	}
}
