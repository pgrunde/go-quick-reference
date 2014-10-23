package main
import "fmt"

type User struct {
    Id                  int
    FirstName, LastName string
}

func main() {
  u := User{Id:1, FirstName:"Peter", LastName:"Grunde"}
  fmt.Println(u)
}
