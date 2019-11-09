package file

import (
	"fmt"
	"os"
)

func WriteStr()  {
	f, err := os.Create("test.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(l, "bytes write ok")

}

func WriteByte()  {
	f, err := os.Create("bytes")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	d2 := []byte{104,101,108,111,32,119,111,114,108,100}

	n2,err := f.Write(d2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n2, "bytes write ok")
}

func WriteLine()  {
	f, err := os.Create("line.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	d2 := []string{"welcome to the world", "it is easy"}

	for _,v := range d2 {
		_,err := fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("line write ok")
}

func WriteAppend()  {
	f, err := os.OpenFile("line.txt", os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	newLine := "file handling is easy"


	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Println("append write ok")
}

