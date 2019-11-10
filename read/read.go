package read

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gobuffalo/packr"
	"io/ioutil"
	"log"
	"os"
)

func ReadAll()  {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("file reading error", err)
		return
	}
	fmt.Println("content of file:", string(data))
}

func Packr()  {
	box := packr.NewBox("../data")
	data,error := box.FindString("hello.txt")
	if error != nil {
		fmt.Println("packr reading error", error)
		return
	}
	fmt.Println("contents of file:", data)
}

func Buffio()  {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for   {
		_, err := r.Read(b)
		if err != nil {
			fmt.Println("error reading file:", err)
			break
		}
		fmt.Println(string(b))
	}
}

func Line()  {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}