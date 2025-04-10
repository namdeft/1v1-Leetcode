package main

type LemonadeStand struct {
	five int
	ten  int
}

func lemonadeChange(bills []int) bool {
	ls := &LemonadeStand{}
	return ls.processTransactions(bills)
}

func (ls *LemonadeStand) processTransactions(bills []int) bool {
	for _, val := range bills {
		if !ls.canExchange(val) {
			return false
		}
	}

	return true
}

func (ls *LemonadeStand) canExchange(bill int) bool {
	switch bill {
	case 5:
		ls.five++
	case 10:
		if ls.five < 1 {
			return false
		}
		ls.ten++
		ls.five--
	case 20:
		if !ls.processCash20() {
			return false
		}
	default:
		return false
	}

	return true
}

func (ls *LemonadeStand) processCash20() bool {
	if ls.ten >= 1 {
		if ls.five < 1 {
			return false
		}
		ls.ten--
		ls.five--
	} else {
		if ls.five < 3 {
			return false
		}
		ls.five = ls.five - 3
	}

	return true
}
