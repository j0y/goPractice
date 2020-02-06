package main

import (
	"testing"
)

func TestSameAge(t *testing.T) {
	a := Person{"a", 34}
	b := Person{"b", 34}
	got := []Person{a, b}
	want := false

	ans := anyExactlyTwiceAsOld(got)
	if ans != want {
		t.Errorf("got %t, want %t", ans, want)
	}
}

func TestUneven(t *testing.T) {
	a := Person{"a", 89}
	b := Person{"b", 44}
	got := []Person{a, b}
	want := false

	ans := anyExactlyTwiceAsOld(got)
	if ans != want {
		t.Errorf("got %t, want %t", ans, want)
	}
}

func TestCorrect(t *testing.T) {
	a := Person{"a", 21}
	b := Person{"b", 42}
	got := []Person{a, b}
	want := true

	ans := anyExactlyTwiceAsOld(got)
	if ans != want {
		t.Errorf("got %t, want %t", ans, want)
	}
}

func TestMultiple(t *testing.T) {
	a := Person{"a", 21}
	b := Person{"b", 42}
	c := Person{"c", 18}
	d := Person{"d", 2}
	got := []Person{a, b, c, d}
	want := true

	ans := anyExactlyTwiceAsOld(got)
	if ans != want {
		t.Errorf("got %t, want %t", ans, want)
	}
}

func TestAtLeastTwice(t *testing.T) {
	a := Person{"a", 32}
	b := Person{"b", 89}
	c := Person{"c", 18}
	d := Person{"d", 2}
	got := []Person{a, b, c, d}
	want := true

	ans := anyAtLeastTwiceAsOld(got)
	if ans != want {
		t.Errorf("got %t, want %t", ans, want)
	}
}

func TestNotAtLeastTwice(t *testing.T) {
	a := Person{"a", 12}
	b := Person{"b", 10}
	c := Person{"c", 7}
	d := Person{"d", 13}
	got := []Person{a, b, c, d}
	want := false

	ans := anyAtLeastTwiceAsOld(got)
	if ans != want {
		t.Errorf("got %t, want %t", ans, want)
	}
}
