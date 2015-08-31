

### 懶人執行
`範例` 會自動執行ex1裡的main的go檔案

	$ ./run.sh ex1 main

### 變數宣告
| Component        | Why is it included? / Remarks |
| ---------------- | ------------------- |
| var variableName type | |
| var variab;eName type = value | |
| var var1, var2, var3 type = v1, v2, v3 | |
| var var1, var2, var3 = v1, v2, v3 | |
| var1, var2, var3 := v1, v2, v3 | |
| _, b := 34, 35 | _(底線)是特殊變數，任何指定的值都會被捨棄 |

### 宣告常數
const var1 = value

### 基礎類型


### array
`var arr [n]type` 在[n]type中，n代表陣列長度，type代表儲存元素類型

`arr := [3]int{1, 2, 3}`

`arr := [...]int{4, 5, 6}` 可使用'...'的方式，Go語言自動根據元素個數來計算長度

`doubleArr := [2][4]int{ [4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8} }` 宣告二維陣列，並指定內部類型

`doubleArr := [2][4]int{ {1, 2, 3, 4}, {5, 6, 7, 8} }` 宣告二維陣列，忽略內部類型

### slice
動態陣列

`var arr []int`

`arr := []byte {'a', 'b', 'c', 'd'}`

`arr = ar[2:5]` a指定陣列的第3個元素開始，並到第五個元素結束

`arr = ar[:3]` 相當於arr = arr[0:3]

`arr = ar[5:]` 相當於arr = arr[5:10]

`arr = ar[:]` 相當於arr = arr[0:10]

### map
	var number map[string] int

	number := make(map[string]int)
	number["one"] = 1
	number["ten"] = 10
	number["three"] = 3

	fmt.Println("number 3 : ", number["three"]);

### 流程函數

## if
	if x > 10 {
		fmt.Println("x > 10");
	} else if x == 10 {
		fmt.Println("x == 10");
	} else {
		fmt.Println("x else 10");
	}

## for
	for x:=0; x < 10; index++ {
	}

## switch
加入`fallthrough`強制執行後面的case

加入`break`比對成功後不會自動向下執行
	switch x {
	case 1:
		fmt.Println("x == 1");
	case 2, 3, 4:
		fmt.Println("x == 2, 3, 4");
	default:
		fmt.Println("x == default");
	}

## 函數
	多個傳回值
	func funcName(input1 type1, input2, type2) (output type1, output2 type2){
		return value1, value2
	}

	變參
	func funcName(arg ...int) {}

	函數作為類型
	type funcName func() string
	func ex1() string {
		return "ex1"
	}
	func ex2() string {
		return "ex2"
	}
	func filter(f funcName) string {
		return f()
	}
	func main() {
		string_ex1 := filter(ex1)
		fmt.Printf("ex1 return => %s\n", string_ex1)
		string_ex2 := filter(ex2)
		fmt.Printf("ex2 return => %s\n", string_ex2)
	}

## 物件導向`ex3`
	type Rectangle struct {
		width, height float64
	}

	type Circle struct {
		radius float64
	}

	func (r Rectangle) area() float64 {
		return r.width * r.height
	}

	func (c Circle) area() float64 {
		return c.radius * c.radius * math.Pi
	}

	func main() {
		r1 := Rectangle{12, 2}
		r2 := Rectangle{9, 4}
		c1 := Circle{10}
		c2 := Circle{25}

		fmt.Println("Area of r1 is : 12 * 2 = ", r1.area())
		fmt.Println("Area of r2 is : 9  * 4 = ", r2.area())
		fmt.Println("Area of c1 is : 10 * 10 * 3.14 = ", c1.area())
		fmt.Println("Area of c2 is : 25 * 25 * 3.14 = ", c2.area())

	}

## 平行處理`ex4` goroutine
	Buffered `main2.go`
	Range、Close `main3.go`


## web sample `ex5`

## import other package of sample `ex6`

## **相同package可以互相使用參數


## 判斷型態
`ex1`
	
	import "reflect"
	import "fmt"

	func main() {
		s := "11111"
		fmt.Println(reflect.TypeOf(s))

		i := 11111
		fmt.Println(reflect.TypeOf(i))
	}

`ex2`

	import (
    	"fmt"
    	"reflect"
	)
	func main() {
		z := "hello"
		zv := reflect.TypeOf(z).Kind()
		fmt.Println(zv)
	}

## 參考網站
	1. [https://github.com/Unknwon/go-fundamental-programming](https://github.com/Unknwon/go-fundamental-programming)



