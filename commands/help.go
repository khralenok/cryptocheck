package commands

import "fmt"

func Help(isForCheck bool, isForTop bool) {
	if isForCheck {
		fmt.Println("Help for check:\nUsage: check symbols=\"btc,usd\" fiat=\"usd\"")
	} else if isForTop {
		fmt.Println("Help for top:\nUsage: top amount=\"5\" fiat=\"usd\"")
	} else {
		fmt.Println("Available commands:\n1. Help\n2. Check\n3. Top")
	}

}
