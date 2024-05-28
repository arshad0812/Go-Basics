package main


import "fmt"


func main(){
	fmt.Println("Loops");
	var arr=[]int{1,2,3};
	var map1=map[int]string{1:"arshad",2:"abi"};

	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i]);
	}


	j:=0;
	for j<len(arr){
		fmt.Println(arr[j]);
        j++;
	} 


	for key,val :=range map1{
		fmt.Println(key,val);
	}
}