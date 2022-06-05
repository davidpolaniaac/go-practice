package main

import (
	"fmt"
	"log"
	"project/channels"
	"project/echo"
	"project/interfaces"
	"project/model"
	"project/routines"
	"strconv"
)

func loops() {
	// loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")
	counter := 10
	for counter > 0 {
		fmt.Println(counter)
		counter--
	}
}

func prints() {
	fmt.Println("Hello world")
	fmt.Println("Variables")
}

func variables() {
	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("pi:", pi)
	fmt.Println("p2:", pi2)

	name := "David"
	var lastname string = "Polania"
	var username = "davidpolania"

	fmt.Println(name, lastname)
	fmt.Println(username)

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)
}

func operations() {
	x := 10
	y := 10
	fmt.Println("suma", x+y)
	fmt.Println("multi", x*y)
	fmt.Println("divi", x/y)
	fmt.Println("rest", x-y)
	fmt.Println("module", x%y)
	x++
	fmt.Println("incremental", x)
	x--
	fmt.Println("decremental", x)
}

func libFmt() {
	fmt.Printf("%s %s\n", "david", "polania")
	fmt.Printf("%v %s\n", "david", "polania")

	message := fmt.Sprintf("%s %s\n", "david", "polania")
	fmt.Println(message)

	fmt.Printf("%v %T\n", "Tipo", "polania")
}

func greeting(message string) {
	fmt.Println(message)
}

func sum(a, b int) int {
	return a + b
}

func division(a int, b int) (c int, d int) {
	valueA := a / b
	valueB := a % b
	return valueA, valueB
}

func functions() {
	greeting("Hello there !")
	result := sum(10, 10)
	fmt.Println(result)

	result, modulo := division(10, 2)
	fmt.Println(result, modulo)

	value, _ := division(10, 2)
	fmt.Println(value)
}
func conditionals() {

	value1 := 1
	value2 := 2

	if value1 == 1 && value2 == 3 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}

func stringToNumber() {
	value, err := strconv.Atoi("asdas")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value: ", value)
}

func switches() {

	//even and odd
	switch a := 4 % 2; a {
	case 0:
		fmt.Println("It is even")
	default:
		fmt.Println("It is odd")
	}

	// major, minor, middle
	value := 200
	switch {
	case value > 100:
		fmt.Println("It is major")
	case value < 0:
		fmt.Println("It is minor")
	default:
		fmt.Println("It is middle")
	}
}

func keywords() {

	defer fmt.Println("Close")

	for i := 0; i < 10; i++ {
		if i == 2 {
			fmt.Println("continue")
			continue
		}

		if i == 8 {
			fmt.Println("break")
			break
		}

		fmt.Println("i:", i)

	}
}

func arraysAndSlides() {

	var array [4]int
	array[0] = 1
	array[1] = 2

	fmt.Println(array, len(array), cap(array))

	slice := []int{0, 1, 2, 3, 4}

	fmt.Println(slice, len(slice), cap(slice))

	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[3:])

	slice = append(slice, 5)

	fmt.Println(slice, len(slice), cap(slice))

	newSlice := []int{6, 7, 8}

	fmt.Println(newSlice, len(newSlice), cap(newSlice))

	slice = append(slice, newSlice...)

	fmt.Println(slice, len(slice), cap(slice))

}

func arrayLoop() {
	slice := []string{"hello", "!", "how", "are", "you", "?"}

	for i, v := range slice {
		fmt.Println(i, v)
	}
}

func maps() {
	m := make(map[string]int)

	m["joseph"] = 18
	m["daniel"] = 23

	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i, v)
	}

	fmt.Println(m["daniel"])

	value, ok := m["test"]
	fmt.Println(value, ok)
}

func structs() {
	myCar := model.Car{Brand: "Ford", Year: 2022}
	fmt.Println(myCar)
	var yourCar model.Car
	yourCar.Brand = "Toyota"
	fmt.Println(yourCar)
}

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) printBrand() {
	fmt.Println("brand:", myPc.brand)
}

func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

func (myPc pc) String() string {
	return fmt.Sprintf("RAM: %d - DISK: %d - BRAND: %s", myPc.ram, myPc.disk, myPc.brand)
}

func pointers() {

	a := 50
	b := &a
	fmt.Println(a, b)
	fmt.Println(*b)
	*b = 100
	fmt.Println(a, b)

	myPc := pc{brand: "msi", ram: 12, disk: 256}
	fmt.Println(myPc)
	myPc.printBrand()
	myPc.duplicateRam()
	fmt.Println(myPc)
}

func stringers() {
	myPc := pc{brand: "msi", ram: 12, disk: 256}
	fmt.Println(myPc)
}

func main() {

	//variables()
	//prints()
	//operations()
	//libFmt()
	//functions()
	//loops()
	//conditionals()
	//stringToNumber()
	//switches()
	//keywords()
	//arraysAndSlides()
	//arrayLoop()
	//maps()
	//structs()
	//pointers()
	//stringers()
	interfaces.Figures()
	routines.Run()
	channels.Run()
	channels.Multi()
	echo.Serve()

}
