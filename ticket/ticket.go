package ticket

import (
	"time"

	"github.com/andresnatanael/go-training2/cinema"
)

const (
	//RegularPrice represent the current price that will
	//be used for all kind of tickets
	RegularPrice    = float32(100)
	guestDiscount   = float32(100)
	retiredDiscount = float32(50)
)

//Ticket interface represent the base
//Operations for tickets
type Ticket interface {
	GetMovie() cinema.Movie
	GetShowTime() time.Time
	GetCurrentPrice() float32
	GetPaidPrice() float32
	SetPaidPrice(price float32)
	GetType() string
	GetDiscount() float32
}
