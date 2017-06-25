package ticket

import (
	"time"

	"github.com/andresnatanael/go-training2/cinema"
)

//Guest structure
type Guest struct {
	movie    *cinema.Movie
	showtime time.Time
}

//GetMovie return the movie for the ticket
func (t *Guest) GetMovie() cinema.Movie {
	return *t.movie
}

//GetShowTime return the dateTime of the movie
func (t *Guest) GetShowTime() time.Time {
	return t.showtime
}

//GetPrice returns the price for FullPrice Tickets
func (t *Guest) GetPrice() float32 {
	return regularPrice - ((regularPrice * guestDiscount) / 100)
}

//GetType return the tiket type
func (t *Guest) GetType() string {
	return "guest"
}

//GetDiscount returns the discount amount
func (t *Guest) GetDiscount() float32 {
	return guestDiscount
}
