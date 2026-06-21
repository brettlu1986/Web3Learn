package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello Go")
// }

//命令行计算器

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("usage: go run .<add|sub|mul|div> <a> <b>")
		return
	}

	op := os.Args[1]

	a, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("invalid number a:", err)
		return
	}

	b, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("invalid num b:", err)
		return
	}

	result, err := Calculate(op, a, b)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println("result:", result)
}

func Calculate(op string, a int, b int) (int, error) {
	switch op {

	case "add":
		return a + b, nil
	case "sub":
		return a - b, nil
	case "mul":
		return a * b, nil
	case "div":
		if b == 0 {
			return 0, fmt.Errorf("Division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operation:%s", op)
	}
}
