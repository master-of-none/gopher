package loops

import "fmt"

func loop1() {
	sum := 0

	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Printf("%d", sum)

}
