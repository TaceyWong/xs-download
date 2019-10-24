package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("LANG", "zh_CN.UTF-8")
	fmt.Println(os.Getenv("LANG"))
	start()
}
