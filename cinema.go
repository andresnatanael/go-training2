package main

import (
	"time"

	"github.com/andresnatanael/go-training2/cinema"
	"github.com/andresnatanael/go-training2/ticket"
)

func main() {

}

func createFullPriceTicket(movie *cinema.Movie, showtime time.Time) ticket.Ticket {
	return &ticket.FullPrice{Movie: movie, Showtime: showtime}
}
