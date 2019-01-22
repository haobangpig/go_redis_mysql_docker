package main
import "fmt"
func main(){
  var a string
  a = "a"
  fmt.Println(a)
  b := true
  fmt.Println(b)
  var c = 123
  fmt.Println(c)
  const (
    one = 111
    twaao = "aaa"
    therr = 123
  )
  fmt.Println(one)
  fmt.Println(twaao)
  fmt.Println(therr)
  a = "changed a"
  fmt.Println(a)
}
