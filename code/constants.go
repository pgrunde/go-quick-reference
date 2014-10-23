package main
import "fmt"

const (
    foo = "A"
    bar = 'A' // any guesses what this will be?
    bin = 2
)

func main() {
    fmt.Printf("%T %v\n", foo, foo)
    fmt.Printf("%T %v\n", bar, bar)
    fmt.Printf("%T %v\n", bin, bin)

    // And for fun
    fmt.Printf("%T %v\n", bar, string(bar))
}
