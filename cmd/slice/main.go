package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]

	// s1 从 slice的索引2 到 索引5，左开右闭。容量没指定，默认到slice结尾
	// s1.len = high - low = 5-2=3
	// sl.cap = max - low = 10-2=8
	// s1=[2 3 4]
	// 但是s1底层数组延伸到9，即[2 3 4 5 6 7 8 9], 只是你直接越界访问如 s1[2] 会 panic

	// s2 从 s1的索引2 到 索引6。容量到 索引7
	// s2.len = high - low = 6-2=4
	// s2.cap = max - low = 7-2=5
	// s2=[4 5 6 7]

	s2 = append(s2, 100)
	s2 = append(s2, 200)
	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}
