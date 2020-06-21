
package main

import "fmt"

func needInt(x int) int { return x*10 +1 }
func needFolat(x float64) float64 { return x*0.1 }

func main(){
    const (
         Big   =  1 << 100 
         Small = Big >> 99 
    )

    fmt.Println(needInt(Small))
    fmt.Println(needInt(Small))
    fmt.Println(needFolat(Big))
}
