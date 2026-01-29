// 로그 시스템
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 1. 미들웨어 정의: 핸들러를 받아서 핸들러를 반환하는 함수
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 다음 핸들러 실행 (진짜 비즈니스 로직)
		next.ServeHTTP(w, r)

		// 실행 완료 후 로그 출력
		log.Printf(
			"[%s] %s %s %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "인프라 모니터링 시스템에 접속하셨습니다.")
}

func main() {
	// 2. 일반 핸들러 생성
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)

	// 3. 미들웨어로 핸들러를 감싸기 (중첩 가능)
	wrappedMux := loggingMiddleware(mux)

	fmt.Println("서버가 8080 포트에서 시작되었습니다...")
	if err := http.ListenAndServe(":8080", wrappedMux); err != nil {
		log.Fatal(err)
	}
}