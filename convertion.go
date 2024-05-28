package main
import("fmt"
"strconv")


func main(){
	fmt.Println("Convertion")
	//Str to int
	var wrd string="123"
	num,_:=strconv.Atoi(wrd)


	num1:=123
	wrd1:=strconv.Itoa(num1)

	var flv=1.2

	num2:=float64(flv)
	num3:=int(flv)


	fmt.Println(num,wrd1,num2,num3)
	
}