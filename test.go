package main

import (
	"fmt"
)

// DEALING WITH FUNCTIONS

// func sayGreeting(n string) {
// 	fmt.Printf("Good Morning %v \n", n)
// }

// func sayBye(n string) {
// 	fmt.Printf("Goodbye %v \n", n)
// }

// func cycleNames(n []string, f func(string)) {
// 	for _, v := range n {
// 		f(v)
// 	}
// }

// func circleArea(r float64) float64 {
// 	return math.Pi * r * r
// }

// func getInitials(n string) (string, string) {
// 	s1 := strings.ToUpper(n)
// 	names := strings.Split(s1, " ")

// 	var initials []string
// 	for _, v := range names {
// 		initials = append(initials, v[:1])
// 	}

// 	if len(initials) > 1 {
// 		return initials[0], initials[1]
// 	}

// 	return initials[0], "_"
// }

func updateName(x string) string {
	x = "Monday"

	return x
}

// func updateMenu(y map[string]float64) {
// 	y["ice"] = 15000
// }

func main() {

	// Group A Types --- Strings, Ints, Floats, Bool, Arrays, Structs
	name := "Tifa"

	name = updateName(name)
	fmt.Println("memory address for name is:", &name)

	// Group B Types --- Maps, Slices, Functions

	// menu := map[string]float64{
	// 	"shawarma": 2000,
	// 	"ice cream": 1000,
	// }

	// updateMenu(menu)
	// fmt.Println((menu))



	// fn1, sn1 := getInitials("covenant Monday")
	// fmt.Println(fn1, sn1)

	// fn2, sn2 := getInitials("Hello world")
	// fmt.Println(fn2, sn2)

	// fn3, sn3 := getInitials("Hello")
	// fmt.Println(fn3, sn3)

	// sayGreeting("Covenant")
	// sayBye("Covenant")

	// cycleNames([]string{ "cloud", "cloudinary", "covenant"}, sayGreeting)
	// cycleNames([]string{ "cloud", "cloudinary", "covenant"}, sayBye)

	// a1 := circleArea(13)
	// a2 := circleArea(19)

	// fmt.Println(a1)
	// fmt.Println(a2)


	// fmt.Println("Hello Covenant! Welcome to Golang")

	// var NameOne string = "Covenant"
	// var NameTwo = "Monday"
	// var NameThree string


	// fmt.Println(NameOne, NameTwo, NameThree)

	// NameOne = "Governor"

	// NameThree = "Browser"
	// fmt.Println(NameOne, NameTwo, NameThree)

	// var ageOne int = 40
	// var ageTwo = 30

	// ageThree := 20

	// fmt.Println(ageOne, ageTwo, ageThree)

	// // Bits and memory


	// // Float
	// var decimalOne float32 = 5.5
	// var decimalTwo float64 = 58.912
	// fmt.Println(decimalOne, decimalTwo)

	// age := 25
	// name := "covenant"
	// // Print 
	// // fmt.Print("Hello")
	// // fmt.Print("World!")

	// fmt.Print("hello")
	// fmt.Println("Covenant Welcome to Golang")
	// // fmt.Println("My name is", name, "I am", age, "years old")

	// // formatted string
	// fmt.Printf("My name is %v and i am %v years old \n", name, age)
	// fmt.Printf("My name is %q and i am %v years old \n", name, age)
	// fmt.Printf("Age is of type %T \n", age)

	// //Sprint
	// // var firstStr = fmt.Sprint("My name is %v and i am %v years old \n", name, age)
	// // fmt.Println("The saved string is:", firstStr)


	// // Array
	// var simpleArr [3]int = [3]int{20, 90, 54}

	// // simpler
	// var simpleArray = [3]int{20, 90, 54}

	// names := [2]string{"covenant", "monday"}

	// fmt.Println(simpleArr, len(simpleArray))
	// fmt.Println(simpleArr, len(simpleArray))
	// fmt.Println(names)

	// // slices
	// numbers := []int{1, 2, 4, 6}

	// // number = append(number, 80)
	// fmt.Println(numbers)

	// //slice ranges
	// rangoOne := numbers[1:3]
	// rangeTwo := numbers[2:]
	// rangeThree := numbers[:4]

	// fmt.Println(rangoOne)
	// fmt.Println(rangeTwo)
	// fmt.Println(rangeThree)

	// rangoOne = append(rangoOne, 35)
	// fmt.Println(rangoOne)
	
	// greeting := "Hello there! Welcome to Golang"

	// fmt.Println(strings.Contains(greeting, "Hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	// fmt.Println(strings.ToUpper(greeting))

	// age := []int{2, 4, 5, 6, 2, 3, 4, 5, 46, 35, 35}

	// sort.Ints(age)
	// fmt.Println(age)

	// Loops

	// x := 0

	// for x < 5 {
	// 	fmt.Println("Value of X is:", x)
	// 	x++
	// }


	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value of I is:", i)
	// }

	// myNames := []string{"covenant", "is", "a", "boy"}

	// for i := 0; i < len(myNames); i++ {
	// 	fmt.Println("My names are:", i)
	// 	fmt.Println(myNames[i])
	// }

	// for index, value := range myNames {
	// 	fmt.Printf("the value at index %v is %v \n", index, value)
	// }

	// for _, value := range myNames {
	// 	fmt.Printf("the value is %v \n", value)
	// }


	// BOOLEANS AND CONDITIONALS
	// age := 45

	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 45)
	// fmt.Println(age != 50)

	// if age < 30 {
	// 	fmt.Println("Age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("Age is less than 40")
	// } else {
	// fmt.Println("Age is not less than 45")
	// }


	// myNames := []string{"covenant", "is", "a", "boy"}

	// for index, value := range myNames {
	// 	if index == 1 {
	// 		fmt.Println("Continuing at position", index)
	// 		continue
	// 	}
	// 	if index > 2 {
	// 		fmt.Println("Breaking at position", index)
	// 		break
	// 	}
	// 	fmt.Printf("The value at %v is %v \n", index, value)
	// }


	// MAPS
	// menu := map[string]float64{
	// 	"soup": 500,
	// 	"jellof": 800,
	// 	"fried rice": 900,
	// 	"mtn rice": 300,
	// }

	// fmt.Println(menu)
	// // to get a single value
	// fmt.Println(menu["jellof"])

	// // loop through maps
	// for k, v := range menu {
	// 	fmt.Println(k, "-", v)
	// }

	// // Using ints as keys and string as value
	// phonebook := map[int]string{
	// 	1234567: "Favour",
	// 	2345678: "Miracle",
	// 	3456789: "Testimony",
	// 	4567890: "Covenant",
	// }
	// fmt.Println(phonebook)
	// fmt.Println(phonebook[1234567])

}