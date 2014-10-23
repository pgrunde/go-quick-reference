package main
import "fmt"

const (
	Apple int = iota
	Orange
	Banana
)

func main() {
	fmt.Println(Apple)
	fmt.Println(Orange)
	fmt.Println(Banana)
}
