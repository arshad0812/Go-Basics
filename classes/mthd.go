package main
import ("fmt")


type obj struct{
	name string
	age int
}

func (obj1 *obj) method1() string{
	fmt.Println(obj1.name)
}

func main(){
    obj2:=obj{name:"Arshad",age:23}
	obj2.method1()
}