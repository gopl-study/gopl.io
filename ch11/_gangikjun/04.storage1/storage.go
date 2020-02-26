package storage

import (
	"fmt"
	"net/smtp"
)

var usage = make(map[string]int64)

func bytesInUse(username string) int64 {return usage[username]}

// Email 송신자 설정
// NOTD: 절대 소스코드에 암호를 넣지 마시오
const sender = "notifications@example.com"
const password = "qwerasdfzxcv"
const hostname = "smtp.example.com"

const template = `Warning: you are using %d bytes of storage,
%d%% of your quota`

// CheckQuota CheckQuota
func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return // OK
	}
	msg := fmt.Sprintf(template, used, percent)
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sener, []string{username},[]byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
	}
}