// Exercise 7.03 – analyzing empty interface{} data

package main

import "fmt"

type record struct {
	key       string
	valueType string
	data      interface{}
}

type person struct {
	lastName  string
	age       int
	isMarried bool
}

type animal struct {
	name     string
	category string
}

func newRecord(key string, i interface{}) record {
	r := record{}
	r.key = key
	switch v := i.(type) {
	case int:
		r.valueType = "int"
		r.data = v
	case bool:
		r.valueType = "bool"
		r.data = v
	case string:
		r.valueType = "string"
		r.data = v
	case person:
		r.valueType = "person"
		r.data = v
	default:
		r.valueType = "unknown"
		r.data = v
	}

	return r
}

func main() {
	m := make(map[string]interface{})
	a := animal{name: "oreo", category: "cat"}
	p := person{lastName: "Danang", isMarried: false, age: 31}
	m["age"] = 54
	m["isMarried"] = true
	m["lastname"] = "Smith"
	m["person"] = p
	m["animal"] = a

	rs := []record{}
	for k, v := range m {
		r := newRecord(k, v)
		rs = append(rs, r)
	}

	for _, v := range rs {
		fmt.Println("Key", v.key)
		fmt.Println("Data", v.data)
		fmt.Println("Type", v.valueType)
		fmt.Println()
	}
}
