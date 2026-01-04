package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// 1. 작업 결과를 담을 구조체 정의
type CopyResult struct {
	FileName string
	Success  bool
	Err      error
}

func main() {
	targetDir := "./new" // 테스트용 디렉토리
	prefix := "result"

	files, _ := os.ReadDir(targetDir)
	
	// 2. 결과를 전달받을 채널 생성 (결과 타입의 파이프)
	// make(chan 타입, 버퍼크기)
	resultChan := make(chan CopyResult, len(files))

	for _, file := range files {
		if !file.IsDir() {
			// 고루틴 실행
			go func(f os.DirEntry) {
				oldPath := filepath.Join(targetDir, f.Name())
				newPath := filepath.Join(targetDir, prefix+"_"+f.Name())

				err := copyFile(oldPath, newPath)
				// 3. 채널에 결과 전송 (파이프에 데이터 던지기)
				// 채널을 사용하면 자동으로 고루틴 작업을 기다림, 채널에 길이를 정해줘서 다 채워질 떄까지 blocking(기다리기) 해줌
				resultChan <- CopyResult{
					FileName: f.Name(),
					Success:  err == nil,
					Err:      err,
				}
			}(file)
		}
	}

	// 4. 메인 함수에서 채널을 통해 결과 수집
	successCount := 0
	failureCount := 0

	fmt.Println("작업 리포트 수집 중...")
	for i := 0; i < len(files); i++ {
		// 채널에서 데이터가 나올 때까지 기다림 (Blocking)
		result := <-resultChan 
		if result.Success {
			successCount++
		} else {
			failureCount++
			fmt.Printf("[실패] %s: %v\n", result.FileName, result.Err)
		}
	}

	fmt.Printf("\n--- 최종 결과 ---\n성공: %d, 실패: %d\n", successCount, failureCount)
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}