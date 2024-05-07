package main
import("fmt")



func main(){

	//int to Float
	number1:=1;
	FloatNumber1:=float64(number1);
	fmt.Println(FloatNumber1);


	//Float to int
	number2:=int(FloatNumber1);
	fmt.Println(number2);
	
}