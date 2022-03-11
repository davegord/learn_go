package dog

import (
	"fmt"
	"testing"

	"github.com/davegord/learn_go/ninja-level-13/01/dog"
)

func TestYears(t *testing.T) {
	got := dog.Years(10)
	if got != 70 {
		t.Error("Expected 70, got", got)
	}
}

func TestYearsTwo(t *testing.T) {
	got := dog.YearsTwo(10)
	if got != 70 {
		t.Error("Expected 70, got", got)
	}

}

func ExampleTestYears() {
	fmt.Println(dog.Years(10))
	// Output:
	// 70
}

func ExampleTestYearsTwo() {
	fmt.Println(dog.YearsTwo(10))
	// Output:
	// 70
}

func BenchmarkTestYears(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dog.Years(10)
	}

}

func BenchmarkTestYearsTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dog.YearsTwo(10)
	}

}
