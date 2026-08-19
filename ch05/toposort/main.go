package main

import (
	"fmt"
	"sort"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])           // 先递归访问它的先修课，把先修课都添加到 order 中
				order = append(order, item) // 自己最后才加入 order，也就是说，一门课能被加入 order，前提是它依赖的所有课程已经在它之前被加入了。
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	//  举例：calculus 和 algorithms 两条互不相关的依赖链，如果某次运行先遍历到 calculus，输出就是 linear algebra, calculus, ..., intro to programming, discrete math, data
	//  structures, algorithms, ...；换一次运行先遍历到 algorithms，两条链的相对顺序就反过来了。两者都是"合法"的拓扑序（先修课都在前），但具体序列不同。
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
