package ticket

import (
	"time"

	"github.com/andresnatanael/go-training2/cinema"
)

//Retired structure
type Retired struct {
	movie    *cinema.Movie
	showtime time.Time
}

//GetMovie return the movie for the ticket
func (t *Retired) GetMovie() cinema.Movie {
	return *t.movie
}

//GetShowTime return the dateTime of the movie
func (t *Retired) GetShowTime() time.Time {
	return t.showtime
}

//GetPrice returns the price for FullPrice Tickets
func (t *Retired) GetPrice() float32 {
	return regularPrice - ((regularPrice * retiredDiscount) / 100)
}

//GetType return the tiket type
func (t *Retired) GetType() string {
	return "guest"
}

//GetDiscount returns the discount amount
func (t *Retired) GetDiscount() float32 {
	return retiredDiscount
}
