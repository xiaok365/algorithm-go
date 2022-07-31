package main

import (
	"collections/collections"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

type Item struct {
	Id   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

func TestList() {
	// 测试int
	intData := []int{1, 2, 3, 4, 5, 6, 7, 8}
	intList := collections.NewList(collections.Int2Inf(intData)).
		Filter(func(i interface{}) bool {
			return i.(int)%2 == 0
		}).
		Filter(
			func(i interface{}) bool {
				return i.(int) > 4
			}).CollectToList()

	for _, v := range intList {
		fmt.Println(v)
	}

	// 测试string
	strData := []string{"aa", "bb", "cc", "ddd"}
	strList := collections.NewList(collections.Str2Inf(strData)).
		Filter(func(i interface{}) bool {
			s := i.(string)
			return len(s) < 3
		}).
		Filter(
			func(i interface{}) bool {
				s := i.(string)
				return s == "aa" || s == "bb"
			}).CollectToList()

	for _, v := range strList {
		fmt.Println(v)
	}

	t := 2
	fmt.Println(reflect.TypeOf(t))

	// 测试item
	itemData := make([]interface{}, 0)
	itemData = append(itemData, Item{
		Id:   1,
		Name: "kyle",
	})
	itemData = append(itemData, Item{
		Id:   2,
		Name: "wilson",
	})
	itemData = append(itemData, Item{
		Id:   3,
		Name: "kyle1",
	})
	itemData = append(itemData, Item{
		Id:   4,
		Name: "kyle2",
	})

	itemList := collections.NewList(itemData).
		Filter(func(i interface{}) bool {
			return i.(Item).Id > 0
		}).
		Map(func(i interface{}) interface{} {
			return i.(Item).Name
		}).CollectToList()

	fmt.Println(fmt.Sprintf("itemList=%+v", itemList))

	itemMap := collections.NewList(itemData).
		Filter(func(i interface{}) bool {
			return i.(Item).Id > 1
		}).
		Filter(func(i interface{}) bool {
			return strings.Contains(i.(Item).Name, "kyle")
		}).
		CollectToMap(func(i interface{}) interface{} {
			return i.(Item).Id
		}, collections.Identity())

	fmt.Println(fmt.Sprintf("itemMap=%+v", itemMap))
}

func TestTreeMap() {

	cmp := func(i1 interface{}, i2 interface{}) int {
		return i1.(int) - i2.(int)
	}
	treeMap := collections.NewTreeMap(cmp)

	//for i := 1; i < 100; i++ {
	//	treeMap.Insert(i, i)
	//}

	a := []int{12, 6, 8, 14, 11, 2}
	for _, v := range a {
		treeMap.Insert(v, v)
	}

	fmt.Println("height=", treeMap.Height())

	treeMap.PrintTree()

	for i := 0; i < 20; i++ {
		p := treeMap.FindNextMin(i)
		p1 := treeMap.FindNextMax(i)
		fmt.Println(fmt.Sprintf("i=%d, front=%d, end=%d", i, p.Key, p1.Key))
	}

	//for i := 1; i < 100; i++ {
	//	p := treeMap.Find(i)
	//	if p != nil {
	//		fmt.Println(fmt.Sprintf("node=%+v", p))
	//	}
	//}
	//
	//for i := 1; i < 100; i++ {
	//	treeMap.Remove(i)
	//	fmt.Println("height=", treeMap.Height())
	//
	//}
}

func TestPriorityQueue() {

	cmp := func(a, b interface{}) int {
		return a.(int) - b.(int)
	}

	rand.Seed(time.Now().UnixNano())
	queue := collections.NewPriorityQueue(100, collections.MAX_HEAP, cmp)
	for i := 0; i < 10; i++ {
		queue.Push(rand.Intn(100))
	}

	for !queue.IsEmpty() {
		fmt.Println(queue.Pop())
	}

}

func TestHash() {

	cmp := func(i1 interface{}, i2 interface{}) int {
		return int(i1.(uint) - i2.(uint))
	}
	treeMap := collections.NewTreeMap(cmp)

	ketama := collections.NewKetamaHash()
	ip := "192.168.0.1"
	for i := 0; i < 16; i++ {
		addr := fmt.Sprintf("%s#%d", ip, i)
		fmt.Println(addr)
		fmt.Println(ketama.GetHash(addr))
		treeMap.Insert(ketama.GetHash(addr), ip)
	}

	for i := 0; i < 10; i++ {
		addr := fmt.Sprintf("%s#%d", ip, i)
		fmt.Println(addr)
		fmt.Println(ketama.GetHash(addr))
		treeMap.Remove(ketama.GetHash(addr))
	}

	fmt.Println(treeMap.Height())

	treeMap.PrintTree()

	fmt.Println("+++")
	for i := 0; i < 10; i++ {
		customerId := "asdfawefasdfasdfasdffsdaf"
		fmt.Println(treeMap.FindNextMin(ketama.GetHash(customerId)))
	}
}

func main() {

	//TestList()
	//TestTreeMap()
	TestPriorityQueue()
	//TestHash()
}
