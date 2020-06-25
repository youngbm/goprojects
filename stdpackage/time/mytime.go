package main

import (
    "fmt"
    "time"
)

func test123() string {
    return fmt.Sprintf("I am in test123")
    
}

func main() {
    fmt.Println(time.Now())
    fmt.Println(time.Now().Year())
    fmt.Println(time.Now().Month())
    fmt.Println(time.Now().Day())
    fmt.Println(time.Now().Hour())
    fmt.Println(time.Now().Minute())
    fmt.Println(time.Now().Second()) 
    fmt.Println(time.Now().Unix()) 
    fmt.Println(time.Now().UnixNano()) 
    fmt.Println(time.Now().Format("2006-01-02 15:04:00.000000"))
    fmt.Println(time.Now().Location())
    // weekday
    var wd time.Weekday 
    wd = 0
    fmt.Printf("%s", wd.String())
    
    // parse
    // longForm shows by example how the reference time would be represented in
    // the desired layout.
    const longForm = "201606060404"
    //const longForm = "Jan 2, 2006 at 3:04pm (MST)"
    t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
    fmt.Println(t)
    // shortForm is another way the reference time would be represented
    // in the desired layout; it has no time zone present.
    // Note: without explicit zone, returns time in UTC.
    const shortForm = "2006-Jan-02"
    t, _ = time.Parse(shortForm, "2013-Feb-03")
    fmt.Println(t)

    //tick
    c := time.Tick(1 * time.Second)
    for now := range c {
        fmt.Printf("%v %s\n", now, test123())
    }

}
