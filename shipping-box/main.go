package main

// based on the area of boxes and products
// 1º loop: got boxes that could contain the products
// 2º loop: got best box from valid boxes

import (
	"fmt"
	"sync"
)

// Product contains products info
type Product struct {
	Name string
	Len  int // milimiters
	Wid  int
	Hei  int
}

// Area returns the product area
func (p *Product) Area() int {
	return p.Len * p.Hei * p.Wid
}

// Box contains products info
type Box struct {
	Len int // milimiters
	Wid int
	Hei int
}

// Area returns the box area
func (b *Box) Area() int {
	return b.Len * b.Hei * b.Wid
}

// ResponseCheck stores information to comunicate the concurrence process
type ResponseCheck struct {
	box   Box
	valid bool
}

// func getBestBox(availableBoxes []Box, products []Product) Box {
func getBestBox(availableBoxes []Box, products []Product) Box {
	outputResponseCheck := make(chan ResponseCheck)
	outputValidBoxes := make(chan []Box)
	var wg sync.WaitGroup
	var bestBox Box

	go handleCheckProductsBox(&wg, outputResponseCheck, outputValidBoxes)

	for _, box := range availableBoxes {
		wg.Add(1)
		go checkProductsBox(box, products, outputResponseCheck)
	}

	wg.Wait()
	close(outputResponseCheck)

	validBoxes := <-outputValidBoxes

	for _, box := range validBoxes {
		if (bestBox == Box{}) {
			bestBox = box
			continue
		}

		if box.Area() < bestBox.Area() {
			bestBox = box
		}
	}

	return bestBox
}

func handleCheckProductsBox(wg *sync.WaitGroup, input chan ResponseCheck, output chan []Box) {
	var results []Box

	for result := range input {
		if result.valid {
			results = append(results, result.box)
		}
		wg.Done()
	}

	output <- results
}

func checkProductsBox(box Box, products []Product, output chan ResponseCheck) {
	area := 0

	for _, product := range products {
		area += product.Area()
	}

	if area <= box.Area() {
		output <- ResponseCheck{
			box:   box,
			valid: true,
		}
	} else {
		output <- ResponseCheck{
			box:   box,
			valid: false,
		}
	}
}

func main() {

	boxes := make([]Box, 0)
	boxes = append(boxes, Box{
		Len: 20,
		Wid: 10,
		Hei: 10,
	})
	boxes = append(boxes, Box{
		Len: 40,
		Wid: 40,
		Hei: 10,
	})
	products := make([]Product, 0)
	products = append(products, Product{
		Name: "p1",
		Len:  10,
		Wid:  10,
		Hei:  10,
	})
	products = append(products, Product{
		Name: "p2",
		Len:  10,
		Wid:  10,
		Hei:  11,
	})

	fmt.Println(fmt.Sprintf("best box for products [%v]: %v", products, getBestBox(boxes, products)))
}
