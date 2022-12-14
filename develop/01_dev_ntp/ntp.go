package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	timeExact, _ := exactTime("0.beevik-ntp.pool.ntp.org")
	fmt.Println("Exact time:", timeExact)
}

// функция определения точного времени
func exactTime(host string) (time.Time, error) {
	response, err := ntp.Query(host)
	if err != nil {
		log.Fatalln(err)
		// выход из программы с кодом 1
	}
	time := time.Now().Add(response.ClockOffset)

	return time, nil
}
