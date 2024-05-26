package main

import (
	"fmt"

	"github.com/porlizm/cinema/movie"
	"github.com/porlizm/cinema/ticket"
)

func main() {
	movieName := movie.FindName("tt4154796")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4)
	ticket.BuyTicket(movieName)

}
