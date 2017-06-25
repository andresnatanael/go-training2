package ticket

import (
	"time"

	"github.com/andresnatanael/go-training2/cinema"
)

//FullPrice structure
type FullPrice struct {
	Movie    *cinema.Movie
	Showtime time.Time
}

//GetMovie return the movie for the ticket
func (t *FullPrice) GetMovie() cinema.Movie {
	return *t.Movie
}

//GetShowTime return the dateTime of the movie
func (t *FullPrice) GetShowTime() time.Time {
	return t.Showtime
}

//GetPrice returns the price for FullPrice Tickets
func (t *FullPrice) GetPrice() float32 {
	return regularPrice
}

//GetType return the tiket type
func (t *FullPrice) GetType() string {
	return "full_price"
}

//GetDiscount returns the discount amount
func (t *FullPrice) GetDiscount() float32 {
	return 0
}
