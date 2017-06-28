package cinema

import "time"

//Movie struct
type Movie struct {
	Name      string      `json:"name"`
	Showtimes []time.Time `json:"showtimes"`
}
