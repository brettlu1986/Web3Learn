package main

import "fmt"

func main() {
	// fmt.Println("------start")
	// fmt.Println(os.Args)
	// fmt.Println("------end")
	//fmt.Println("hello world")
outter:
	for i := 1; i <= 3; i++ {
		fmt.Printf("外部, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("内部 j = %d\n", j)
			if j >= 7 {
				continue outter
			}
			fmt.Println("内部， continue后")
		}
	}
}
