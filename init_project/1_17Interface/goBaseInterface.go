package main

import (
	"fmt"
)
// PaymentMethod 接口定义了支付方法的基本操作
type PayMethod interface {
	Account
	Pay(amount int) bool
}
// 另一种定义方式，接口中声明的方法，参数可以没有名称。
type PayMethod1 interface {
    Pay1(int)
}

type Account interface {
	GetBalance() int
	// Deposit(amount int) bool
}

// 信用卡
type CreditCard struct {
	balance int
	limit   int
}
// 接收者 是 指针类型，
func (c *CreditCard) GetBalance() int {
	return c.balance
}

func (c *CreditCard) Pay(amount int) bool {
	if c.balance + amount <= c.limit {
		c.balance += amount
		fmt.Println("CreditCard 支付成功")
		return true
	} else {
		fmt.Println("CreditCard 支付失败,超出信用卡限额")
		return false
	}
}
// 返回 balance
func (c *CreditCard) Pay1(amout int) {
    if c.balance < amout {
        fmt.Println("余额不足")
        return
    }
    c.balance -= amout
}

// 借记卡
type DebitCard struct {
	balance int
}

func (d *DebitCard) GetBalance() int {
	return d.balance
}

func (d *DebitCard) Pay(amount int) bool{
	if d.balance >= amount {
		d.balance -= amount
		fmt.Println("DebitCard 支付成功")
		return true
	} else {
		fmt.Println("DebitCard 支付失败,余额不足")
		return false
	}
}

func purchaseItem(pm PayMethod, price int) {
	if pm.Pay(price) {
		fmt.Printf("购买成功，剩余余额: %d\n", pm.GetBalance())
	} else {
		fmt.Println("购买失败")
	}
}
// 如果函数参数使用 interface{}可以接受任何类型的实参。
// 同样，可以接收任何类型的值也可以赋值给 interface{}类型的变量
func anyParam(param interface{}) {
    fmt.Println("param: ", param)
}

func main() {
	// 实例化信用卡
	creditCard := &CreditCard{balance: 0, limit: 1000}
	// 实例化借记卡
	debitCard := &DebitCard{balance: 500}
	// 购买商品
	fmt.Println("第一次使用信用卡购买:")
	purchaseItem(creditCard, 800)
	fmt.Println("第一次使用借记卡购买:")
	purchaseItem(debitCard, 300)
	fmt.Println("第二次使用借记卡购买:")
    purchaseItem(debitCard, 300)
	fmt.Println("第二次使用信用卡购买:")
    purchaseItem(creditCard, 300)
	
	c := CreditCard{balance: 100, limit: 1000}
    var a Account = &c
    fmt.Println(a.GetBalance())
	//c.GetBalance() 也可以调用，但是不建议，
	// go隐式接口把c转为&c，再使用(*c).GetBalance(),是一个语法糖
	fmt.Println(c.GetBalance()) 
	// 下面会报错，任何方法是指针接收者，则必须用指针。因为GetBalance不是指针接收者
	//purchaseItem(c, 50)

	c.Pay1(300)
    var p PayMethod1 = &c
    fmt.Println("p.Pay1: ", p)

    var b interface{} = &c
    fmt.Println("b: ", b)
    // 这里的anyParam接收任何类型参数，包括int,string,CreditCard,PayMethod等
    anyParam(c)
    anyParam(1)
    anyParam("123")
    anyParam(p)
}


