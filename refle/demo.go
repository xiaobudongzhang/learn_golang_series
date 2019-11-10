package refle

import (
	"fmt"
	"reflect"
)

type order struct {
	orderId int
	customerId int
}

func Demo()  {
	i := 10
	fmt.Printf("%d %T", i, i)
}
func Demo2()  {
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)

	b := "Na"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)
}
func createQuery(q interface{})   {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()

	fmt.Println("type", t)
	fmt.Println("value", v)
	fmt.Println("kind", k)
}

func createQuery2(q interface{})   {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("num of fields", v.NumField())

		for i := 0;i < v.NumField();i++  {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}

func createQuery3(q interface{})   {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values (", t)
		v := reflect.ValueOf(q)
		for i := 0; i< v.NumField();i++  {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("unsupported type")
				return
			}
		}
		 query = fmt.Sprintf("%s)", query)
		 fmt.Println(query)
		return
	}
	fmt.Println("unsuported type")
}
func DemoSql()  {
	o := order{
		orderId:      456,
		customerId: 56,
	}
	createQuery3(o)


	i := 90
	createQuery3(i)
}
