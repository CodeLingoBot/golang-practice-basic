# Introducing Go Note

# Type

Go is a statically typed programming language. This means that variables always have a specific type and that type cannot change.

### Integer

Go’s integer types are uint8, uint16, uint32, uint64, int8, int16, int32, and int64.

### Floating-Point Numbers

- Floating-point numbers are inexact.
- Floating-point numbers have a certain size (32 bit or 64 bit). Using a larger-sized floating-point number increases its precision.
- In addition to numbers, there are several other values that can be represented: “not a number” (NaN, for things like 0/0) and positive and negative infinity (+∞ and −∞).

### Defining Multiple Variables

    var (
    	a = 5
    	b = 10
    	c = 15
    )

### Arrays

An array is a numbered sequence of elements of a single type with a **fixed length**. 

Rarely see arrays used directly in Go code.

    var x [5]int

### Slices

A slice is a segment of an array. Like arrays, slices are indexable and have a length. Unlike arrays, this length is allowed to change.

    var x []float64

Create a slice, use the built-in make function:

    x := make([]float64, 5)

### append

    slice1 := []int{1,2,3}
    slice2 := append(slice1, 4, 5)

### copy

Copy takes two arguments: dst and src. All of the entries in src are copied into dst overwriting whatever is there. If the lengths of the two slices are not the same, the smaller of the two will be used.

    slice1 := []int{1,2,3}
    slice2 := make([]int, 2)
    copy(slice2, slice1)

After running this program slice1 has [1,2,3] and slice2 has [1,2]. The contents of slice1 are copied into slice2, but because slice2 has room for only two elements, only the first two elements of slice1 are copied.

### Maps

A map is an unordered collection of key-value pairs (maps are also sometimes called associative arrays, hash tables, or dictionaries). Maps are used to look up a value by its associated key.

    var x map[string]int

initialized before they can be used.

    x := make(map[string]int)
    x["key"] = 10
    fmt.Println(x["key"])

### Variadic Functions

    func add(args ...int) int { 
    total := 0
    	for_,v:=rangeargs{ total += v
    }
    	return total 
    }

    func main() {
    	xs := []int{1,2,3} 
    	fmt.Println(add(xs...))
    }

### Closure

It is possible to create functions inside of functions.

    func main() {
    	add := func(x, y int) int {
    					returnx+y 
    				}
            fmt.Println(add(1,1))
        }

One way to use closure is by writing a function that returns another function, which when called, can generate a sequence of numbers.

    func makeEvenGenerator() func() uint { 
    	i := uint(0)
    	return func() (ret uint) { 
    				ret=i
    				i+=2
    		return
    }}
    func main() {
    	nextEven := makeEvenGenerator() 
    	fmt.Println(nextEven()) // 0 
    	fmt.Println(nextEven()) // 2 
    	fmt.Println(nextEven()) // 4
    }

### Recursion

    func factorial(x uint) uint { 
    	if x == 0{
    		return 1 
    }
    	return x * factorial(x-1) 
    }

### Pointers

    func zero(xPtr *int) { 
    	*xPtr = 0
    }
    func main() {
    	x:=5
    	zero(&x)
    	fmt.Println(x) // x is 0
    }

### Structs

    type Circle struct {
    	x float64
    	y float64
    	r float64
    }

    type Circle struct {
    	x, y, r float64
    }

### Fields

Arguments are always copied in Go. If we attempted to modify one of the fields inside of the circle Area function, it would not modify the original variable. Because of this, we would typically write the function using a pointer to the Circle:

    func circleArea(c *Circle) float64 {
    	return math.Pi * c.r*c.r
    }

    c := Circle{0, 0, 5}
    fmt.Println(circleArea(&c))

### Methods

    type Rectangle struct {
    	x1, y1, x2, y2 float64
    }
    func (r *Rectangle) area() float64 {
    	l := distance(r.x1, r.y1, r.x1, r.y2)
    	w := distance(r.x1, r.y1, r.x2, r.y1)
    	return l * w
    }

### Embedded Types

    type Android struct {
    	Person Person
    	Model string
    }

We use the type (Person) and don’t give it a name. When defined this way, the Person struct can be accessed using the type name:

    a := new(Android)
    a.Person.Talk()

But we can also call any Person methods directly on the Android:

    a := new(Android)
    a.Talk()

### Interfaces

    type Shape interface { 
    	area() float64
    }

Define a method set. A method set is a list of methods that a type must have in order to implement the interface.

Can’t write a function that contains two variadic parameters.

Nothing additional is required to implement an interface (there is no implements or extends keyword). It’s sufficient to merely have a method with the same name and signature.

Interfaces can also be used as fields.

    type MultiShape struct { 
    	shapes []Shape
    }
    
    multiShape := MultiShape{
            shapes: []Shape{
                Circle{0, 0, 5},
                Rectangle{0, 0, 10, 10},
            },
    }

### Goroutines

A goroutine is a function that is capable of running concurrently with other functions.

    func f(n int) { fori:=0;i<10;i++{
                fmt.Println(n, ":", i)
            }
    }
    func main() { go f(0)
    var input string
            fmt.Scanln(&input)
        }

### Channels

Channels provide a way for two goroutines to communicate with each other and synchronize their execution.

    func pinger(c chan string) { 
    	for i:=0; ;i++{
    		c <- "ping" 
    	}
    }
    func printer(c chan string) { 
    	for {
    		msg:=<-c 
    		fmt.Println(msg) 
    		time.Sleep(time.Second * 1)
    	} 
    }
    func main() {
    	var c chan 
    	string = make(chan string)
    	go pinger(c) 
    	go printer(c)
    	var input string
      fmt.Scanln(&input)
    }

    func thoaiu(th [5]string, thoachan chan<- string) {
    	for _, item := range th {
    		thoachan <- item
    	}
    }
    
    func thoaiu(th [5]string, thoachan <-chan<- string) {
    
    }
    
    <-chan : chanel dispatch
    chan<- receive chanel

The left arrow operator (<-) is used to send and receive messages on the channel.

msg := <- c means receive a message and store it in msg.

The fmt line could also have been written like fmt.Println(<-c).

### Channel Direction

We can specify a direction on a channel type, thus restricting it to either sending or receiving.

    func pinger(c chan<- string) // Now pinger is only allowed to send to c.

    func printer(c <-chan string)

### Select

Go has a special statement called select that works like a switch but for channels

    c1 := make(chan string) 
    c2 := make(chan string)
    go func() { 
    	for {
    		c1 <- "from 1"
        time.Sleep(time.Second * 2)
    }
    }()
    go func() { 
    	for {
    		c2 <- "from 2"
        time.Sleep(time.Second * 3)
    }
    }()
    go func() { 
    	for {
    		select {
    			case msg1 := <- c1:
    				fmt.Println(msg1) 
    			case msg2 := <- c2:
            fmt.Println(msg2)
    		}
    	} 
    }()

### Buffered Channels

    c := make(chan int, 1)

This creates a buffered channel with a capacity of 1. Normally, channels are synchro‐ nous; both sides of the channel will wait until the other side is ready. A buffered chan‐ nel is asynchronous; sending or receiving a message will not wait unless the channel is already full. If the channel is full, then sending will wait until there is room for at least one more int.
