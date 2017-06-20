package cinema

import (
	"github.com/andresnatanael/go-training2/customer"
)

func sellTicket(c customer.Customer, tickets *[]customer.Customer) {
	*tickets = append(*tickets, c)
}
