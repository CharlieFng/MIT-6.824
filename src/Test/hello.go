package main

import ("fmt"; "math")


var sharedVar = "Shared Variable"

var (
	a = 5
	b = 6
	c = 10
)

func main(){

	fmt.Println(len("hello, world"))
	fmt.Println("hello, world"[1])
	fmt.Println("Hello " + "World")
	fmt.Println(true || false)

	var x string = "Test variable!"
	fmt.Println(x)

	var y  string
	y = "Alternative definition of variable"
	fmt.Println(y)

	var z = "complier inference" // omit type
	fmt.Println(z)

	w := 10.          // omit var and type
	fmt.Println(w)

	test()

	fmt.Println("a=",a, "b=",b, "c=",c )

	/*
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
	*/

	for i := 1; i<=10; i++{
		if i%2 == 0{
			fmt.Println(i, "even")
		}else{
			fmt.Println(i, "odd")
		}
		switch i{
		case 1 : fmt.Println("One")
		case 3 : fmt.Println("Three")
		case 5 : fmt.Println("Five")
		default: fmt.Println("Unknown number")
		}

	}

	// Array
	var array [5]int
	array[4] = 100
	fmt.Println(array)


	xx := [5]float64{98,84,34,25,75,}
	var total float64 = 0
	for _, value := range xx {
		total += value
	}
	fmt.Println(total / float64(len(xx)))

	//Slices

	slice := make([]float64, 5, 10) // 5 is current size, 10 is upper limit, for create a slice
	fmt.Println(slice[1:5])

	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := make([]int, 2)
	copy(slice3, slice1)
	fmt.Println(slice3)


	//Map
	mapp := make(map[string]int)
	mapp["Ten"] = 10
	mapp["Twenty"] = 20
	fmt.Println(mapp["Ten"])
	delete(mapp, "Ten")
	fmt.Println(mapp)
	fmt.Println(mapp["Three"] == 0) // if no exist, return zero value

	if value, ok := mapp["Twenty"]; ok{ // ok represents where lookup success or not
		fmt.Println(value, ok)
	}


	inputSlice := []float64{23,34,45,13,87}
	fmt.Println(average(inputSlice))


	r1, r2 := f3()
	fmt.Println(r1,r2)

	fmt.Println(add(1,2,3))

	xss := []int{2,3,4,5}
	fmt.Println(add(xss...))


	cs := 0
	increment := func() int {
		cs++
		return cs
	}
	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	/*
	defer func() {         // using defer to make sure the execution of recover
	    str := recover() // extract input parameter to panic
	    fmt.Println(str)
	}()
	panic ("PANIC")
	*/

	test := 5
	zero(&test)
	fmt.Println(test)
	fmt.Println(*(&test))  //& get address of variable, * to dereference pointer to original value

	xPtr := new(int) // using 'new', another way to get a pointer, it takes a type as as argument and return a pointer
	zero(xPtr)
	fmt.Println(*xPtr)


	//Test structs

	c := Circle{x:0,y:0,r:5} // c := Cigrcle{0,0,5}
	c.x, c.y = 10, 10
	fmt.Println(c.x,c.y,c.r)

	fmt.Println(circleArea(&c))

	//Test method
	fmt.Println(c.area())

	d := new(Android)
	d.Talk()   // Talk is the method of type Person, while Andriod is a Person


	e := Circle{x:5,y:5,r:24}
	fmt.Println(totalArea(&c, &e))

}

/*---------------------------------------------------------------------------------------------*/

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total/ float64(len(xs))
}


func test(){

	const x string = "Constant string"
	fmt.Println(sharedVar)

	fmt.Println(x)


}

// name the return type in function
func f2() (r int) {
	r = 1
	return
}


// multiple return values

func f3() (int, int) {
	return 5, 6
}


//variadic functions
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}


// Closure
func makeEvenGenerator() func() uint{
	i := uint(0)
	return func() (ret uint){
		ret = i
		i += 2
		return
	}
}


//Pointer
//Pointer reference a locaiton in memory where a value is stored, by using a pointer(*int) the zero
//function is able to modify the original variable
// Pointer is represented using * followed by the type of stored value
//* is also used to "dereference" pointer variables,give us access to value the pointer points to
func zero(xPtr *int){
	*xPtr = 100000000
}



//struct, used for defining fields
type Circle struct {
	x, y, r float64
}


func circleArea(c *Circle) float64{
	return math.Pi * c.r * c.r    //math.Pi * (*c).r * (*c).r
}



// Method, a special function
func (c *Circle) area() float64{
	return math.Pi * c.r * c.r
}


// Test embeded type

type Person struct {
	Name string
}
func (p *Person) Talk(){
	fmt.Println("Hi, my name is ", p.Name)
}


type Android struct {
	Person
	Model string
}


//Test interface, used for define method set
type Shape interface {
	area() float64
}

// use interface types as argumentsw to functions
func totalArea(shapes ...Shape) float64{
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}


// interfaces can also be used as fields:
type MultipleShape struct {
	shapes []Shape
}

func (m *MultipleShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

