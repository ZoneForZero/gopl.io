package main

import (
	"fmt"
	"os"
)

func main() {
	// 1-1
	fmt.Printf("The os.Args[] is %s \n", os.Args[0]);

	// 1-2
	for index, value := range os.Args {
		// f结尾的函数大多数有格式化，但是ln会以%v去处理，并在结尾追加换行符
		fmt.Println("index: ",index, ", value:",value);
	}
}