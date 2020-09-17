package demo

import "fmt"

func main() {

	//4 collection types.

	//1. Arrays
	fmt.Println("Arrays _________________")
	//Method one for array Declaration
	array1 := [3]int64{1, 2, 3}
	fmt.Println(array1)

	var array2 [3]int
	array2[0] = 1
	array2[1] = 2
	array2[2] = 3
	fmt.Println(array2)

	//2. Slices
	//Slices - Demo One
	/*Slices can be extended from arrays as demonstrated in the code snippet below.
	Slices are dynamic
	*/
	fmt.Println("Slice _________________")
	slice1 := array1[:]
	fmt.Println(slice1)
	/*Any change made to the array will also be reflected in the splice
	for this case because slice1 is acting as a pointer to array1
	*/
	array1[2] = 67
	slice1[1] = 88
	//Result should be the same
	fmt.Println(array1, slice1)
	//End of Slices part one.

	//Slice Demo 2
	// This demo shows how to create a slice without it being a fixed size as demonstrated below
	slice2 := []int{10, 20, 30}
	fmt.Println(slice2)
	//append values
	slice2 = append(slice2, 40, 50, 60)
	fmt.Println(slice2)
	// Create Slices out of a Slice example.

	slice4 := slice2[1:] // Creates a slice copying elements of slice2 starting from index 1 e.g(20 30 40 50 60])
	fmt.Println(slice4)

	slice5 := slice2[:3] //Creates a slice copying elements of slice2 starting from index 0 upto but not inclusive of value at index 2 e.g(10 20 30])
	fmt.Println(slice5)

	slice6 := slice2[3:5] //Creates a slice copying elements of slice2 starting from index 3 upto but not inclusive of value at index 5 e.g(40 50])
	fmt.Println(slice6)

	//3. Maps
	fmt.Println("Maps _________________")
	//Maps defined as usual similar to java, specify the key data type and the value data type
	nameAgeMap := map[string]int{"ken": 45}
	fmt.Println(nameAgeMap)
	fmt.Println(nameAgeMap["ken"])
	nameAgeMap["ken"] = 56
	nameAgeMap["bill"] = 45
	fmt.Println(nameAgeMap["ken"])
	fmt.Println(nameAgeMap)
	delete(nameAgeMap, "bill")
	fmt.Println(nameAgeMap)

	//4. Struct

	fmt.Println("Structs _________________")

	type user struct {
		id         int
		firstName  string
		secondName string
		alive      bool
	}

	user1 := user{
		id:         12,
		firstName:  "Mar",
		secondName: "Malang",
		alive:      false,
	}
	fmt.Println(user1)

	var user2 user

	user2.id = 34
	user2.firstName = "Jimmy"
	user2.secondName = "Bullard"
	user2.alive = true
	fmt.Println(user2)

}
