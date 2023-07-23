package concurrency

import "fmt"

func init() {
	fmt.Println("init concurrency package")
}

func Init() {
	fmt.Println("concurrency init from external")
}
