package data

import "testing"

func TestCheckValodation(t *testing.T) {
	p := &Product{
		Name:  "Artem",
		Price: 1.0,
		SKU:   "abs-asd-asd",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
