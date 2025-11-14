package main

// import ("fmt"; "os"; "path/filepath")
import (
	"fmt"
	"os"
	"path/filepath"
)

// 递归函数计算阶乘
func factorial(n int) int {
	/* 代码解释
    基准条件：当 n 等于 0 时，函数返回 1，因为 0! 定义为 1。
    递归条件：函数返回 n 乘以 factorial(n-1) 的结果，逐步将问题分解为更小的子问题。
 */
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	/* 代码解释
    基准条件：当 n 等于 0 或 1 时，函数返回 n。
    递归条件：函数返回 fabonacci(n-1) + fabonacci(n-2) 的结果，逐步将问题分解为更小的子问题。
 */
	if n == 0 || n == 1 { // 即n<2 时，返回n
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func sqrtRecursive(x, guess, prevGuess, epsilon float64) float64 {
	/* sqrtRecursive 函数接受四个参数：
    x 表示待求平方根的数
    guess 表示当前猜测的平方根值
    prevGuess 表示上一次的猜测值
    epsilon 表示精度要求（即接近平方根的程度）
递归的终止条件是当前猜测的平方根与上一次猜测的平方根非常接近，差值小于给定的精度 epsilon。 */
        if diff := guess*guess - x; diff < epsilon && -diff < epsilon {
                return guess
        }

        newGuess := (guess + x/guess) / 2
        if newGuess == prevGuess {
                return guess
        }

        return sqrtRecursive(x, newGuess, guess, epsilon)
}
// 定义 sqrt 函数，调用 sqrtRecursive 函数 求平方根
func sqrt(x float64) float64 {
        return sqrtRecursive(x, 1.0, 0.0, 1e-9)
}

func main() {
	// 5! = 5 * 4 * 3 * 2 * 1 = 120 阶乘是一个正整数的乘积，表示为 n!。
	fmt.Println("0的阶乘为：", factorial(0))
	fmt.Println("1的阶乘为：", factorial(1))
	fmt.Println("2的阶乘为：", factorial(2))
	fmt.Println("3的阶乘为：", factorial(3))
	fmt.Println("5的阶乘为：", factorial(5))
	var i int
	fmt.Println("斐波那契数列")
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t",fibonacci(i))
	}
	fmt.Println()
	x :=25.0
    result := sqrt(x)
    fmt.Printf("%.2f 的平方根为 %.6f\n", x, result)
	
	// 输出当前目录下的文件和目录
	walkDir(".","")
	// 输出指定目录下的文件和目录
	walkDir("D:\\web3","")
}

func walkDir(dir string, indent string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, entry := range entries {
		fmt.Println(indent + entry.Name())
		if entry.IsDir() {
			fmt.Println("filepath.Join路径",filepath.Join(dir, entry.Name()))
			walkDir(filepath.Join(dir, entry.Name()), indent+" ")
		}
	}
}