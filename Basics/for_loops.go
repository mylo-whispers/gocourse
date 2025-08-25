package basics

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Println(index, " and ", value)
	}
}
