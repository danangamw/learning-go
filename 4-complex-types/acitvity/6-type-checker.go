// Activity 4.06 – Type checker

package main

import "fmt"

func getData() []interface{} {
	return []interface{}{
		1,
		3.14,
		"hello",
		true,
		struct{}{},
	}
}

func typeCheck(v any) string {
	switch t := v.(type) {
	case string:
		return t + " is string"
	case int, int32, int64:
		return fmt.Sprintf("%v is int", t)
	case float32, float64:
		if f, ok := t.(float64); ok {
			return fmt.Sprintf("%v is float", f)
		}

		return fmt.Sprintf("%v is float", t.(float32))
	case bool:
		if t {
			return "true is bool"
		}

		return "false is bool"
	default:
		return fmt.Sprintf("%v is unknown", t)
	}

}

func main() {
	data := getData()

	for i := 0; i < len(data); i++ {
		fmt.Println(typeCheck(data[i]))
	}
}
