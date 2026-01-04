package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 일감 정의
type Job int

func worker(id int, jobs <-chan Job, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("일꾼 %d: 작업 %d 시작\n", id, j)
		time.Sleep(time.Millisecond * 500) // 작업 시뮬레이션
		results <- int(j * 2)              // j는 Job타입이라 변환해주어야함.
	}
}

func main() {
	// 1. CPU 개수에 맞게 워커 수 설정
	// runtime.NumCPU()는 현재 머신의 논리적 CPU 코어 수를 반환합니다.
	numWorkers := runtime.NumCPU()
	fmt.Printf("시스템 CPU 코어 수: %d\n", numWorkers)

	jobs := make(chan Job, 100)
	results := make(chan int, 100)
	var wg sync.WaitGroup

	// 2. 워커 풀 가동 (CPU 개수만큼 고루틴 생성)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 3. 일감 투입
	for j := 1; j <= 20; j++ {
		fmt.Printf("일감 %d. \n", j)
		jobs <- Job(j)
	}
	close(jobs) // 일감이 끝났음을 알림

	// 4. 모든 작업 완료 대기
	wg.Wait()
	close(results)

	fmt.Println("모든 작업 완료!")
}
// 동작 순서
// 1. 고루틴을 워커 개수에 맞게 생성한다
// 2. 고루틴은 Job(일감)이 들어올 때까지 대기한다.
// 3. 3번 일감 투입에서 반복문이 돌면서 일감을 추가한다.
// 4. 3번 일감 투입 반복문이 돌아감과 동시에 고루틴은 일감이 들어왔으니 일을 시작한다

// 잼미니 부연설명
// 1. 고루틴 생성 (일꾼 배치)
// go worker(...)가 실행되는 순간, 메모리 상에 아주 가벼운 일꾼(고루틴)들이 배치됩니다. 이들은 생성되자마자 자기 코드를 실행하기 시작합니다.

// 2. Job 대기 (채널의 Blocking 특징)
// 워커 내부의 for j := range jobs 코드가 핵심입니다. 채널에 데이터가 없으면 고루틴은 CPU 자원을 거의 쓰지 않고 '대기(Blocking)' 상태로 멈춰있습니다. 마치 식당 주방장들이 주문서가 들어오길 기다리며 대기하는 것과 같습니다.

// 3. 일감 투입 및 즉시 시작
// 메인 루프에서 jobs <- Job(j)를 하는 순간, 대기 중이던 워커 중 하나가 즉시 깨어나서 일감을 가져갑니다.

// 말씀하신 대로 메인 루프가 다음 일감(4번, 5번...)을 채널에 넣는 동안, 이미 일감을 가져간 워커들은 동시에(Parallel) 각자의 업무를 처리합니다.