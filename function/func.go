package function

import "fmt"

func Demo()  {
	a := func() {
		fmt.Println("hello world ")
	}
	a()
	fmt.Printf("%T\n", a)

	func(n string){
		fmt.Println("welocome", n)
	}("go")
}

type add func(a int, b int) int

func TypeFunc()  {
	var a add = func(a int, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println("Sum", s)
}

func simpleFunc(a func(a, b int) int)  {
	fmt.Println(a(60, 7))
}

func SimpleFunc()  {
	f := func(a, b int) int {
		return a + b
	}
	simpleFunc(f)
}

func simplereturn() func(a, b int) int  {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func Simplereturn()  {
	s := simplereturn()
	fmt.Println(s(60 ,7 ))
}

func Closure()  {
	a := 5
	func() {
		fmt.Println("a=", a)
	}()
}

func appendStr() func(string) string {
	t := "hello"
	c := func(b string) string{
		t = t + " " +b
		return t
	}
	return c
}

func AppendStr()  {
	a := appendStr()
	b := appendStr()

	fmt.Println(a("World"))
	fmt.Println(b("everyone"))

	fmt.Println(a("go"))
	fmt.Println(b("!"))
}