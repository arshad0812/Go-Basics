package main
import "fmt"




func main(){
	var ptr *string;
	var wrd string="Hi";
	ptr=&wrd;
	fmt.Println(*ptr);
}