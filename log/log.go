package cLog

import (
	"fmt"
	"time"
)

/*
Custom Log
*/

// The basic logs: [Go_auth](date) message

// Print is similar to fmt.Print()
func Print(msg string) {
	now := time.Now()
	date := fmt.Sprintf("%d/%d %d:%d:%d", now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("[Go_auth](%s) %s", date, msg)
}

// Println is similar to fmt.Println()
func Println(msg string, a ...any) {
	now := time.Now()
	date := fmt.Sprintf("%d/%d %d:%d:%d", now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	text := fmt.Sprintf(fmt.Sprintf("[Go_auth](%s) %s \n", date, msg), a...)
	fmt.Printf(text)
}

// Printf Similar to fmt.Printf()
func Printf(msg string, a ...any) {
	text := fmt.Sprintf(msg, a...)
	Println(text)
}

// Warning Warnings Logs: [Go_auth] [WARNING] (date) message
func Warning(msg string) {
	now := time.Now()
	date := fmt.Sprintf("%d/%d %d:%d:%d", now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("[Go_auth] \x1b[31m%s\x1b[0m (%s) %s \n", "[WARNING]", date, msg)
}

// Fatal is FATAL
func Fatal(msg string) {
	now := time.Now()
	date := fmt.Sprintf("%d/%d %d:%d:%d", now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(fmt.Sprintf("[Go_auth] \x1b[31m%s\x1b[0m (%s) %s \n", "[FATAL]", date, msg))
}

func Fatalf(msg string, a ...any) {
	Fatal(fmt.Sprintf(msg, a...))
}
