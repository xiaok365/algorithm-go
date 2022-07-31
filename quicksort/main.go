package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var Data []int

func main() {

	Input("a.input")

	for k, v := range Data {
		fmt.Printf("k=%d, v=%d\n", k, v)
	}

	Sort(0, len(Data)-1)

	for k, v := range Data {
		fmt.Printf("sorted: k=%d, v=%d\n", k, v)
	}
}

//read data from file
func Input(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("os.Open(): failed:%s", err.Error())
	}
	defer f.Close()

	Data = make([]int, 0)

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		//fmt.Println(line)
		var s int64
		if s, err = strconv.ParseInt(line[0:len(line)-1], 10, 32); err != nil {
			fmt.Printf("strconv(): failed:%s, %s\n", line, err.Error())
		}

		//fmt.Println(s)
		Data = append(Data, int(s))
	}
}

//sort the data
func Sort(l, r int) {
	if l >= r {
		return
	}

	key := Data[l]
	i, j := l, r

	for i < j {
		for i < j && Data[j] >= key {
			j--
		}
		Data[i] = Data[j]

		for i < j && Data[i] <= key {
			i++
		}
		Data[j] = Data[i]
	}

	Data[i] = key
	Sort(l, i-1)
	Sort(i+1, r)
}
