package main


import "fmt"

import "time"


func main() {

    t,_ := time.Parse(time.UnixDate,"Mon Jan 14 21:50:45 EST 2013")

    fmt.Println(t.Format(time.RFC3339))  // prints time as Z


    t2,_:=time.Parse(time.RFC3339,t.Format(time.RFC3339))

    fmt.Println(t2.Format(time.UnixDate)) // prints time as UTC

}