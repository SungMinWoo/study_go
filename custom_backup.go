package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// 1. 사용할 플래그 정의 (포인터 개념이 등장합니다!)
	// flag.String(이름, 기본값, 설명)
	dirPtr := flag.String("dir", ".", "복사할 대상 디렉토리 경로")
	prefixPtr := flag.String("prefix", "backup", "파일명 앞에 붙일 접두어")

	// 2. 파싱 (터미널에서 입력한 값을 읽어옵니다)
	flag.Parse()

	// 3. 포인터이므로 값을 꺼낼 때는 *를 붙입니다.
	targetDir := *dirPtr
	prefix := *prefixPtr

	fmt.Printf("설정된 경로: %s, 접두어: %s\n", targetDir, prefix)

	// 해당 디렉토리 읽기
	files, err := os.ReadDir(targetDir)
	if err != nil {
		fmt.Printf("에러: %v\n", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			oldPath := filepath.Join(targetDir, file.Name())
			newPath := filepath.Join(targetDir, prefix+"_"+file.Name())

			err := copyFile(oldPath, newPath)
			if err != nil {
				fmt.Printf("실패: %s\n", file.Name())
			} else {
				fmt.Printf(_)
				fmt.Printf("복사 완료: %s\n", newPath)
			}
		}
	}
}

// 파일 복사 함수 (이전과 동일)
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
