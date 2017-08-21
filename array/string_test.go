package array_test

import (
	"testing"
	"strconv"
	"fmt"
	"github.com/RomanosTrechlis/Tools/array"
)

func TestArrayString(t *testing.T) {
	a := make(array.StringArray, 10)
	for i := 0; i < 10; i++ {
		a[i] = strconv.Itoa(i + 1)
	}

	if fmt.Sprintf("%s", a.Head()) != "1" {
		t.Errorf("expected '1', got %s", a.Head())
	}
	if fmt.Sprintf("%T", a.Head()) != "string" {
		t.Errorf("expected 'string', got %T", a.Head())
	}

	if fmt.Sprintf("%s", a.Tail()) != "[2 3 4 5 6 7 8 9 10]" {
		t.Errorf("expected '[2 3 4 5 6 7 8 9 10]', got %s", a.Tail())
	}
	if fmt.Sprintf("%T", a.Tail()) != "array.StringArray" {
		t.Errorf("expected 'array.StringArray', got %T", a.Tail())
	}
}


