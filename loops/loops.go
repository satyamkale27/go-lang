package loops

import "fmt"

func loop() {

	// 	i :=1
	// 	for i <= 3{
	//    fmt.Println(i)
	//     i = i + 1
	//	}

	// for j:=0; j<3;j++{
	// fmt.Println(j)
	// }

	// for i:= range 3 {
	// 	fmt.Println(i);    // prints 0,1,2 //
	// }

	// for {
	// 	fmt.Println("break")    // break condition breaks the loop, without break loop will be infinite //
	// 	break
	// }

	for n := range 6 {
		if n%2 == 0 {
			continue // continue is used to skip the current iteration of the loop and proceed to the next iteration. // bypass the remaining code in the current iteration //
		}
		fmt.Println(n)
	}

}
