package ticket

import (
	"time"

	"github.com/andresnatanael/go-training2/cinema"
)

const (
	regularPrice    = float32(100)
	guestDiscount   = float32(100)
	retiredDiscount = float32(50)
)

//Ticket interface represent the base
//Operations for tickets
type Ticket interface {
	GetMovie() cinema.Movie
	GetShowTime() time.Time
	GetPrice() float32
	GetType() string
	GetDiscount() float32
}
