package main

import (
	_ "fmt"
	"sort"
	"sync"
)

type Key struct {
	Path, Country string
}

func main()  {

	var m map[string] string
	m = make(map[string]string)
	m["route"] = "1"
	test(m["route"] == "1")

	n := len(m)
	test(n == 1)

	delete(m, "route")
	_, ok := m["route"]
	test(ok == false)
	m["root"] = "admin"

	for key,value := range m{
		test(key == "root")
		test(value == "admin")
	}

	commits := map[string]int{
		"rsc":1234,
		"r":2138,
		"gri":1908,
		"adg":912,
	}
	test(len(commits) == 4)

	empty := map[string]int{}
	test(len(empty)== 0)

	hits := make(map[string]map[string]int)
	test(hits["/doc/"]["au"] == 0)

	add := func(m map[string]map[string]int, path, country string){
		mm, ok := m[path]
		if !ok{
			mm = make(map[string]int)
			m[path] = mm
		}
		mm[country]++
	}

	add(hits, "/doc/", "au")
	test(hits["/doc/"]["au"] == 1)

	hit2 := make(map[Key]int)
	hit2[Key{"/","vn"}]++
	test(hit2[Key{"/ref/spec", "cn"}] ==0)

	var counter  = struct {
		sync.RWMutex
		m map[string]int
	}{m : make(map[string]int)}

	counter.RLock()
	count := counter.m["some_key"]
	counter.RUnlock()
	test(count == 0)

	var unordered map[int]string
	var kes []int
	for k := range unordered{
		kes = append(kes, k)
	}
	sort.Ints(kes)
	test(len(unordered) == 0)

}

func test(expr bool) {
	if !expr{
		panic(0)
	}
}

func use(vals ...interface{})  {
	for _, val := range vals{
		_ = val
	}
}