package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// 设置一下输出格式
	log.SetFlags(log.Lshortfile | log.Ltime)
	log.Println("Hello World!")

	// 面向对象学习
	puTongRen := Human{
		name:  "普通人",
		age:   48,
		phone: "13812345678",
	}

	xuesheng := Student{
		Human: Human{
			name:  "何同学",
			age:   8,
			phone: "123001",
		},
		school: "北京邮电大学",
		loan:   -1800,
	}

	dagongren := Employee{
		Human: Human{
			name:  "打工人",
			age:   30,
			phone: "996996",
		},
		company: "字节跳动",
		money:   30000,
	}

	puTongRen.SayHi()
	xuesheng.SayHi()
	dagongren.SayHi()

	// 启动一个Get请求
	listenGetWelcomeRequest()
	listenRegister()

	// 阻塞端口，监听请求
	http.ListenAndServe(":9000", nil)

}

// 模拟一个不含参的Get请求
func listenGetWelcomeRequest() {
	http.HandleFunc("/welcome", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			writer.Write([]byte("Let's go"))
		}
	})
}

// 声明结构体（类） ``json别名
type User struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Code    int
	Message string
}

// 注册的内部实现
func register(writer http.ResponseWriter, request *http.Request) {
	var res Response
	// 延时返回，return之后执行
	defer func() {
		//将model对象序列化成json对象
		arr, _ := json.Marshal(res)
		//取第一个参数通过writer返回
		writer.Write(arr)
	}()
	if request.Method != "POST" {
		return //无效返回
	}

	// 解析入参
	bds, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Println(err)
		return
	}

	user := User{}
	json.Unmarshal(bds, &user)

	if len(user.Name) < 4 || len(user.Password) < 6 {
		res.Code = 100001
		res.Message = "注册参数不合法"
	} else {
		res.Code = 100000
		res.Message = "注册成功"
	}

	return
}

// 模拟一个注册的POST请求
func listenRegister() {
	http.HandleFunc("/register", register)
}

// 模拟一个登陆的POST请求
