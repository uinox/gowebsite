package main

import (
 "fmt"
 "net/http"
 "io/ioutil"
)

func main(){
   fmt.Println("aa")
   for i := 0; i < 10; i++ {
     geturl()
   }
}
func geturl(){
 url := "http://www.tmubei.com/home"
 for i := 0; i < 1000; i++ {
   var a = i
   fmt.Println(a)
   resp, err := http.Get(url)
   if err != nil {
      fmt.Println("err")
   }
   defer resp.Body.Close()
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
     fmt.Println("err1")
   }
   fmt.Printf("%s",body)
 }
}
