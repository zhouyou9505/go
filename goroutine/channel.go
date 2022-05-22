package goroutine

import "fmt"

func sum(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

func invokeChannel() {

	arr := []int{1, 2, 3, 4, 5}
	c := make(chan int)

	go sum(arr[:len(arr)/2], c)
	go sum(arr[len(arr)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

}

func MyChannel() {

	//通道是一个引用类型的数据 0xc0000240c0
	myChan := make(chan int)

	fmt.Println(myChan)

}

func Child2ParentChanTest() {

	var mychan chan bool
	mychan = make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("innner print...")
		}

		mychan <- true
	}()

	data := <-mychan

	fmt.Println("main over....", data)
}

func Parent2ChildChanTest() {

	var mychan chan bool
	mychan = make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("print ...")
		}

		data := <-mychan
		fmt.Println("data---->", data)
	}()

	mychan <- true
	fmt.Println("main over...")
}
