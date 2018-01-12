package bitvector

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)

	if s.Has(1) == false || s.Has(2) == false || s.Has(3) == false || s.Has(4) == false {
		t.Error("Failed")
	}

	if s.Has(0) == true || s.Has(5) == true || s.Has(100) == true {
		t.Error("Failed")
	}
}
