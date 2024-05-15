package main

import "fmt"

/*
 *   reference:
 *     https://www.sohamkamani.com/golang/options-pattern
 */

/*
 * Define a function type that accepts a pointer to a House.
 * This is the signature of our option functions
 */
type HouseOption func(*House)

type House struct {
	Material     string
	HasFireplace bool
	Floors       int
}

/* Define functional options that modify the *House instance.
 * Each of the below functions are "option constructors" and return another function that takes a *House as an argument & returns nothing
 * In the case of the last functional option, notice that it even accepts a parameter.
 */
func WithConcrete() HouseOption {
	return func(h *House) {
		h.Material = "concrete"
	}
}
func WithoutFireplace() HouseOption {
	return func(h *House) {
		h.HasFireplace = false
	}
}
func WithFloors(floors int) HouseOption {
	return func(h *House) {
		h.Floors = floors
	}
}

// This signature type allows for invoking function directly on object instance
func (houseP *House) setHouseMaterialToSteel() {
	houseP.Material = "steel"
}

func setHouseMaterialToClay(houseP *House) {
	houseP.Material = "Clay"
}

// NewHouse is a constructor function for `*House`
// This constructor is a variadic function that accepts a list of any number of function option arguments.
// It then applies to the *House instance before returning it
func NewHouse(opts ...HouseOption) *House {
	const (
		defaultFloors       = 2
		defaultHasFireplace = true
		defaultMaterial     = "wood"
	)

	h := &House{
		Material:     defaultMaterial,
		HasFireplace: defaultHasFireplace,
		Floors:       defaultFloors,
	}

	// Loop through each option

	for _, opt := range opts {
		// call the option giving the instantiated *House as the argument
		opt(h)
	}

	return h
}

func main() {
	h := NewHouse(
		WithConcrete(),
		WithoutFireplace(),
		WithFloors(3),
	)
	//h.setHouseMaterialToSteel()
	setHouseMaterialToClay(h)
	fmt.Printf("house details:   %+v\n", h)

}
