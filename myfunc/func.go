package myfunc
import "fmt"

func Locate(x int,y int) {
	fmt.Printf("\033[%d;%dH", x, y);
}
func Cls(){
	fmt.Printf("\033[2J")
}
func ArrayPrint(array []string){
	for i := 0; i < len(array); i++ {
		fmt.Printf("%s",array[i])
	}
}