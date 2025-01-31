package main

import "fmt"

type User struct {
	first_name, last_name string
	age int
}

func (u User) SayHello() {
	age := u.age
	var msg string
	if age <= 10 {
		msg = fmt.Sprintf("Hello! Your name is %s %s. You are %d years old. You are a child", u.first_name, u.last_name, u.age)
	}else if age > 10 && age < 18 {
		msg = fmt.Sprintf("Hello! Your name is %s %s. You are %d years old. You are a teenager", u.first_name, u.last_name, u.age)
	}else {
		msg = fmt.Sprintf("Hello! Your name is %s %s. You are %d years old. You are a old", u.first_name, u.last_name, u.age)
	}
	fmt.Println(msg)
}

func main() {
	user := User{first_name: "John", last_name: "Doe", age: 7}
	second_user := User{"Diana","Teen",17}
	grandFather := User{"Andrey", "Oldman", 54}
	user.SayHello()
	second_user.SayHello()
	grandFather.SayHello()
}