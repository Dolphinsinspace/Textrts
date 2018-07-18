package main

import ("fmt"
		"time")
	
func main() {
	var input string
	for input !="quit" {
		fmt.Println("Send Dummy to mountain, forest, or river:")
		fmt.Scanln(&input)
			ptimer := time.NewTimer(1 * time.Second)
			<-ptimer.C
			fmt.Printf("Dummy has made it to %v.",input)
	}
}

