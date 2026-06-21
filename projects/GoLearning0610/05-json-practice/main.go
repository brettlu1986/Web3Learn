package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	RPCURL     string `json:"rpc_url"`
	DatabasURL string `json:"database_url"`
	Port       int    `json:"port"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	//按照json格式解析，并放到cfg指针
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func main() {
	//go run ./05-json-practice   运行这个的时候，因为还是在 根目录下，所以config.json还是要带上05-json-practice/
	//如果是 debug， 因为配置里面配置的是 "program": "${workspaceFolder}/05-json-practice"，所以工作目录已经变到根目录了，所以要改成
	//cfg, err := LoadConfig("config.json")
	cfg, err := LoadConfig("05-json-practice/config.json")
	if err != nil {
		fmt.Println("load config failed:", err)
	}

	fmt.Println("rpc:", cfg.RPCURL)
	fmt.Println("db:", cfg.DatabasURL)
	fmt.Println("port", cfg.Port)
}
