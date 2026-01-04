package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync" // 1. 동기화를 위한 패키지 추가
)

func main() {
	dirPtr := flag.String("dir", ".", "대상 디렉토리")
	prefixPtr := flag.String("prefix", "parallel", "접두어")
	flag.Parse()

	targetDir := *dirPtr
	prefix := *prefixPtr

	files, err := os.ReadDir(targetDir)
	if err != nil {
		fmt.Printf("에러: %v\n", err)
		return
	}

	// 2. WaitGroup 선언 (일꾼들의 출석부라고 생각하세요)
	var wg sync.WaitGroup

	for _, file := range files {
		if !file.IsDir() {
			// 3. 고루틴 하나가 추가될 때마다 대기 목록에 +1
			wg.Add(1)

			// 4. 익명 함수를 고루틴으로 실행
			go func(f os.DirEntry) {
				// 5. 함수가 끝나면(복사 완료) WaitGroup에 보고 (-1)
				defer wg.Done()
				// defer은 Python에서 try: finally:에 finally같은 역할, 고루틴 중간에 return으로 끝날 수도 있으니 미리 등록
				// 없으면 데드락 상태에 빠짐 있음
				oldPath := filepath.Join(targetDir, f.Name())
				newPath := filepath.Join(targetDir, prefix+"_"+f.Name())

				err := copyFile(oldPath, newPath)
				if err != nil {
					fmt.Printf("실패: %s\n", f.Name())
				} else {
					fmt.Printf("병렬 복사 완료: %s\n", newPath)
				}
			}(file) // 현재 순서의 file 정보를 고루틴에 전달
		}
	}

	// 6. 모든 일꾼(고루틴)이 보고(Done)를 마칠 때까지 여기서 대기
	fmt.Println("모든 고루틴이 작업을 마칠 때까지 기다립니다...")
	wg.Wait()
	fmt.Println("모든 작업이 완료되었습니다!")
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