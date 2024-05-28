package main
import("fmt")


type obj struct{
	name string
	surname string
}

func main(){
	var obj1=obj{name:"arshad",surname:"ahamed"};
	// obj1.name = "ali";
	// obj1.surname = "ahamed";
	fmt.Println(obj1.name, obj1.surname);
}