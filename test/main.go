package main

import (
	"fmt"
	"log"
	"time"
)

type struct1 struct {
	struct2  struct2
	struct22 *struct2
}
type struct2 struct {
	struct3
	s3 *struct3
}

type struct3 struct {
	name string
}

var structObj = new(struct1)

type student struct {
	Name string
	Age  int
}

func addMap(mm map[string]string) {

	mm["3"] = "3"
}

func addSlice(mm []string) {
	mm = append(mm, "addslice")
	fmt.Println(mm)
}

type Per struct {
	Name string
	Age  int
}

func main() {
	fmt.Println(compareDate(time.Now().Add(-100*time.Minute).Format(time.RFC3339), time.Now().Format(time.RFC3339)))

	pers := []Per{{Name: "abc", Age: 12}, {Name: "bcd", Age: 22}}
	for _, v := range pers {
		aa := &v
		aa.Age = 33
		aa.Name = "abcddd"
	}
	fmt.Printf("pers: %v\n", pers[0])
	fmt.Printf("pers: %v\n", pers[1])
	for i := 0; i < len(pers); i++ {
		pers[i].Age = 33
		pers[i].Name = ""
	}
	fmt.Printf("pers: %v\n", pers[0])
	fmt.Printf("pers: %v\n", pers[1])

	fmt.Printf("structObj.struct2.struct3: %#v\n", structObj.struct2.struct3)
	fmt.Printf("structObj.struct2.struct3: %+v\n", structObj.struct22.s3)

	fmt.Printf("f(5): %v\n", f(5))

	sStr := make([]string, 2)
	sStr = append(sStr, "0")
	sStr = append(sStr, "1")
	sStr = append(sStr, "2")
	fmt.Println(sStr)
	addSlice(sStr)
	fmt.Println(sStr)

	mStr := make(map[string]string, 2)

	mStr["0"] = "0"
	mStr["1"] = "1"
	mStr["2"] = "2"
	fmt.Println(mStr)
	addMap(mStr)
	fmt.Println(mStr)

	fmt.Println(compareDate(time.Now().Add(-100*time.Minute).Format(time.RFC3339), time.Now().Format(time.RFC3339)))
	m := make([]*student, 3)
	m1 := make([]*student, 0)
	stus := []*student{
		{Name: "sa", Age: 10},
		{Name: "sb", Age: 11},
		{Name: "sc", Age: 12},
	}
	log.Println("################ 错误做法 ##################")
	for k, stu := range stus {
		m[k] = stu
		m1 = append(m1, stu)
	}
	for _, s := range m {
		log.Println(s.Name, s.Age)
	}

	for _, s := range m1 {
		log.Println(s.Name, s.Age)
	}
	log.Println("################ 正确做法 ##################")
	for k, _ := range stus {
		m[k] = stus[k]
	}
	for _, s := range m {
		log.Println(s.Name, s.Age)
	}
	time.Sleep(2 * time.Second)
}

func compareDate(timeStr1, timeStr2 string) bool {
	fmt.Printf("timeStr1: %v\n", timeStr1)
	fmt.Printf("timeStr2: %v\n", timeStr2)
	t, err := time.Parse(time.RFC3339, timeStr1)
	if err != nil {
		return true
	}

	tt, err := time.Parse(time.RFC3339, timeStr2)
	if err != nil {
		return true
	}
	return t.After(tt)
}

func f(n int) []int {
	if n <= 2 {
		return []int{1}
	}
	sliceInt := make([]int, n)
	sliceInt[0] = 1
	sliceInt[1] = 1
	for i := 2; i < n; i++ {
		sliceInt[i] = sliceInt[i-1] + sliceInt[i-2]
	}
	return sliceInt
}
