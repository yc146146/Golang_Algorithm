package main

import "fmt"

func LargeNumberMul(a, b string)string{
	var result string
	a = Rev(a)
	b = Rev(b)
	c := make([]byte, len(a)+len(b))


	for i:=0;i<len(a);i++{
		for j:=0;j<len(b);j++{
			//叠加计算
			c[i+j]+=(a[i]-'0')*(b[j]-'0')
		}
	}
	var plus byte = 0
	for i:=0;i<len(c);i++{
		if c[i]==0{
			break
		}
		temp := c[i]+plus
		plus=0
		if temp >9{
			plus = temp/10
			result += string(temp-plus*10+'0')
		}else{
			result += string(temp+'0')
		}
	}
	return Rev(result)
}


func Rev(mystr string)string{

	//开辟内存
	newstring := make([]uint8, len(mystr), len(mystr))

	for i:=0;i<len(newstring);i++{
		newstring[i] = mystr[i]
	}

	for i,j := 0,len(newstring)-1;i<j;i,j=i+1,j-1{
		newstring[i], newstring[j] = newstring[j],newstring[i]
	}
	return string(newstring)


}


func main() {
	var str1 = "123"
	var str2 = "789"

	fmt.Println(Rev(str1))
	fmt.Println(Rev(str2))

	//97047
	fmt.Println(LargeNumberMul(str1,str2))
}
