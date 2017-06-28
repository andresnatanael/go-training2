package ticket

import (
	"time"

	"github.com/andresnatanael/go-training2/cinema"
)

//Guest structure
type Guest struct {
	Movie     *cinema.Movie
	Showtime  time.Time
	PaidPrice float32
}

//GetMovie return the movie for the ticket
func (t *Guest) GetMovie() cinema.Movie {
	return *t.Movie
}

//GetShowTime return the dateTime of the movie
func (t *Guest) GetShowTime() time.Time {
	return t.Showtime
}

//GetCurrentPrice returns the actual price for guests
func (t *Guest) GetCurrentPrice() float32 {
	return RegularPrice - ((RegularPrice * guestDiscount) / 100)
}

//GetPaidPrice returns the original price of the buy
func (t *Guest) GetPaidPrice() float32 {
	return t.PaidPrice
}

//SetPaidPrice records the paid price for the ticket
func (t *Guest) SetPaidPrice(price float32) {
	t.PaidPrice = price
}

//GetType return the tiket type
func (t *Guest) GetType() string {
	return "guest"
}

//GetDiscount returns the discount amount
func (t *Guest) GetDiscount() float32 {
	return guestDiscount
}
