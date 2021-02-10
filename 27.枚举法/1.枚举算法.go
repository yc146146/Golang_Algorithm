package main

import "fmt"

//x+y+z = 100




func main() {

	for x:=0;x<21;x++{
		for y:=0;y<=33;y++{
			z:=100-x-y
			if z%3==0 && 5*x+3*y+z/3==100{
				fmt.Println(x,y,z)
			}
		}
	}

}
