package main

import "fmt"

// "math"

// error if a variable is defined but not used
// cannot redeclare same variable in the same scope
// In go uninitialized value is 0
// using fisrt letter is uppercase then it will be export
var i int16 = 34
var num float32 = 45
var a string = "abc"
var b int

// iota is scoped to a constant block
const (
	//enumerated constants
	_ = iota + 2 //+2 has no effect on p ,q ,r and so on
	p = iota     //like counter
	q = iota
	r = iota + 4 //+4 will have effect on n1,n2,n3 only if n1 !=iota n2=! iota n3!=iota
	n1
	n2
	n3
	// { p=iota q r } will produce same result

)

type student struct {
	name      string
	course    string
	contactNo int
	subjects  []string
}

func main() {
	/*var i int = 60
	var j float32 = 78
	m := "anjali"
	j = float32(i)
	j = i //error
	var I int = 56
	a = strconv.Itoa(i)
	fmt.Println(num, i, I)
	fmt.Println(j, a)
	fmt.Printf("%v,%T\n", m, m)
	fmt.Printf("%v,%T\n", a, a)

	//primitives bool,int8 ,int16,int32,int64 ,
	// var n bool = false
	// var n uint16 = 34
	var n int16 = -23

	fmt.Printf("%v,%T\n", n, n)
	fmt.Println(b)
	*/
	//operation of same data types

	/*var num1 int = 23
	var num2 int = 45
	// fmt.Println(num1+num2, num1*num2, num2-num1, num1-num2, num2/num1, num2%num1)

	//bit operators
	num1 = 10
	num2 = 3
	fmt.Println(num1&num2, num1|num2, num1^num2, num1&^num2)
	fmt.Println(num2<<3, num1>>2) //num2 * (2^3) num1/(2^2)
	var complexNo = complex(4, 5)
	fmt.Printf("%v,%T\n", complexNo, complexNo)*/

	//strings
	/*str1 := "my name is "

	// str2 := "anjali singh"
	b := []byte(str1) //string is a collection of bytes
	fmt.Printf("%v,%T\n", b, b)
	r := 'a'
	fmt.Printf("%v,%T", r, r) //rune is   int32
	*/
	//constants created at the compile time and express assign must be constant
	// const myName string = "Anjali Singh"

	// const myConst float32 = math.Sin(1.57) //error since this will compute at the compile time hence not constant
	const i int32 = 89
	// fmt.Println(p, q, r, n1, n2, n3)

	//Arrays

	/*languages := [...]string{"cpp", "python", "Go", "javascript", "java"}
	fmt.Println(languages)
	languages[4] = "ruby"
	fmt.Println(languages, len(languages))
	*/
	//2d array
	/*var matrix [3][3]int
	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{8, 6, 9}
	fmt.Println(matrix)*/

	//slice
	/*array := []int{4, 5, 6, 2, 1, 8}
	barray := &array
	array[2] = 11
	fmt.Println(array, cap(array), len(array), barray, cap(*barray), len(*barray))
	carray := array[:] //slice of all elements
	darray := array[2:]
	karray := array[3:5] //3 inclusive 5 exclusive
	array[1] = 100
	fmt.Println(carray)
	fmt.Println(darray)
	fmt.Println(karray)
	*/
	//slice
	/*a := []int{3, 4, 1, 9, 9}
	fmt.Println(len(a), cap(a))
	a = append(a, 5)
	fmt.Println(len(a), cap(a))
	*/
	//slices via make
	/*a := make([]int, 10)      //create slice with capacity=10 and len=10
	a := make([]int, 10, 100) //slice with len=10 and capacity=100
	*/

	//MAPS AND STRUCTS
	//manipulating map at one place is going to have impact on other maps made up of that map
	marks := map[string]int{
		"math":    88,
		"english": 79,
		"social":  67,
		"science": 99,
		"hindi":   45,
	}
	// mp := make(map[string]int)
	// mp = marks
	fmt.Println(marks, len(marks))

	// fmt.Println(marks["maths"]) //if key is not present then 0
	fmt.Println(marks["math"])
	marks["humanity"] = 65

	delete(marks, "hindi")

	pop, ok := marks["social"]

	fmt.Println(pop, ok)

	//Struct
	student1 := student{
		name:      "anjali",
		course:    "CS",
		contactNo: 38080380803,
		subjects:  []string{"MTH-1", "Env", "C++", "Chem", "OS"},
	}
	fmt.Println(student1)

}
