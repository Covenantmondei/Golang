package main

import "fmt"

type bill struct {
	name string
	items map[string]float64
	tip float64
}

// Funvtion to create a new bill
func newBill(name string) bill {
	b := bill {
		name: name,
		items: map[string]float64{"pie": 5.99, "cake": 10.99},
		tip: 0,
	}
	return b
}


func (b *bill) format() string {
	// Formatted string
	fs := "Bill Breakdown \n"
	var total float64 = 0

	// list items
	for k, v := range b.items{
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}
	//tip
	fs += fmt.Sprintf("%-25v ...$%v \n", "tip:", b.tip)
	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}


// Function to add a new item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}