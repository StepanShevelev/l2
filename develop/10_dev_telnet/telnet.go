package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// парсим таймаут, либо задаем стандарные 10 сек
	timeout := flag.Int("timeout", 10, "timeout for connection")
	flag.Parse()
	// проверка на количество аргументов
	if flag.NArg() != 2 {
		log.Fatal("incorrect args")
	}
	// складываем адресс из ip и port
	args := flag.Args()
	address := args[0] + ":" + args[1]
	// подключаемся по tcp к адрессу, и таймаутом на подключение
	conn, err := net.DialTimeout("tcp", address, time.Duration(*timeout)*time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	// настраиваем контекст для завершения программы
	ctx, cancel := context.WithCancel(context.Background())
	// канал на прием сигнала из ОС
	signalChan := make(chan os.Signal, 1)
	// запись в канал сигнала из ОС
	signal.Notify(signalChan, syscall.SIGINT)
	// горутина на ожидание сигнала из ОС
	go func() {
		<-signalChan
		cancel()
	}()
	// запуск в горутине функции общения с сервером
	go Connect(ctx, conn)
	// ожидание завершения контекста
	<-ctx.Done()
	fmt.Println("\nConnection closed")
}

func Connect(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	// читаем из conn, записываем в stdout
	for {
		select {
		case <-ctx.Done():
			return
		default:

			// Чтение входных данных от stdin
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Text to send: ")
			text, _ := reader.ReadString('\n')
			// Отправляем в socket
			fmt.Fprintf(conn, text+"\n")
			// Прослушиваем ответ
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Message from server: " + message)

		}

	}
}
