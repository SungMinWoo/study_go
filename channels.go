// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sendData(ch chan string) {
// 	// 채널을 통해 데이터를 전송
// 	ch <- "Hello"
// 	ch <- "World"
// 	ch <- "test"
// }

// func receiveData(ch chan string) {
// 	// 채널을 통해 데이터를 수신
// 	msg1 := <-ch
// 	msg2 := <-ch
// 	msg3 := <-ch

// 	fmt.Println("Received messages:", msg1, msg2, msg3)
// }

// func main() {
// 	// 문자열을 전송하는 채널 생성
// 	messageChannel := make(chan string, 1)
	
// 	// 고루틴으로 데이터를 전송하는 함수 실행
// 	go sendData(messageChannel)

// 	// 고루틴으로 데이터를 수신하는 함수 실행
// 	go receiveData(messageChannel)

// 	// 메인 고루틴이 일정 시간 동안 대기
// 	time.Sleep(1 * time.Second)
// }

package main

import (
	"fmt"
	"time"
)

func main() {
	messageChannel := make(chan string, 1) // 크기가 3인 채널 생성

	go func() {
		messageChannel <- "Message 1"
		messageChannel <- "Message 2"
		messageChannel <- "Message 3"
		close(messageChannel)
	}()

	time.Sleep(1 * time.Second)

	// 채널로부터 데이터 받아오기
	for msg := range messageChannel {
		fmt.Println("Received:", msg)
	}

	fmt.Println("Received:", <-messageChannel)
}
