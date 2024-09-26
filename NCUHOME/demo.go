package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var container = make(map[string]string)

func usemain() {
	fmt.Println("欢迎使用命令行工具")
	fmt.Println("1) 命令行")
	fmt.Println("2) 使用说明")
	fmt.Println("3) 退出程序")
}

func use() {
	fmt.Println("命令行")
	fmt.Println("1) SET")
	fmt.Println("2) SETNX")
	fmt.Println("3) GET")
	fmt.Println("4) DEL")
	fmt.Println("5) 保存到JSON文件")
}

func SET(key, value string) {
	container[key] = value
	fmt.Println("存储成功")
	SAVE("data.json")
}

func SETNX(key, value string) int {
	if _, ok := container[key]; !ok {
		container[key] = value
		fmt.Println("该键值不存在，已存储")
		SAVE("data.json")
		return 1
	}
	fmt.Println("该键值已存在")
	return 0
}

func GET(key string) string {
	value, exists := container[key]
	if !exists {
		fmt.Println("未找到该键")
		return ""
	}
	fmt.Println("获取成功:", value)
	return value
}

func DEL(key string) {
	if _, exists := container[key]; exists {
		delete(container, key)
		fmt.Println("删除成功")
		SAVE("data.json")
	} else {
		fmt.Println("未找到该键，无法删除")
	}
}

func readInput(prompt string) string {
	fmt.Printf(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

func SAVE(filename string) {
	data, err := json.Marshal(container)
	if err != nil {
		fmt.Println("无法序列化为JSON:", err)
		return
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("无法写入文件:", err)
		return
	}
	fmt.Println("数据已保存到JSON文件")
}

func main() {
	var choice1, choice2 int
	var key, value string

	for {
		usemain()
		choice1Str := readInput("请输入(1~3)：")
		fmt.Sscanf(choice1Str, "%d", &choice1)
		switch choice1 {
		case 1:
			use()
			choice2Str := readInput("请输入：")
			fmt.Sscanf(choice2Str, "%d", &choice2)
			switch choice2 {
			case 1:
				key = readInput("请输入想存入的键：")
				value = readInput("请输入想存入的值：")
				SET(key, value)
			case 2:
				key = readInput("请输入想检验的键：")
				value = readInput("请输入想检验的值：")
				SETNX(key, value)
			case 3:
				key = readInput("请输入想要获取的键：")
				GET(key)
			case 4:
				key = readInput("请输入想要删除的键：")
				DEL(key)
			case 5:
				SAVE("data.json")
			default:
				fmt.Println("请输入1~5：")
			}
		case 2:
			fmt.Println("这是使用说明")
		case 3:
			fmt.Println("退出程序")
			return
		default:
			fmt.Println("请输入1~3：")
		}
	}
}
