package main

import "fmt"

func main() {
	///* 定义局部变量 */
	//var grade string = "B"
	//var marks int = 60
	//
	//switch marks {
	//	case 90: grade = "A"
	//	case 80: grade = "B"
	//	case 50,60,70 : grade = "C"
	//	default: grade = "D"
	//}
	//
	//switch {
	//	case grade == "A" :
	//		fmt.Printf("优秀!\n" )
	//	case grade == "B", grade == "C" :
	//		fmt.Printf("良好\n" )
	//	case grade == "D" :
	//		fmt.Printf("及格\n" )
	//	case grade == "F":
	//		fmt.Printf("不及格\n" )
	//	default:
	//		fmt.Printf("差\n" );
	//}
	//fmt.Printf("你的等级是 %s\n", grade );

	var x interface{}

	switch i := x.(type) {
		case nil:
			fmt.Printf(" x 的类型 :%T",i)
		case int:
			fmt.Printf("x 是 int 型")
		case float64:
			fmt.Printf("x 是 float64 型")
		case func(int) float64:
			fmt.Printf("x 是 func(int) 型")
		case bool, string:
			fmt.Printf("x 是 bool 或 string 型" )
		default:
			fmt.Printf("未知型")
	}

	var x2 interface{}

	switch x2.(type) {
		case nil:
			fmt.Printf(" x 的类型是 nil")
		case int:
			fmt.Printf("x 是 int 型")
		case float64:
			fmt.Printf("x 是 float64 型")
		case func(int) float64:
			fmt.Printf("x 是 func(int) 型")
		case bool, string:
			fmt.Printf("x 是 bool 或 string 型" )
		default:
			fmt.Printf("未知型")
	}

	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}
