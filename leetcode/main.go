package main

import (
	"fmt"
	"reflect"
	"strings"
)

/*
Go语言的优势
1 简单的部署方式
	可直接编译成机械码
	不信赖其他库
	直接运行即可部署
2 静态类型语言
	编译的时候可以检测出隐藏的大多数问题
	强类型方便阅读与重构
3 语言层面的并发
	天生的基于支持
	充分利用多核
4 工程化比较优秀
	GoDoc可以直接从代码和注释生成漂亮的文档
	GoFmt 统一的代码格式
	GoLint代码语法提示
	测试框架内置
5 强大的标准库
	Runtime系统调度机制
	高效的GC垃圾回收
	丰富的标准库
6 简单易学
	25个关键字
	面向对象特征
	跨平台


*/
func main() {

	//1 无重复字符串的最长子串
	// s := "abcdfabdfcc"
	// lengthOfLongestSubstring(s)

	//2 快速排序
	// nums := []int{3, 5, 2, 3, 4, 6, 7, 4, 9}
	// fmt.Printf("QuickSort(nums): %v\n", QuickSort(nums))

	//3 连表反转
	// node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	// ln := reverseList(node)

	// for ln != nil {
	// 	fmt.Printf("ln: %#v\n", ln)
	// 	ln = ln.Next
	// }

	//4 翻转字符串
	//reverseString()

	//5 打印tag
	// printTag(&UserInfo{Name: "zhouli", Age: 33})

	//6 slice 扩容
	sliceDouble()
}

func sliceDouble() {
	var arr0 []int16
	var arr1 []int32
	var arr2 []int64
	var arr3 []int64

	for i := 0; i < 1025; i++ {
		arr0 = append(arr0, int16(i))
		arr1 = append(arr1, int32(i))
		arr2 = append(arr2, int64(i))
	}
	fmt.Printf("len=%d,cap=%d\n", len(arr0), cap(arr0))
	//1024*125＊4 = 5120 查表(sizeclasses.go) 5376／4 ＝ 1344//sizeclassse 最接近5120的值是5376
	fmt.Printf("len=%d,cap=%d\n", len(arr1), cap(arr1))
	//1024*125＊8 = 10240 查表(sizeclasses.go) 10240／8 ＝ 1280
	fmt.Printf("len=%d,cap=%d\n", len(arr2), cap(arr2))
	arr3 = append(arr3, arr2...)
	//1025*8 = 8200 查表(sizeclasses.go) 9472／8 ＝ 1184
	fmt.Printf("len=%d,cap=%d\n", len(arr3), cap(arr3))
}

type UserInfo struct {
	Name string `json:"name" bilbil:"name bil"`
	Age  int    `json:"age" bilbil:"age bil"`
}

func printTag(tagObj interface{}) {
	reType := reflect.TypeOf(tagObj)
	fmt.Printf("reType: %T,\n%#v,\n%+v\n", reType, reType, reType)
	fmt.Printf("reType: %v,\n%v,\n%+v\n", reType.Elem().Kind(), reType.Kind(), reType)

	reValue := reflect.ValueOf(tagObj)
	fmt.Printf("reValue: %v\n,%v", reValue, reValue.Elem())
	for i := 0; i < reValue.Elem().NumField(); i++ {
		field := reValue.Elem().Type().Field(i)
		tag := field.Tag
		// fmt.Printf("tag: %v\n", tag)
		fmt.Printf("tag.Get(\"json\"): %v\n", tag.Get("json"))
		fmt.Printf("tag.Get(\"bilbil\"): %v\n", tag.Get("bilbil"))
	}
}

func reverseString() {
	s := "我是中国人，I love china!"

	r := []rune(s)
	fmt.Printf("r: %v\n", string(r))
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	fmt.Printf("r: %v\n", string(r))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil

	for cur != nil {
		fmt.Printf("cur: %#v\n", cur)
		pre, cur, cur.Next = cur, cur.Next, pre

		//不知道为什么使用下面的三行代码不行
		// pre = cur
		// cur = cur.Next
		// cur.Next = pre
	}

	return pre
}

//快速排序
/*
取切片第一个值做为标识值，把大于这个值的放高位切片，小于这个值的放低位切片，一直递归
*/
func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	low := make([]int, 0, 0)
	high := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	flag := nums[0]
	mid = append(mid, flag)
	for i := 1; i < len(nums); i++ {
		if nums[i] > flag {
			high = append(high, nums[i])
		} else if nums[i] < flag {
			low = append(low, nums[i])
		} else {
			mid = append(mid, nums[i])
		}
	}
	low, high = QuickSort(low), QuickSort(high)
	return append(append(low, mid...), high...)
}

// 无重复字符串的最长子串
func lengthOfLongestSubstring(s string) {
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		fmt.Printf("%s \n", fmt.Sprintf("s[start:i]==%s,s[i]==%s", s[start:i], string(s[i])))
		fmt.Printf("start,end ,i: %d,%d,%d,%v,%v\n", start, end, i, (i+1) > end, index == -1)
		if index == -1 && (i+1) > end {
			end = i + 1
		} else {
			start += index + 1
			end += index + 1
		}
	}
	fmt.Printf("start,end  %d,%d,\n", end, start)
}
