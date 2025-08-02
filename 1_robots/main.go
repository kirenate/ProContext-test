package main

import (
	"fmt"
	"time"
)

func main() {
	var black, left, right int
	_, err := fmt.Scan(&black, &right)
	if err != nil {
		panic(err)
	}
	now := time.Now()

LeftRobot:
	if CheckIfMet(left, black) {
		goto RightRobot
	}
	if left < black {
		MR(&left)
	}
	goto LeftRobot
RightRobot:
	if CheckIfMet(right, black) {
		passed := time.Since(now)
		fmt.Printf("robots have met after %v seconds", int(passed.Seconds()))
		return
	}
	if right > black {
		ML(&right)
	}
	goto RightRobot
}

func ML(robot *int) {
	*robot--
	time.Sleep(1 * time.Second)
}

func MR(robot *int) {
	*robot++
	time.Sleep(1 * time.Second)
}

func CheckIfMet(robot, black int) bool {
	if robot == black {
		return true
	}
	return false
}
