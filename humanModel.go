package main

import (
	. "fmt"
)

// Human 人类
type Human struct {
	name  string
	age   int
	phone string
}

// Student 学生，也属于人类
type Student struct {
	Human  // 匿名字段 Human
	school string
	loan   float32
}

// Employee 职员，也属于人类
type Employee struct {
	Human   // 匿名字段 Human
	company string
	money   float32
}

// SayHi Human 对象实现SayHi方法
func (h Human) SayHi() {
	Printf("Hi, I am %s, you can call me %s \n", h.name, h.phone)
}

// Sing Human 对象实现Sing方法
func (h Human) Sing(lyrics string) {
	Printf("La la ,la la la, la la la ... \n", lyrics)
}

// SayHi Employee 重载Human SayHi方法
func (e Employee) SayHi() {
	Printf("Hi, I am %s, I Work at %s, Call me on %s \n", e.name, e.company, e.phone)
}
