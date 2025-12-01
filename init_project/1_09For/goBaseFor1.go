package main

import (
	"fmt"
	"time"
)

func main() {
	// -------------------break语句------------------
	// 中断for循环
	for i := 0; i < 5; i++ {
		if i == 3 {
            break
        }
        fmt.Println("break语句，第", i, "次循环") // 输出第1次循环、第2次循环、第3次循环
	}
	// 中断switch
	switch i := 1; i {
	case 1:
		fmt.Println("break语句，i等于1") // 输出: i等于1
		if i == 1 {
            break
        }
	case 2:
		fmt.Println("break语句，i等于2")
	default:
		fmt.Println("break语句，i不等于1或2")
	}
	// 中断select
	select {
		case <-time.After(time.Second*2):
			fmt.Println("break语句，过了2秒")
			break
		case <- time.After(time.Second*1):
			fmt.Println("break语句，进过了1秒") // 这里
			if true {
				break
			}
        	fmt.Println("break 之后")
		// default:
		// 	fmt.Println("没有信号") // 有default语句，则不会执行break语句
	}

	// 不使用标记 , 循环3次外部，3次内部因为内部一次就break了，所以只输出一次
    for i := 1; i <= 3; i++ {
        fmt.Printf("break语句，不使用标记,外部循环, i = %d\n", i) 
        for j := 5; j <= 10; j++ {
            fmt.Printf("break语句，不使用标记,内部循环 j = %d\n", j)
            break
        }
    }

	// 使用标记，循环1次外部，1次内部就break outter了，所以只输出一次
outter:
    for i := 1; i <= 3; i++ {
        fmt.Printf("break语句，使用标记,外部循环, i = %d\n", i)
        for j := 5; j <= 10; j++ {
            fmt.Printf("break语句，使用标记,内部循环 j = %d\n", j)
            break outter
        }
    }
	// --------------------continue语句-------------------
	// 跳过当前循环，继续下一次循环
	// 中断for循环
    for i := 0; i < 5; i++ {
        if i == 3 { // i等于3时，跳过当前循环，继续下一次循环
            continue
        }
        fmt.Println("continue语句，第", i, "次循环")
    }

    // 不使用标记
    for i := 1; i <= 2; i++ {
        fmt.Printf("continue语句，不使用标记,外部循环, i = %d\n", i)
        for j := 5; j <= 10; j++ {
            fmt.Printf("continue语句，不使用标记,内部循环 j = %d\n", j)// 7\8\9\10都输出
            if j >= 7 { // j大于等于7时，跳过当前循环，继续下一次循环
                continue
            }
			// 7\8\9\10不输出
            fmt.Println("continue语句，不使用标记，内部循环，在continue之后执行")
        }
    }

    // 使用标记，j=7时continue outter1,跳出内部循环  剩余逻辑 ，执行下一次外部循环
outter1:
    for i := 1; i <= 3; i++ {
        fmt.Printf("continue语句，使用标记,外部循环, i = %d\n", i)
        for j := 5; j <= 10; j++ {
            fmt.Printf("continue语句，使用标记,内部循环 j = %d\n", j)
            if j >= 7 {
                continue outter1
            }
            fmt.Println("continue语句，使用标记，内部循环，在continue之后执行")
        }
    }

	// --------------------goto语句-----------------
	gotoPreset := false

	preset:
		a := 5

	process:
		if a > 0 {
			a--
			fmt.Println("goto语句，当前a的值为：", a)
			goto process
		} else if a <= 0 {
		// elseProcess:
			if !gotoPreset {
				gotoPreset = true
				goto preset
			} else {
				goto post
			}
		}

	post:
		fmt.Println("goto语句，main将结束，当前a的值为：", a)
}