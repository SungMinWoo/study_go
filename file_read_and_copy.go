package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var targetFolder string = "test"
var resultFolder string = "new"

func main() {
	// 1. 현재 디렉토리의 파일 읽기 (Python의 os.listdir()과 유사)
	files, err := os.ReadDir(fmt.Sprintf("./%s", targetFolder))
	if err != nil {
		fmt.Println("디렉토리를 읽을 수 없습니다:", err)
		return
	}

	// 2. 오늘 날짜 구하기 (YYYYMMDD 형식)
	today := time.Now().Format("20060102")

	fmt.Printf("오늘 날짜(%s)를 붙여 복사를 시작합니다...\n", today)

	for _, file := range files {
		// 디렉토리는 제외하고 파일만 처리

		if !file.IsDir() {
			originalName := file.Name()
			fmt.Printf("현재 처리 중인 파일: %s\n", originalName)
			newName := fmt.Sprintf("%s_%s", today, originalName)

			// 3. 파일 복사 실행
			err := copyFile(originalName, newName)
			if err != nil {
				fmt.Printf("실패: %s -> %s (사유: %v)\n", originalName, newName, err)
			} else {
				fmt.Printf("성공: %s -> %s\n", originalName, newName)
			}
		}
	}
}

// 파일 복사를 담당하는 함수
func copyFile(src, dst string) error {
	// 원본 파일 열기
	sourceFile, err := os.Open(fmt.Sprintf("./%s/%s", targetFolder, src))
	if err != nil {
		return err
	}
	defer sourceFile.Close() // 함수가 끝날 때 파일을 닫음 (Python의 with 문과 유사)

	if _, err := os.Stat(resultFolder); os.IsNotExist(err) {
		err := os.Mkdir(resultFolder, 0700) // 뒤에 0700은 mode 변수를 넣어야하는데 mode는 권한을 뜻함
		if err != nil {
			fmt.Printf("파일 생성 실패")
		}
	}

	// 대상 파일 생성
	destFile, err := os.Create(fmt.Sprintf("./%s/%s", resultFolder, dst))
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 내용 복사
	_, err = io.Copy(destFile, sourceFile)
	return err
}
