package main

import (
	"fmt"
	"go-hero/bank"
	"go-hero/geometry"
	"go-hero/media"
)

func main() {
	r := geometry.Rectangle{Length: 5, Breadth: 4}
	fmt.Println(r.Area())      // 20
	fmt.Println(r.Perimeter()) // 18

	acc := bank.BankAccount{Owner: "Ravi", Balance: 1000}
	acc.PrintBalance()
	acc.Deposit(500)
	acc.PrintBalance()
	acc.Withdraw(200)
	acc.PrintBalance() // should print: Current balance: 1300

	playlistSongs := []media.Song{
		{Title: "1", Artist: "A", DurationInSeconds: 300},
		{Title: "2", Artist: "B", DurationInSeconds: 200},
		{Title: "3", Artist: "C", DurationInSeconds: 250},
	}

	playlist := media.Playlist{Songs: playlistSongs}
	playlist.AddSong(media.Song{Title: "4", Artist: "D", DurationInSeconds: 150})
	fmt.Println(playlist.TotalDuration())
	playlist.PrintSongs()

}