package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	name string
	age  int
}

func NewPerson(name string) Person {
	p := Person{name: name}
	p.age = rand.Intn(120)
	return p
}

func anyExactlyTwiceAsOld(people []Person) bool {
	var doubleAges []int
	var regularAges []int
	for i := range people {
		doubleAges = append(doubleAges, people[i].age*2)
		regularAges = append(regularAges, people[i].age)
	}
	var xors []int
	xors = xor(doubleAges, regularAges)
	if len(xors) < len(regularAges)*2 {
		return true
	}
	return false
}

func anyAtLeastTwiceAsOld(people []Person) bool {
	minAgeIndex := 0
	for i := range people {
		if people[i].age < people[minAgeIndex].age {
			minAgeIndex = i
		}
	}
	for i := range people {
		if minAgeIndex != i && people[i].age >= people[minAgeIndex].age*2 {
			return true
		}
	}
	return false
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var people []Person
	for i := 0; i < rand.Intn(20)+2; i++ {
		name := string('a' + byte(i))
		people = append(people, NewPerson(name))
	}

	fmt.Println(people)
	if anyExactlyTwiceAsOld(people) {
		fmt.Println("Some people are exactly twice as old")
	}
	if anyAtLeastTwiceAsOld(people) {
		fmt.Println("Some people are at least twice as old")
	}
}

func xor(list1, list2 []int) []int {
	set1 := make(map[int]bool)
	for _, s := range list1 {
		set1[s] = true
	}
	set2 := make(map[int]bool)
	for _, s := range list2 {
		set2[s] = true
	}

	var c []int
	for _, s := range list1 {
		if !set2[s] {
			c = append(c, s)
		}
	}
	for _, s := range list2 {
		if !set1[s] {
			c = append(c, s)
		}
	}
	return c
}
