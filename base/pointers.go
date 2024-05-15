package main

import "fmt"

type House struct {
	Material     string
	HasFireplace bool
	Floors       int
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("ptr:", &i)

	h1 := &House{
		Material:     "clay",
		HasFireplace: true,
		Floors:       3,
	}
	h2 := &House{
		Material:     "steel",
		HasFireplace: false,
		Floors:       2,
	}

	fmt.Printf("h1 details:   %+v\n", h1)
	fmt.Printf("h2 details:   %+v\n", h2)

}
