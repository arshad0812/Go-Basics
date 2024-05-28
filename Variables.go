package main
import ("fmt")

func main() {
  fmt.Println("GoLang Basics");

  //   Declaring Var and type
  var number1 int=1;
  var condition1 bool=false;	 
  var wrd1 string="Hi";
  fmt.Println(number1," ",condition1," ",wrd1);

  
  //Decalring wihtout var and type using :=sign
  number2:=2;
  condition2:=true;
  percentage:=1.1;
  wrd2:="Hello";
  fmt.Println(number2," ",condition2," ",percentage," ",wrd2);

  //Late Declaration
  var number3 int;
  number3=4;
  print(number3);

  //Decalring multiple varaibles
  var a,b,c=1,2,3;
  fmt.Println(a," ",b," ",c);
  d,e:=1.1,true;
  fmt.Println(d," ",e);


  //Cant redeclare value
  const val int=1;
  val=2;
  
  print(val)

}