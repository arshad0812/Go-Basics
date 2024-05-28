package main

import "fmt"


func main(){
	Hmap1:=map[int]string{
		1:"Arshad",
		2:"Abi",
	}

	Hmap2:=map[string]int{
		"arshad":1,
		"abi":2,
	}

	var Hmap3=map[int]int{
		1:1,
		2:2,
	}

	//Empty Map
	Hmap4:=map[int]int{};
	var Hmap5=map[int]int{};

	hmap6:=make(map[int]int);
	hmap6[0]=1
	
	fmt.Println(Hmap1,Hmap2,Hmap3,Hmap4,Hmap5,hmap6);

}