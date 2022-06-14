/* Alta3 Research | RZFeeser
   init - order of initialization  */

package main 

import "fmt"

//var WhatIsThe = AnswerToLife() // testing a fix

func init() {
    WhatIsThe = 0
    fmt.Println("init method")
}

var WhatIsThe = AnswerToLife()

func AnswerToLife() int {
    fmt.Println("answer to life")
    return 42
}

func main() {
    if WhatIsThe == 0 {
        fmt.Println("The cake is a lie.")
    }
}

