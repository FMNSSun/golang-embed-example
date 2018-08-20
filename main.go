package main

import "fmt"
import "github.com/FMNSSun/golang-embed-example/sub"

func main() {
	fmt.Println(string(TestTxt))
	fmt.Println(string(HwTxt))
	fmt.Println(string(sub.SecretPassword))
}
