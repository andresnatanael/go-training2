package ticket

import (
	"time"

	"github.com/andresnatanael/go-training2/cinema"
)

//Retired structure
type Retired struct {
	Movie     *cinema.Movie
	Showtime  time.Time
	PaidPrice float32
}

//GetMovie return the movie for the ticket
func (t *Retired) GetMovie() cinema.Movie {
	return *t.Movie
}

//GetShowTime return the dateTime of the movie
func (t *Retired) GetShowTime() time.Time {
	return t.Showtime
}

//GetCurrentPrice returns the actual price for retired
func (t *Retired) GetCurrentPrice() float32 {
	return RegularPrice - ((RegularPrice * retiredDiscount) / 100)
}

//GetPaidPrice returns the original price of the buy
func (t *Retired) GetPaidPrice() float32 {
	return t.PaidPrice
}

//SetPaidPrice records the paid price for the ticket
func (t *Retired) SetPaidPrice(price float32) {
	t.PaidPrice = price
}

//GetType return the tiket type
func (t *Retired) GetType() string {
	return "guest"
}

//GetDiscount returns the discount amount
func (t *Retired) GetDiscount() float32 {
	return retiredDiscount
}
