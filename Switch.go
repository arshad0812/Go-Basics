package main
import "fmt"


func main(){
	color:="red";
	switch color{

	case "red":
		fmt.Println("Its red");
		break;

	case "green":
		fmt.Println("Its green");
		break;
		
	default:
		fmt.Println("Its not red or green");
		break;
	}
}