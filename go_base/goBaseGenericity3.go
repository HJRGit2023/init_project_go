package main

import (
"fmt"
"sync"
)
/* 1. 定义泛型结构体 Stack[T any] 
2.泛型映射到具体类型（Map） */
// 1.1 泛型栈实践，泛型结构体
type Stack[T any] struct {
	elements []T // 元素列表切片
}
// 1.2 入栈
func (stack *Stack[T]) Push(value T) {
	stack.elements = append(stack.elements, value)
}
// 1.3 出栈
func (stack *Stack[T]) Pop() (T, bool) {
	if len(stack.elements) == 0 {
		var zero T
		return zero, false
	}
	lastIndex := len(stack.elements) - 1
	value := stack.elements[lastIndex]
	stack.elements = stack.elements[:lastIndex] // 切片截取, 丢弃最后一个元素
	return value, true
}

// 1.4 查看栈顶元素
func (stack *Stack[T]) Peek() (T, bool) {
	if len(stack.elements) == 0 {
		var zero T
		return zero, false
	}
	return stack.elements[len(stack.elements)-1], true
}

// 1.5 判断栈是否为空
func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.elements) == 0 // 切片长度为0判断为空
}

// 2.1 泛型映射到具体类型
type Safemap[K comparable, V any] struct {
	data map[K]V // 映射表
	mutex sync.RWMutex // 互斥锁
}

// 2.2 创建新的 SafeMap
func NewSafeMap[K comparable, V any]() *Safemap[K, V] {
	return &Safemap[K, V]{
		data: make(map[K]V),
	}
}

// 2.3设置键值对
func (m *Safemap[K, V]) Set(key K, value V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.data[key] = value
}

// 2.4 获取值
func (m *Safemap[K, V]) Get(key K) (V, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	value, ok := m.data[key]
	return value, ok
}

// 2.5删除键
func (m *Safemap[K, V]) Delete(key K) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.data, key)
}

// 2.6 获取所有键
func (m *Safemap[K, V]) Keys() []K {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	keys := make([]K, 0, len(m.data))
	for key := range m.data {
		keys = append(keys, key)
	}
	return keys
}

func main() {
	// 整数栈
	var intStack Stack[int] // 声明整数栈
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	fmt.Println(intStack.Pop()) // 3, true

	// 字符串栈
	var strStack Stack[string] // 声明字符串栈
	strStack.Push("hello")
	strStack.Push("world")
	strStack.Push("go")
	fmt.Println(strStack.Pop()) // go, true

	// 创建字符串到整数的映射
	scores := NewSafeMap[string,int]()
	scores.Set("Alice", 90)
	scores.Set("Bob", 80)
	scores.Set("Charlie", 70)
	fmt.Println(scores.Get("Alice")) // 90, true
	if score, exist := scores.Get("Alice"); exist {
		fmt.Printf("Alice's score is%d\n", score)
	}
	fmt.Println("Keys：", scores.Keys()) // [Alice Bob Charlie]

}