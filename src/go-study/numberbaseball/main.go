package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	strikes int
	balls   int
}

// 1
// 랜덤 숫자 3개 뽑기
func MakeNumbers() [3]int {
	var rst [3]int
	for i := 0; i < 3; i++ {
		for {
			n := rand.Intn(10)
			duplicated := false
			// 뽑은 수가 기존 수와 중복되는 지 확인
			for j := 0; j < i; j++ {
				if rst[j] == n {
					duplicated = true
					break
				}
			}
			// 중복되지 않으면 배열에 추가
			if !duplicated {
				rst[i] = n
				break
			}
		}
	}
	return rst
}

// 2
// 사용자 입력
func InputNumbers() [3]int {
	var rst [3]int

	for {
		fmt.Println("겹치지 않는 0~9 사이의 숫자 3개를 입력하세요")
		var no int

		// enter까지 포함 -> error 방지
		_, err := fmt.Scanf("%d\n", &no)

		// error가 null이 아니면(error가 있다)
		if err != nil {
			fmt.Println("잘못 입력하셨습니다")
			continue
		}

		success := true
		idx := 0
		// 숫자 3개 한꺼번에 입력 받기
		// 예외 처리
		for no > 0 {
			// 세 자리 수 하나씩 뽑기
			n := no % 10
			no = no / 10

			duplicated := false

			for j := 0; j < idx; j++ {
				if rst[j] == n {
					duplicated = true
					break
				}
			}

			// 숫자 중복
			if duplicated {
				fmt.Println("숫자가 겹치지 않아야 합니다")
				success = false
				break
			}

			// 입력 수 초과
			if idx >= 3 {
				fmt.Println("3개보다 많은 숫자를 입력하셨습니다")
				success = false
				break
			}

			// 예외 통과 시, 배열에 담기
			rst[idx] = n
			idx++
		}

		if success && idx < 3 {
			fmt.Println("3개의 숫자를 입력하세요")
			success = false
		}

		if !success {
			continue
		}

		break
	}

	// 사용자 입력을 1의 자리부터 가져오므로 역순으로 정렬
	rst[0], rst[2] = rst[2], rst[0]

	return rst
}

// 3
// 수 비교
func CompareNumbers(numbers, inputNumbers [3]int) Result {
	s := 0
	b := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// 두 수가 같을 때
			if numbers[i] == inputNumbers[j] {
				// 위치도 같으면
				if i == j {
					s++
				} else { // 수만 같으면
					b++
				}
				break
			}
		}
	}

	return Result{s, b} // 구조체 타입 반환
}

// 4
// 결과 출력
func PrintResult(result Result) {
	fmt.Printf("Your Result: %dS%dB\n", result.strikes, result.balls)
}

// 5
// 게임 종료 조건
func IsGameEnd(result Result) bool {
	return result.strikes == 3
}

func main() {
	// random seed 초기화
	// 항상 변하는 값 - 시간
	rand.Seed(time.Now().UnixNano())

	numbers := MakeNumbers()

	// 게임 시도 횟수
	cnt := 0

	for {
		cnt++

		inputNumbers := InputNumbers()

		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(result)

		if IsGameEnd(result) {
			break
		}
	}

	fmt.Printf("%d번만에 성공!", cnt)

}
