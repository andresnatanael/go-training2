package ticket

import (
	"time"

	"github.com/andresnatanael/go-training2/cinema"
)

//FullPrice structure
type FullPrice struct {
	Movie     *cinema.Movie
	Showtime  time.Time
	PaidPrice float32
}

//GetMovie return the movie for the ticket
func (t *FullPrice) GetMovie() cinema.Movie {
	return *t.Movie
}

//GetShowTime return the dateTime of the movie
func (t *FullPrice) GetShowTime() time.Time {
	return t.Showtime
}

//GetCurrentPrice returns the current price for FullPrice Tickets
func (t *FullPrice) GetCurrentPrice() float32 {
	return RegularPrice
}

//GetPaidPrice returns the original price of the buy
func (t *FullPrice) GetPaidPrice() float32 {
	return t.PaidPrice
}

//GetType return the tiket type
func (t *FullPrice) GetType() string {
	return "full_price"
}

//GetDiscount returns the discount amount
func (t *FullPrice) GetDiscount() float32 {
	return 0
}
