package main

import (
	"./cycle"
)

func main() {
	{
		v := 0
		type test struct {
			v1 *int
			v2 *int
			v3 *int
			link *test
			v4 *int
			v5 *int
		}
		c := &test{
			v1:   &v,
			v2:   &v,
			v3:   &v,
			link: &test{
				v1:   &v,
				v2:   &v,
				v3:   &v,
				v4:   &v,
				v5:   &v,
			},
			v4:   &v,
			v5:   &v,
		}
		print("struct : ")
		println(cycle.IsCycle(c))
	}

	{
		root := make(map[string]interface{})
		key := make(map[string]interface{})
		root["key"] = key
		key["root"] = root
		print("map cycle : ")
		println(cycle.IsCycle(root))
	}

	{
		type Cycle struct { Value int; Tail *Cycle }
		var c Cycle
		c = Cycle{42, &c}
		print("cycle : ")
		println(cycle.IsCycle(c))
	}

	{
		sli1 := make([]interface{}, 0, 10)
		sli2 := make([]interface{}, 0, 10)
		sli1 = append(sli1, sli2) // 1
		sli2 = append(sli2, sli1) // 2
		print("slice : ")
		println(cycle.IsCycle(sli1))
	}

	{
		sli1 := make([]interface{}, 0, 10)
		sli2 := make([]interface{}, 0, 10)
		sli1 = append(sli1, &sli2)
		sli2 = append(sli2, &sli1)
		print("slice pointer : ")
		println(cycle.IsCycle(sli1))
	}

	{
		var arr1 [1]interface{}
		var arr2 [1]interface{}
		arr1[0] = arr2
		arr2[0] = arr1
		print("array : ")
		println(cycle.IsCycle(arr1)) // false
	}


	{
		var arr1 [1]interface{}
		var arr2 [1]interface{}
		arr1[0] = &arr2
		arr2[0] = &arr1
		print("array pointer : ")
		println(cycle.IsCycle(arr1)) //
	}
}
