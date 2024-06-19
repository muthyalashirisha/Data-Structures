package main

import (
	"fmt"
)

// type MySlice []string

// func Clone1[E any](s []E) []E {
// 	s1 := []E{}
// 	return append(s1, s...)
// }
// func Clone2[S ~[]E, E any](s S) S {
// 	s1 := []E{}
// 	return append(s1, s...)
// }
// func (ms MySlice) String() string {
// 	return strings.Join(ms, "+")
// }

// func printMySlice(ms MySlice) {
// 	c := Clone2(ms)
// 	slices.Sort(c)
// 	// The Go assignment rules allow us to pass a value of type MySlice to a parameter of type []string, so calling Clone1 is fine. But Clone1 will return a value of type []string, not a value of type MySlice. The type []string doesn’t have a String method, so the compiler reports an error.
// 	// fmt.Println(c.String())
// }
// func main4() {
// 	slice := [][]int{{1, 3, 3, 4, 2, 6}, {1}, {2}, {3}, {4}}
// 	fmt.Println(slice)
// 	slice1 := slices.Clone(slice)
// 	sliceCopy := make([][]int, len(slice))
// 	copy(sliceCopy, slice)

// 	fmt.Println(slice1, sliceCopy)
// 	fmt.Println(slice1[:1], "test")
// 	slice[0][4] = 23
// 	fmt.Println(slice1, slice, sliceCopy)
// 	slice2 := slices.Delete(slice, 0, 1)
// 	fmt.Println(slice, slice2, slice1, sliceCopy)

// 	s := make([]int, 3)
// 	fmt.Println(len(s), cap(s))
// 	s1 := []string{"a", "b", "c"}
// 	s2 := s1[:]
// 	fmt.Printf("%T, %T", s1, s2)

// 	o := [8]int{1, 4, 2, 8, 4, 1, 0, 8}
// 	o1 := o[2:]
// 	o2 := o[4:]
// 	fmt.Println(o1, o2, o)

// 	copy(o2, o1)
// 	fmt.Println(o1, o2, o, len(o2), cap(o2))
// 	o2 = append(o2, o1...)
// 	fmt.Println(o1, o2, o, len(o2), cap(o2))

// 	main3()
// 	var prints []func()
// 	for i := 1; i <= 3; i++ {
// 		prints = append(prints, func() { fmt.Println(i) })
// 	}
// 	for _, print := range prints {
// 		print()
// 	}
// 	done := make(chan bool)

// 	values := []string{"a", "b", "c"}
// 	for _, v := range values {
// 		go func() {
// 			fmt.Println(v)
// 			done <- true
// 		}()
// 	}

// 	// wait for all goroutines to complete before exiting
// 	for _ = range values {
// 		<-done
// 	}

// }
// func main3() {
// 	type Person struct {
// 		Name string
// 		Age  int
// 	}
// 	people := []Person{
// 		{"Alice", 55},
// 		{"Bob", 24},
// 		{"Gopher", 13},
// 	}
// 	n, found := slices.BinarySearchFunc(people, Person{"Bob", 0}, func(a, b Person) int {
// 		return cmp.Compare(a.Name, b.Name)
// 	})
// 	fmt.Println("Bob:", n, found)
// }
// func sample(s string) bool {
// 	revStr := reverseString(s)

// 	for i := 0; i < len(s)-1; i++ {
// 		if strings.Contains(revStr, string(s[i:i+2])) {
// 			return true
// 		}
// 	}
// 	return false
// }
// func reverseString(s string) string {
// 	reversed := make([]rune, len(s))
// 	for i, r := range s {
// 		reversed[len(s)-1-i] = r
// 	}

// 	return string(reversed)
// }
// func minimumDeletions(word string, k int) int {
// 	freq := map[byte]int{}
// 	for i := 0; i < len(word); i++ {
// 		freq[word[i]]++
// 	}
// 	type freqStruct struct {
// 		Key   byte
// 		Value int
// 	}
// 	var freqSlice []freqStruct
// 	for k, v := range freq {
// 		freqSlice = append(freqSlice, freqStruct{k, v})
// 	}

// 	// Sort the slice by frequency values
// 	sort.Slice(freqSlice, func(i, j int) bool {
// 		return freqSlice[i].Value < freqSlice[j].Value
// 	})

// 	// Count the minimum number of decreases needed
// 	decreases := 0
// 	for i := 1; i < len(freqSlice); i++ {
// 		diff := freqSlice[i].Value - freqSlice[i-1].Value
// 		if diff > k {
// 			// Calculate the number of decreases needed
// 			decreases += (diff - k)
// 			// Update the frequency to make the difference <= k
// 			freqSlice[i].Value -= (diff - k)
// 		}
// 	}

// 	return decreases
// }

func main2() {
	var n int
	fmt.Scan(&n)                   // Number of prisoners
	prisoners := make([][2]int, n) // Slice to store coordinates of prisoners

	// Read coordinates of prisoners
	for i := 0; i < n; i++ {
		fmt.Scan(&prisoners[i][0], &prisoners[i][1])
	}

	// Find the escaped prisoner's coordinate
	escapedPrisoner := findEscapedPrisoner(prisoners)
	fmt.Println("Escaped Prisoner's Coordinate:", escapedPrisoner[0], ",", escapedPrisoner[1])
}

// Function to find the escaped prisoner's coordinate
func findEscapedPrisoner(prisoners [][2]int) [2]int {
	ans := [2]int{-1, -1}
	// Iterate through each prisoner's coordinate pair and check conditions
	for _, prisoner := range prisoners {
		x := prisoner[0]
		y := prisoner[1]
		condition1 := false
		condition2 := false

		// Check if there is at least one prisoner at [x, ay] where ay != y
		for _, pair := range prisoners {
			x1 := pair[0]
			y1 := pair[1]
			if y1 != y && x1 == x {
				condition1 = true
				break
			}
		}

		// Check if there is at least one prisoner at [bx, y] where bx != x
		// Check if there is at least one prisoner at [x, ay] where ay != y
		for _, pair := range prisoners {
			x1 := pair[0]
			y1 := pair[1]
			if y1 == y && x1 != x {
				condition2 = true
				break
			}
		}

		// If any condition is not met, return the coordinates of this prisoner
		if !condition1 || !condition2 {
			ans = [2]int{x, y}
			break
		}
	}

	return ans // If no escaped prisoner found
}
