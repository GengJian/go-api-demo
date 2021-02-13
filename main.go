package main

import (
	"log"
	"net/http"
)

func main() {
	// 设置一下输出格式
	log.SetFlags(log.Lshortfile | log.Ltime)
	log.Println("Hello World!")

	// 启动一个Get请求
	listenGetWelcomeRequest()
}

// 模拟一个不含参的Get请求
func listenGetWelcomeRequest() {
	http.HandleFunc("/welcome", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			writer.Write([]byte("Let's go"))
		}
	})

	http.ListenAndServe(":9000", nil)
}
