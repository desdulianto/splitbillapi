package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/desdulianto/splitbill"
)

func parsePeople(peopleFlag *string) splitbill.People {
	people := make(splitbill.People, 0)

	for _, p := range strings.Split(*peopleFlag, ",") {
		if p != "" {
			person := splitbill.Person(p)
			people = append(people, person)
		}
	}

	return people
}

func parseOptions() splitbill.Bill {
	amount := flag.Int("amount", 0, "Bill amount (required).")
	paidBy := flag.String("paid-by", "", "Who paid the bill (required).")
	peopleFlag := flag.String("people", "",
		"List of people in group, use ',' to separate (required).")

	flag.Parse()

	return splitbill.Bill{
		Amount: splitbill.Money(*amount),
		PaidBy: splitbill.Person(*paidBy),
		People: parsePeople(peopleFlag),
	}
}

func main() {
	bill := parseOptions()

	// output
	amount, err := bill.SplitEvenly()
	payTo := bill.PaidBy
	payFrom := bill.GetPeople()

	if err == nil {
		fmt.Printf("Bill Amount Total: %v\n", bill.Amount)
		fmt.Printf("Bill paid by: %v\n", bill.PaidBy)
		fmt.Printf("People: %v\n", payFrom)
		for _, p := range payFrom {
			fmt.Printf("%v need to pay %v to %v\n", p, amount, payTo)
		}
	} else {
		fmt.Println("Error: ", err.Error())
	}
}
