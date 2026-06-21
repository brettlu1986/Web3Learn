package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//go run ./06-log-parser   运行这个的时候，因为还是在 access.log 还是要带上06-log-parser/
	//如果是 debug， 因为配置里面配置的是 "program": "${workspaceFolder}/06-log-parser"，所以工作目录已经变到根目录了，所以要改成
	//stats, err := ParseLog("access.log")
	stats, err := ParseLog("06-log-parser/access.log")
	if err != nil {
		fmt.Println("parse log failed", err)
		return
	}

	for key, count := range stats {
		fmt.Printf("%s: %d \n", key, count)
	}
}

func ParseLog(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close() //注册一个file defer close，在函数执行返回前最后执行，也可以防止中途返回没有close

	stats := make(map[string]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		if strings.HasPrefix(line, "INFO") {
			stats["INFO"]++
		}

		if strings.HasPrefix(line, "ERROR") {
			stats["ERROR"]++
		}

		if strings.Contains(line, "rpc timeout") {
			stats["rpc timeout"]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}
