package main

import (
	"time"

	"github.com/andresnatanael/go-training2/cinema"
	"github.com/andresnatanael/go-training2/ticket"
)

func main() {

}

func createFullPriceTicket(movie *cinema.Movie, showtime time.Time) ticket.Ticket {
	var t ticket.Ticket = &ticket.Retired{Movie: movie, Showtime: showtime}
	return t
}

func createGuestTicket(movie *cinema.Movie, showtime time.Time) ticket.Ticket {
	return &ticket.Guest{Movie: movie, Showtime: showtime}
}

func createRetiredTicket(movie *cinema.Movie, showtime time.Time) ticket.Ticket {
	return &ticket.FullPrice{Movie: movie, Showtime: showtime, PaidPrice: 0}
}
