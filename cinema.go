package main

import (
	"errors"
	"time"

	"github.com/andresnatanael/go-training2/cinema"
	"github.com/andresnatanael/go-training2/ticket"
	"gopkg.in/gin-gonic/gin.v1"
)

var tickets []ticket.Ticket = make([]ticket.Ticket, 0)
var movies [2]cinema.Movie

type ticketMsg struct {
	Movie    string `json:"movie"`
	Showtime string `json:"showtime"`
}

func main() {
	initData()

	r := gin.Default()

	r.GET("/movies", func(c *gin.Context) {
		c.JSON(200, movies)
	})

	r.GET("/tickets", func(c *gin.Context) {
		c.JSON(200, tickets)
	})

	r.GET("/cinema/profits", func(c *gin.Context) {
		profit := float32(0)
		for _, v := range tickets {
			profit = profit + v.GetPaidPrice()
		}
		c.JSON(200, profit)
	})

	r.POST("/buy_ticket/full_price", func(c *gin.Context) {
		var message ticketMsg
		if err := c.BindJSON(&message); err != nil {
			c.JSON(500, err)
		}
		movie, t, parseError := parseData(message)

		if parseError != nil {
			c.JSON(500, parseError.Error())
		} else {
			var ticket ticket.Ticket = createFullPriceTicket(movie, t)
			tickets = append(tickets, ticket)

			c.JSON(200, ticket)
		}

	})

	r.POST("/buy_ticket/guest", func(c *gin.Context) {

		var message ticketMsg
		if err := c.BindJSON(&message); err != nil {
			c.JSON(500, err)
		}
		movie, t, parseError := parseData(message)

		if parseError != nil {
			c.JSON(500, parseError.Error())
		} else {
			var ticket ticket.Ticket = createGuestTicket(movie, t)
			tickets = append(tickets, ticket)

			c.JSON(200, ticket)
		}
	})

	r.POST("/buy_ticket/retired", func(c *gin.Context) {

		var message ticketMsg
		if err := c.BindJSON(&message); err != nil {
			c.JSON(500, err)
		}
		movie, t, parseError := parseData(message)

		if parseError != nil {
			c.JSON(500, parseError.Error())
		} else {
			var ticket ticket.Ticket = createRetiredTicket(movie, t)
			tickets = append(tickets, ticket)

			c.JSON(200, ticket)
		}
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func findMovie(movie string) *cinema.Movie {
	for _, v := range movies {
		if v.Name == movie {
			return &v
		}
	}
	return nil
}

func parseData(message ticketMsg) (*cinema.Movie, time.Time, error) {
	t, err := time.Parse(time.RFC3339, message.Showtime)

	if err != nil {
		return &cinema.Movie{}, time.Now(), err
	}
	movie := findMovie(message.Movie)

	if movie == nil {
		return &cinema.Movie{}, time.Now(), errors.New(message.Movie + " not found!")
	}

	return movie, t, nil
}

func initData() {
	var showtimes = make([]time.Time, 0)
	showtimes = append(showtimes, time.Date(2017, 11, 20, 20, 10, 0, 0, time.Local))
	showtimes = append(showtimes, time.Date(2017, 11, 21, 20, 10, 0, 0, time.Local))

	movies[0] = cinema.Movie{Name: "Matrix", Showtimes: showtimes}
	movies[1] = cinema.Movie{Name: "Armageddon", Showtimes: showtimes}
}

func createFullPriceTicket(movie *cinema.Movie, showtime time.Time) ticket.Ticket {
	var t ticket.Ticket = &ticket.FullPrice{Movie: movie, Showtime: showtime}
	t.SetPaidPrice(t.GetCurrentPrice())
	return t
}

func createGuestTicket(movie *cinema.Movie, showtime time.Time) ticket.Ticket {
	var t ticket.Ticket = &ticket.Guest{Movie: movie, Showtime: showtime}
	t.SetPaidPrice(t.GetCurrentPrice())
	return t
}

func createRetiredTicket(movie *cinema.Movie, showtime time.Time) ticket.Ticket {
	var t ticket.Ticket = &ticket.Retired{Movie: movie, Showtime: showtime}
	t.SetPaidPrice(t.GetCurrentPrice())
	return t
}
