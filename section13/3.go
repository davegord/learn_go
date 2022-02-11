package main

import "fmt"

type vehicle struct {
	door  int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	lexusES := sedan{
		vehicle: vehicle{
			door:  4,
			color: "white",
		},
		luxury: true,
	}

	fordF150 := truck{
		vehicle: vehicle{
			door:  2,
			color: "black",
		},
		fourWheel: true,
	}

	fmt.Println(lexusES)
	fmt.Printf("\t%v\n", lexusES.color)
	fmt.Println(fordF150)
}
