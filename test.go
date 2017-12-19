package main

import "fmt"


//定义接口  里面说明包含的方法
type Phone interface{
  call()
}

//定义结构体
type NokiaPhone struct{

}

//实现接口方法
func (nokiaPhone NokiaPhone) call() {
    fmt.Println(" I am NokiaPhone, I can call you now yes  ")  
}

//调用
func  main() {
  var phone Phone
  phone = new(NokiaPhone)
  phone.call()
  
}