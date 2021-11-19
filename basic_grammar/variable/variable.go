package variable

// 变量

import "fmt"

import "go_playground/basic_grammar/utils"

//Go语言的基本类型有：
//bool
//string
//int、int8、int16、int32、int64
//uint、uint8、uint16、uint32、uint64、uintptr
//byte // uint8 的别名
//rune // int32 的别名 代表一个 Unicode 码
//float32、float64
//complex64、complex128

// 集中定义变量
var (
	// a int         // 整形
	// b float64     // 64位浮点数
	c [3]string // 字符串数组
	// d func() bool // 函数指针
	e struct {
		x int
		y float32
	} // 结构体
)

func PrintVars() {
	utils.FuncStart("Print variables")

	c[0] = "I"
	c[1] = "am"
	c[2] = "OK"
	fmt.Println("var c is:", c)

	e.x, e.y = 1, 2
	fmt.Println("var e is:", e)

	//定义变量并初始化
	//不提供数据类型
	//只能在函数内部
	f := 1
	fmt.Println("var f is:", f)
	utils.FuncEnd("Print variables")
}

// GetData 匿名变量、哑元变量
// _ :匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。
func GetData() (int, int) {
	return 100, 200
}

func AnonymousVar() {
	utils.FuncStart("Test Anonymous")
	a, _ := GetData()
	_, b := GetData()
	fmt.Println("a is", a, ",b is", b)
	utils.FuncEnd("Test Anonymous")
}
