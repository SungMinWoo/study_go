package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond) // 잠시 대기
		fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for char := 'a'; char <= 'e'; char++ {
		time.Sleep(100 * time.Millisecond) // 잠시 대기
		fmt.Printf("%c ", char)
	}
}

func main() {
	// 고루틴을 사용하여 함수들을 동시에 실행
	go printNumbers()
	go printLetters()

	// 메인 고루틴은 일정 시간 동안 대기
	time.Sleep(1 * time.Second)

	fmt.Println("\n메인 고루틴 종료")
}
