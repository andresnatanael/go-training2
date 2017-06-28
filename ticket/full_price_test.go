package ticket

import (
	"fmt"
	"testing"
	"time"

	"github.com/andresnatanael/go-training2/cinema"
)

func TestFullPrice(t *testing.T) {

	movie := &cinema.Movie{Name: "Armageddon"}

	showtime := time.Date(2017, 11, 20, 20, 10, 0, 0, time.Local)

	var ticket Ticket = &FullPrice{movie, showtime, 0}

	var ticket2 Ticket = &FullPrice{movie, showtime, 0}

	movies := []Ticket{ticket, ticket2}

	fmt.Println(movies)

	fmt.Println(ticket.GetMovie())
	if ticket.GetCurrentPrice() != float32(100) {

		t.Errorf("GOT: %f EXPECTED: %f", ticket.GetCurrentPrice(), float32(100))
	}

}
