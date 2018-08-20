package main

import "fmt"

//go:generate go run ./buildtools/embed.go -in=./assets/test.txt -out=test-asset.go -name=TestTxt
//go:generate go run ./buildtools/embed.go -in=./assets/hw.txt -out=hw-asset.go -name=HwTxt

func main() {
	fmt.Println(string(TestTxt))
	fmt.Println(string(HwTxt))
}
