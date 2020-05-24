package main

import (
	"fmt"
	"os"
	"time"
)

type accessChecker struct {
	logArr []string
}

func (checker *accessChecker) newLog(message string) {

	// timestamp 추가
	now := time.Now().Format("2006/01/02 15:04:05")
	timestampMessage := fmt.Sprintf("%s %s", now, message)

	// memory에 log 저장
	checker.logArr = append(checker.logArr, timestampMessage)

	// log가 일정 수치 이상 쌓이면 file write하고 array 초기화
	if len(checker.logArr) >= 100 {
		outFile, err := os.OpenFile("access.log", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer outFile.Close()

		for _, log := range checker.logArr {
			if _, err := outFile.WriteString(log + "\n"); err != nil {
				panic(err)
			}
		}

		checker.logArr = checker.logArr[0:0]
	}
}

// TODO : 이상접속 확인 시 메세지 혹은 이메일 등으로 알림
