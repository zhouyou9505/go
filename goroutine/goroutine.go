package goroutine

import (
	"fmt"
	"time"
)

func Say(str string) {

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str)
	}

}

func SayHi() {
	fmt.Println("say hi")
}


