package main

//go:generate go run ./buildtools/embed.go -in=./assets/test.txt -out=test-asset.go -name=TestTxt -package=main
//go:generate go run ./buildtools/embed.go -in=./assets/hw.txt -out=hw-asset.go -name=HwTxt -package=main
