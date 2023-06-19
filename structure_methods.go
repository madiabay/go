package main

import "fmt"

// type Car struct {
// 	name        string
// 	year        int32
// 	horse_power int32
// }

// func (c Car) get_power() int32 {
// 	return c.horse_power
// }

// func (c *Car) set_name(new_name string) {
// 	c.name = new_name
// }

// func (c Car) IsNewer() bool {
// 	if c.year >= 2010 {
// 		return true
// 	}
// 	return false
// }

// func main() {
// 	mers := Car{name: "Mersedes", year: 2009, horse_power: 260}
// 	fmt.Println(mers)
// 	fmt.Println(mers.get_power())

// 	mers.set_name("Mersedes_S_Class")
// 	fmt.Println(mers)
// 	fmt.Println(mers.IsNewer())
// }

type Score struct {
	name   string
	points []int
}

func (s Score) max() int {
	maximum := 0
	for i := 0; i < len(s.points); i++ {
		if maximum < s.points[i] {
			maximum = s.points[i]
		}
	}
	return maximum
}

func main() {
	madi_maraphon := Score{name: "Madi", points: []int{543, 1, 2, 3, 4}}
	fmt.Println(madi_maraphon)
	fmt.Println(madi_maraphon.max())
}
