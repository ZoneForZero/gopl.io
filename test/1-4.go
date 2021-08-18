// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)


// 输出出现重复行的文件的名称
func main() {
	files := os.Args[1:]
	resultChan := make(chan int)
	if len(files) == 0 {
		fmt.Println("you need to input the file!")
		return
	} else {
		for _, file := range files {
			go func(fileName string) int {
				numberMap := make(map[string]int)
				fileData, err := os.Open(fileName)
				defer fileData.Close()
				if err != nil {
					fmt.Printf("The file %s is error!\n", fileName)
					resultChan <- 1001
					return 1001
				}
				input := bufio.NewScanner(fileData)
				for input.Scan() {
					if numberMap[input.Text()] == 1 {
						fmt.Printf("file: %s , value: %s \n",fileName,input.Text())
						resultChan <- 1
						return 1
					}
					numberMap[input.Text()]++
				}
				resultChan <- 2
				return 2
			}(file)
		}
	}
	for index, _ := range files {
		if value1 := <-resultChan; value1 != 100 {
			fmt.Printf("等到第%d个结果了\n",index)
		}
	}
	return
}