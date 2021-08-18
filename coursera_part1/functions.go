package main

import "fmt"

func deferTest(){
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println("panic happened: ",err)
		}
	}()
	fmt.Println("Something cool...")
	panic("something bad happened!")
	return
}

func main(){
	printer:=func(in string){
		fmt.Println("Hello ",in)
	}

	printer("Anya")

	type strFunc func(string)

	//функция принимает коллбэк
	worker:=func(callback strFunc){
		callback("as callback")
	}
	worker(printer)

	//замыкания

	prefixer:=func(prefix string)strFunc{
		return func(in string){
			fmt.Printf("[%s] %s\n",prefix,in)
		}
	}
	successLoger:=prefixer("SUCCESS")
	successLoger("expected behavior")
	deferTest()
}
