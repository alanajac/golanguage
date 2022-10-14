package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Planets struct {
	planetName    string
	moons         []string
	semimajoraxis float64
	eccentricity  float64
}
type Animal struct {
	Name   string `required max:"100"`
	Origin string
}
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	//three different ways of how to declare variables
	//var then name of the variable then the type of the variable
	//looks like you speak
	var i int
	i = 42
	j := 28.5
	k := "love"
	l := strconv.Itoa(i)
	var (
		actress_name    = "Elizabeth"
		actress_surname = "Reed"
		actress_place   = "California"
	)
	fmt.Println("Hello World!")
	fmt.Println(i, j, k)
	fmt.Printf("%v=%T\t%v=%T\t%v=%T\t%T\n", i, i, j, j, k, k, l)
	fmt.Printf("nome da atriz"+"%v\t"+"sobrenome da atriz: "+"%v\t"+"ela mora em:"+"%v\n", actress_name, actress_surname, actress_place)
	//boolean types
	//var n bool = false
	//fmt.Println("%v, %T\n", n, n)
	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", m, m)
	//numeric types
	//int->unspecified size-at least 32 bits. It can be more.
	//int8,int16,int32,int64(signed integer)
	//uint8,uint16.uint32
	// other operations
	a := 10 //1010
	b := 3  //0011
	fmt.Printf("results of binary operations:\n")
	//which bits are set in BOTH numbers--> result 2: 0010
	fmt.Println(a & b)
	//which bits are set in one or the other--> result 11: 1011
	fmt.Println(a | b)
	//which bits are set in one or the other, but not both--> result 9 1001
	fmt.Println(a ^ b)
	//only if neither of them is set-->result 8 0100
	fmt.Println(a &^ b)
	//bit shift
	c := 8
	fmt.Println(c << 3) //2^(3+3)
	fmt.Println(c >> 3) //2^(3-3)
	//complex numbers!
	var imagined_number complex64 = 1 + 2i
	fmt.Printf("%v,%T\n", real(imagined_number), real(imagined_number))
	fmt.Printf("%v,%T\n", real(imagined_number), real(imagined_number))

	s := "this is a string"
	fmt.Printf("%v\n", s)
	fmt.Printf("%v,%T\n", string(s[2]), s[2])
	//text variables- string and

	//CONSTANTS
	//It has to be lowered cased-internal constant or global constant. Global constant -> upper case.
	const myCOnst int = 42
	fmt.Printf("%v,%T\n", myCOnst, myCOnst)
	//we can determine constants by the data types we know, the primitives one, but not as the result of a function or something. Such as sin(1.57)
	//const test float32 = sin(1.57)
	//fmt.Printf(test)
	//we can change the type of the constant if it is shadowed.
	const test1 = 42
	var d int16 = 27
	fmt.Printf("%v,%T\n", test1+d, test1+d)
	//iota atributes sequencial value. its a counter.
	const (
		a1 = iota
		b1
		c1
	)
	const (
		a2 = iota
	)
	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", b1)
	fmt.Printf("%v\n", c1)
	fmt.Printf("%v\n", a2)

	//arrays-elements are cotiguous in memory
	grades := [3]int{97, 85, 93} //string, etc
	//we can use the following syntax too:
	gradesx := [...]int{97, 85, 93} //string, etc
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Grades: %v\n", gradesx)
	//we have the built-in lenght function
	//var students [5]string
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)
	//copying arrays-copy all array. it will be in a different memory space
	//Slices are referenced types
	//Slices are very similar to lists
	//make function
	cake := make([]int, 3, 100)
	fmt.Println(cake)
	fmt.Printf("Length: %v\n", len(cake))
	fmt.Printf("Capacity: %v\n", cap(cake))
	cookie := []int{}
	fmt.Println(cookie)
	fmt.Printf("Cookie Length: %v\n", len(cookie))
	fmt.Printf("Cookie Capacity: %v\n", cap(cookie))
	cookie = append(cookie, 1)
	fmt.Println(cookie)
	fmt.Printf("Cookie Length: %v\n", len(cookie))
	fmt.Printf("Cookie Capacity: %v\n", cap(cookie))

	//stack operations
	//how to pop elements out and add objects to the stack
	slice_bread := []int{1, 2, 3, 4, 5}
	slice_cake := slice_bread[:len(slice_bread)-1]
	slice_pie := append(slice_bread[:2], slice_bread[3:]...)
	fmt.Print("print the slice of cake")
	fmt.Println(slice_cake)
	fmt.Print("print the slice of pie")
	fmt.Println(slice_pie)
	fmt.Println(slice_bread)

	//maps
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"Recife":    1,
		"Olinda":    2,
		"Caruaru":   3,
		"Garanhuns": 4,
		"Bonito":    5,
		"Palmares":  6,
		"Jaboatão":  7,
	}
	statePopulations["Catende"] = 10
	fmt.Println(statePopulations)
	delete(statePopulations, "Catende")
	fmt.Println(statePopulations)
	//using a struct
	planet1 := Planets{
		planetName: "Neptune",
		moons: []string{
			"Triton",
			"Larissa",
			"Naiada",
			"Thalassa",
		},
		semimajoraxis: 30.1,
		eccentricity:  0.009,
	}
	planet2 := Planets{
		planetName: "Mars",
		moons: []string{
			"Phobos",
			"Deimons",
		},
		semimajoraxis: 1.524,
		eccentricity:  0.0934,
	}
	planet3:=Planets{
		planetName: "Earth",
		moons: []string{
			"Moon",
		},
		semimajoraxis: 1.0,
		eccentricity: 0.0167
	}
	fmt.Println(planet1)
	fmt.Println(planet1.moons[2])

	//we can initialize a struct:
	planet2 := struct{ name string }{name: "Mars"}
	fmt.Println(planet2)
	//Go does not have inheritance.See above with some structures detailed: Bird and Animal struct
	bird1 := Bird{
		Animal: Animal{Name: "Emu", Origin: "Australia"}, SpeedKPH: 30, CanFly: false}
	bird1.Name = "Emu"
	bird1.Origin = "Australia"
	bird1.SpeedKPH = 48

	//We can add tags
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
	//looping
	//simple loops
	//Exiting early loops
	//loopings through collections
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	counter := 0
	for ; counter < 5; counter++ {
		fmt.Println(counter)
	}

	//how to work with collections and for loops
	fmt.Printf("here a mark.")
	sx := []int{1, 2, 3}
	fmt.Println(sx)
	//for range loop
	for k, v := range s {
		fmt.Println(k, v)
	}
	for _, v := range statePopulations {
		fmt.Println(v)
	}

	//another flow controls
	//defer
	//panic
	//recover
	defer fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end")

	variavel1 := 42
	variavel2 := variavel1
	fmt.Println(variavel1, variavel2)
	variavel1 = 27
	fmt.Println(variavel1, variavel2)
	//lets crate a pointer
	var variavel3 int = 42
	var variavel4 *int = &variavel3
	fmt.Println(&variavel3, variavel4)
	fmt.Println(&variavel3, *variavel4)
	*variavel4 = 14
	fmt.Println(variavel3, *variavel4)
	variavel5 := [3]int{1, 2, 3}
	variavel6 := &variavel5[0]
	variavel7 := &variavel5[1]
	fmt.Printf("values are:\n")
	fmt.Printf("%v\t%p\t%p\t%v\t%v\n", variavel5, *variavel6, *variavel7, *variavel6, *variavel7)
	//other ways:
	var ms *myStruct
	fmt.Println(ms)
	ms = &myStruct{foo: 42}
	fmt.Println(*ms)
	//de-reference has less precedence than "." operator. So we can write:
	(*ms).foo = 45
	fmt.Println((*ms).foo)
	//we get the same behaviour as:
	fmt.Println(ms.foo)

	//calling sayMessage
	sayMessage("Hello Gorld!")
	//defer will run at the end of the main
}

type myStruct struct {
	foo int
}

func sayMessage(msg string) {
	fmt.Println(msg)
}
