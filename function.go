package main

import (
	"fmt"
)

func fun1(name string) string {
	fmt.Println(name)
	return "fun1 " + name
}

func fun2(arr []int)[]int {
	for i := 0; i < len(arr); i++ {
		arr[i]+=1;
		// fmt.Println(arr[i])
	}
	return arr;
}


func fmap(Hmap map[int]int)map[int]int{
	for key,val := range Hmap {
        Hmap[key]=val+1;
    }
    return Hmap;
    
}

func main() {
	var arr = []int{1, 2, 3}

	fmt.Println(fun1("Arshad"))
	var brr=fun2(arr)
	
	for i := 0; i < len(brr); i++ {
		fmt.Println(brr[i])
	}

	var hmap=map[int]int{
		1:1,
        2:2,
        3:3,
	}

	var hmap1=map[int]int{};
	hmap1=fmap(hmap);
	for key,val := range hmap1{
		fmt.Println(key,val);
	}

}
