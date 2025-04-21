package main

import (
	"fmt"
	"go-hero/bank"
	"go-hero/logging"

	// "go-hero/geometry"
	// "go-hero/logging"
	// "go-hero/media"
	"sync"
	"time"
)

func PrintNumbers(from string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("From %s - %d\n", from, i)
		time.Sleep(time.Second)
	}
}

func main() {
	// rectangle := geometry.Rectangle{Length: 5, Breadth: 4}
	// circle := geometry.Circle{Radius: 14}

	// geometry.PrintShapeDetails(rectangle)
	// geometry.PrintShapeDetails(circle)

	// acc := bank.BankAccount{Owner: "Ravi", Balance: 1000}
	// acc.PrintBalance()
	// acc.Deposit(500)
	// acc.PrintBalance()
	// acc.Withdraw(200)
	// acc.PrintBalance() // should print: Current balance: 1300

	// playlistSongs := []media.Song{
	// 	{Title: "1", Artist: "A", DurationInSeconds: 300},
	// 	{Title: "2", Artist: "B", DurationInSeconds: 200},
	// 	{Title: "3", Artist: "C", DurationInSeconds: 250},
	// }

	// playlist := media.Playlist{Songs: playlistSongs}
	// playlist.AddSong(media.Song{Title: "4", Artist: "D", DurationInSeconds: 150})
	// fmt.Println(playlist.TotalDuration())
	// playlist.PrintSongs()

	// fileLogger, err := logging.NewFileLogger("testlog.log", true)
	// if err != nil {
	// 	fmt.Println("error opening the file")
	// } else {
	// 	loggers := []logging.Logger{
	// 		&logging.ConsoleLogger{BaseLogger: logging.BaseLogger{Utc: false}},
	// 		fileLogger,
	// 	}

	// 	for _, logger := range loggers {
	// 		logger.Info("Info Message")
	// 		logger.Error("Error Message", nil)
	// 		logger.Error("Error Message with Error", fmt.Errorf("error error error"))
	// 	}
	// }

	// go PrintNumbers("First")
	// go PrintNumbers("Second")

	// time.Sleep(6 * time.Second)
	// fmt.Println("main function done")

	account := &bank.BankAccount{Owner: "Rajeev", Balance: 1000, Logger: &logging.ConsoleLogger{}}

	var wg sync.WaitGroup

	// Launch 5 concurrent deposits
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			account.Deposit(100)
		}(i)
	}

	// Launch 5 concurrent withdrawals
	for i := 0; i < 75; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			account.Withdraw(50)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final Balance: $%.2f\n", account.GetBalance())
}
