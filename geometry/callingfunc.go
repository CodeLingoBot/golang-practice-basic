package main
var ac int

func main(){
	ac = 7
	println(ac)
	f()
}

func f() {
	ac:= 8
	println(ac)
	g()
}

func g() {
	println(ac)
}

