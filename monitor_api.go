// 서버 상태 모니터링 API
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

// 1. JSON으로 응답할 데이터 구조체 정의 (구조체 태그 사용)
type StatusResponse struct {
	NumCPU       int    `json:"cpu_cores"`
	NumGoroutine int    `json:"active_goroutines"`
	MemoryAlloc  uint64 `json:"memory_alloc_mb"`
	OS           string `json:"os"`
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	// 2. 시스템 정보 수집
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	res := StatusResponse{
		NumCPU:       runtime.NumCPU(),
		NumGoroutine: runtime.NumGoroutine(), // 현재 떠있는 고루틴 개수!
		MemoryAlloc:  m.Alloc / 1024 / 1024,  // MB 단위 변환
		OS:           runtime.GOOS,
	}

	// 3. 응답 헤더 설정 (JSON 형태임을 명시)
	w.Header().Set("Content-Type", "application/json")

	// 4. 구조체를 JSON으로 변환하여 응답 스트림으로 바로 쏨
	json.NewEncoder(w).Encode(res)
}

func main() {
	// 5. 경로(Path)와 핸들러 함수 연결
	http.HandleFunc("/status", statusHandler)

	fmt.Println("서버가 8080 포트에서 시작되었습니다. http://localhost:8080/status 에 접속하세요.")

	// 6. 서버 시작 (에러 발생 시 출력)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("서버 시작 실패: %v\n", err)
	}
}
