package sliceDemo

import "fmt"

// Diffreent way to create a slices
var myList = make([]string, 5)
var fruitList = []string{"Orange"}
var sliceCon []int

func main() {

	//declaring the slice in one way
	fruitList = append(fruitList, "Banana")

	//declaring the slice in second way
	var vegList []string
	vegList = append(vegList, "Tomatto")
	vegList = append(vegList, "Carot")

	myList = append(vegList, "many")
	fmt.Print(myList)
	myList[0] = "sswsw"
	fmt.Print(myList)

	for i := 0; i < 100; i++ {
		sliceCon = append(sliceCon, i+1)
	}
	fmt.Println("iterating a slice")
	for i, v := range sliceCon {
		fmt.Println(i, v)
	}

	var nameSlice = []string{
		"darshan",
		"Ramesh",
		"brinda",
		"Varun",
		"Naveen",
		"Hemanth",
	}

	var copySlice = nameSlice[0:3]
	fmt.Println("Copied slice is ", copySlice)

	//Copy All the elements
	var allCopy = nameSlice[:]
	fmt.Println("All copied elements", allCopy)

	//Copy the elements upto 3 elements from starting index
	var uptoCopy = nameSlice[:3]
	fmt.Println("Copied upto 3 elements from starting", uptoCopy)

}
