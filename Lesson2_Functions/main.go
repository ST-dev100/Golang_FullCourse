package main


import "fmt"
//multiple argument with the same type we can use(x,y int)
func first(x int,y int) string{
	
	return "d"
}
//return multiple values
func first2(x int,y int) (int,int){
	
	return x,y
}
//named return values
func first3(x int,y int) (c int,b int){
	
	return 
}
func main(){
	fmt.Println("ff")
	fmt.Println(first(10,2));
	fmt.Println(first2(10,2));
	fmt.Println(first3(10,2));
}