package cLog

import (
	"fmt"
	"time"
)

/*
Custom Log

The basic log: [Go_auth](date) message
*/

// fmt.Print()
func Print(msg string) {
	now := time.Now()
	date := fmt.Sprintf("%d/%d %d:%d:%d", now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("[Go_auth](%s) %s", date, msg)
}

// Similar to fmt.Println()
func Println(msg string, a ...any) {
	now := time.Now()
	date := fmt.Sprintf("%d/%d %d:%d:%d", now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	text := fmt.Sprintf(fmt.Sprintf("[Go_auth](%s) %s \n", date, msg), a...)
	fmt.Printf(text)
}

// Similar to fmt.Printf()
func Printf(msg string, a ...any) {
	text := fmt.Sprintf(msg, a...)
	Println(text)
}
