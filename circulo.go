package main

import "fmt" //format

func main(){
	var r float32

	//a notação & é usada para passagem de referência
	fmt.Scan(&r)

	area := calcularArea(r)
	
	//formatando e mostrando a string formatada
	fmt.Printf("A=%.4f\n", area)

}