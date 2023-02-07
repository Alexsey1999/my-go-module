package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	fiber2 "github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Println(fiber2.New())
	fmt.Println(fiber.New())
	type Person struct {
		name string
		age  int
	}

	// nums1 := []float64{1.0, 2.0, 3.0}
	// nums2 := []int64{10, 20, 30}

	// fmt.Println(sum(nums1))
	// // fmt.Println(sum(nums2))
	// fmt.Println(sum[int64](nums2))

	names := []string{"Alexsey", "John", "Steve"}
	ages := []int{23, 44, 31}
	persons := []Person{
		Person{
			name: "Alexsey",
			age:  23,
		},
		Person{
			name: "John",
			age:  45,
		},
	}

	fmt.Println(contains(names, "John"))
	fmt.Println(contains(ages, 44))
	fmt.Println(contains(persons, Person{
		name: "Alexsey",
		age:  23,
	}))

	fmt.Println(contains(persons, Person{
		name: "Unknown",
		age:  23,
	}))

	fmt.Println()

	show(1, 2, 3)
	show("a", "b", "c")
	show([]int{10, 20, 30}, []int{40, 50, 60})
	show(persons[0], persons[1])
	show(interface{}(1), interface{}("hello"), any(struct{ name string }{name: "Test name"}))

	var ints Numbers[int64]
	ints = append(ints, []int64{10, 20, 30, 44}...)
	fmt.Println(ints)

	floats := Numbers[float64]{1.0, 20.3, 3}
	fmt.Println(floats)

	print(123, "\n")
	print("\n")
	println(12333)
	println("hellooo")
	// fmt.Println(PrintGreeting())
}

func show[T any](entities ...T) {
	fmt.Println(entities)
}

func contains[T comparable](elems []T, searchable T) bool {
	for _, el := range elems {
		if el == searchable {
			return true
		}
	}

	return false
}

type Number interface {
	~int64 | float64
}

type Numbers[T Number] []T

func sum[V Number](nums []V) V {
	var sumRes V = 0

	for _, n := range nums {
		sumRes += n
	}

	return V(sumRes)
}
