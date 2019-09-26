package main

import (
    "fmt"
    "time"
    "flag"
)

func main() {
    var timeType = flag.String("type", "standard", "help message for flag 'type'" )
    flag.Parse()
    now := time.Now()

    switch *timeType {
        case "standard":
            fmt.Println(now)
        case "duration":
            h, _ := time.ParseDuration("4h30m")
            fmt.Printf("I've got %.1f hours of work left.\n", h.Hours())
        case "kitchen":
            fmt.Println(now.Format(time.Kitchen))
    } 


}
//var formatType = flag.String("type", "default", "help message for flag 'type'")
    // flag.Parse()
    // var name string = "Jeffrey B. Daube"
    // switch *formatType {
    //     case "default":
    //         fmt.Printf("%v\n", name)
    //         return
    //     case "%#v":
    //         fmt.Printf("%#v\n", name)
    //         return
    //     case "%T":
    //         fmt.Printf("%T\n", name)
    //         return
    // }