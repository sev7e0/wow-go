package main

import "fmt"

type Test struct {
	a int
	b int
}

func main() {

	i, j := 42,2001
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j

	*p = *p / 74

	fmt.Println(j)

	fmt.Println(Test{1,2})
	test := Test{1, 2}
	fmt.Println(test.a)

}