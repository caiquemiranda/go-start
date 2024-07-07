package main

import "fmt"

func main(){
	oneNumber(3)
	oneNumber(4)
	manyNumbers(10)
	}

func oneNumber(n int){
	
	if n%2 == 0 {
		fmt.Println("Number is Par")
	} else {
		fmt.Println("Number is Impar")
	}

}

func manyNumbers(n int){
	for i:=1; i<=n; i++ {
        if i%2 == 0 {
            fmt.Println(i, "is Par")
        } else {
            fmt.Println(i, "is Impar")
        }
    }
}