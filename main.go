package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// 设置一下输出格式
	log.SetFlags(log.Lshortfile | log.Ltime)
	log.Println("Hello World!")

	/* 面向对象学习 */
	//oopTest()

	/* Interface 学习*/
	interfaceTest()

	/* 启动Web服务 */
	//requestTest()

}

// 面向对象 oop test
//goland:noinspection GoUnusedFunction
func oopTest() {
	commonPeople := Human{
		name:  "普通人",
		age:   48,
		phone: "13812345678",
	}

	studentA := Student{
		Human: Human{
			name:  "何同学",
			age:   8,
			phone: "123001",
		},
		school: "北京邮电大学",
		loan:   -1800,
	}
	workerB := Employee{
		Human: Human{
			name:  "打工人",
			age:   30,
			phone: "996996",
		},
		company: "字节跳动",
		money:   30000,
	}

	commonPeople.SayHi()
	studentA.SayHi()
	workerB.SayHi()
}

// 注册接口 request test
//goland:noinspection ALL
func requestTest() {
	// 启动一个Get请求
	listenGetWelcomeRequest()
	listenRegister()

	// 阻塞端口，监听请求
	http.ListenAndServe(":9000", nil)
}

// Men 接口测试 interface test
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func interfaceTest() {
	mike := Student{Human{"Mike", 25, "111-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i（类似泛型+协议）
	var i Men

	//i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee
	i = tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}

	// 测试实现String的接口,重写desperation
	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)

}

// 模拟一个不含参的Get请求
func listenGetWelcomeRequest() {
	http.HandleFunc("/welcome", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			writer.Write([]byte("Let's go"))
		}
	})
}

// User 声明结构体（类） ``json别名
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
