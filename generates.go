package main

//go:generate go run ./buildtools/embed.go -in=./assets/test.txt -out=test-asset.go -name=TestTxt -package=main
//go:generate go run ./buildtools/embed.go -in=./assets/hw.txt -out=hw-asset.go -name=HwTxt -package=main -as-string=true
//go:generate go run ./buildtools/embed.go -in=./assets/secret.txt -out=./sub/secret-asset.go -name=SecretPassword -package=sub -as-string=true
