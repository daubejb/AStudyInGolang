package main

import "fmt"
import "flag"

func main() {
    var formatType = flag.String("type", "default", "help message for flag 'type'")
    flag.Parse()
    var name string = "Jeffrey B. Daube"
    switch *formatType {
        case "default":
            fmt.Printf("%v\n", name)
            return
        case "%#v":
            fmt.Printf("%#v\n", name)
            return
        case "%T":
            fmt.Printf("%T\n", name)
            return
    }
}