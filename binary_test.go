package main

import (
	"fmt"
	"os"
	"testing"
)

func TestBinaryPrint(t *testing.T) {
	binaries := []Binary{
		{decimal: 0, expected: "00000000"},
		{decimal: 1, expected: "00000001"},
		{decimal: 2, expected: "00000010"},
		{decimal: 3, expected: "00000011"},
		{decimal: 8, expected: "00001000"},
		{decimal: 10, expected: "00001010"},
		{decimal: 16, expected: "00010000"},
	}
	for _, binary := range binaries {
		testBinary(t, binary)
	}
}

func TestBitwiseOperations(t *testing.T) {
	b170 := Binary{
		decimal: 170, expected: "10101010",
	}
	b15 := Binary{
		decimal: 15, expected: "00001111",
	}
	b170Or15 := Binary{
		decimal: 170 | 15, expected: "10101111",
	}
	b170And15 := Binary{
		decimal: 170 & 15, expected: "00001010",
	}
	testBinary(t, b170)
	testBinary(t, b15)
	testBinary(t, b170Or15)
	testBinary(t, b170And15)
}

func TestFileOs(t *testing.T) {
	fmt.Printf("O_RDONLY: \t%08b - %d\n", os.O_RDONLY, os.O_RDONLY)
	fmt.Printf("O_WRONLY: \t%08b - %d\n", os.O_WRONLY, os.O_WRONLY)
	fmt.Printf("O_RDWR: \t%08b - %d\n", os.O_RDWR, os.O_RDWR)
	fmt.Println()
	fmt.Printf("O_APPEND: \t%012b - %d\n", os.O_APPEND, os.O_APPEND)
	fmt.Printf("O_SYNC: \t%012b - %d\n", os.O_SYNC, os.O_SYNC)
	fmt.Printf("O_CREATE: \t%012b - %d\n", os.O_CREATE, os.O_CREATE)
	fmt.Printf("O_TRUNC: \t%012b - %d\n", os.O_TRUNC, os.O_TRUNC)
	fmt.Printf("O_EXCL: \t%012b - %d\n", os.O_EXCL, os.O_EXCL)
	fmt.Println()
	fmt.Printf("O_WRONLY | O_APPEND: \t%012b\n", os.O_WRONLY|os.O_APPEND)
}

type Binary struct {
	decimal  int
	expected string
}

func testBinary(t *testing.T, binary Binary) {
	actual := fmt.Sprintf("%08b", binary.decimal)
	if actual != binary.expected {
		t.Errorf("Binary of %d is expected to be %s, but was %s", binary.decimal, binary.expected, actual)
	}
}
