package main

import "fmt"



func main(){

	//Intializing
	var arr1=[]int{1,2,3};
	fmt.Println(arr1);

	arr2:=[]int{1,2,3};
	fmt.Println(arr2);

	//Appending
	arr1=append(arr1,4);
	fmt.Println(arr1);

	//Indexing
	fmt.Println(arr1[0]);


	//Slicing
	arr3:=arr1[0:3];
	fmt.Println(arr3);


    //empty arr
	var arr4 [5]int;
	arr4[0]=1;
	fmt.Println(arr4);

	arr5:=make([]int,5);
	fmt.Println(arr5);


	
}