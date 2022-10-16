package main

import "time"

func sendingString(channel chan<- string, s string) {
	channel <- s
}

func receiver(helloCh, goodbyeCh <-chan string, quitCh chan bool) {
	for {
		select {
		case s := <-helloCh:
			println(s)
		case s := <-goodbyeCh:
			println(s)
		case <-time.After(time.Second * 2):
			println("Timeout")
			quitCh <- true
			break
		}
	}
}

func main() {
	helloCh := make(chan string, 1)
	goodbyeCh := make(chan string, 1)
	quitCh := make(chan bool)

	go receiver(helloCh, goodbyeCh, quitCh)
	go sendingString(helloCh, "Hello World!")
	time.Sleep(time.Second * 1)
	go sendingString(goodbyeCh, "Goodbye World!")
	<-quitCh
}
