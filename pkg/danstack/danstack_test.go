package danstack_test

import (
	"testing"
	"github.com/danwhitford/danstack/pkg/danstack"
)

func assertEqual(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Fatalf("Wanted: %d. Got: %d", expected, actual)
	}
}

func assertEqualStr(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Fatalf("Wanted: `%s`. Got: `%s`", expected, actual)
	}
}

func assertNil(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

// Test can initialise a new stack
func TestInit(t *testing.T) {
	danstack.New[int]()
}

// Test can push to the stack without issue
func TestPushSimple(t *testing.T) {
	s := danstack.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
}

// Test we can pop what we have just pushed
func TestPopSimple(t *testing.T) {
	s := danstack.New[int]()
	s.Push(10)
	v, err := s.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if v != 10 {
		t.Fatalf("Wanted: %d. Got: %d", 10, v)
	}
}

// Test order of popping is correct
func TestPopOrder(t *testing.T) {
	s := danstack.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	v, err := s.Pop()
	assertNil(t, err)
	assertEqual(t, 3, v)

	v, err = s.Pop()
	assertNil(t, err)
	assertEqual(t, 2, v)

	v, err = s.Pop()
	assertNil(t, err)
	assertEqual(t, 1, v)
}

// Test error for underflow
func TestPopUnderflow(t *testing.T) {
	s := danstack.New[int]()
	v, err := s.Pop()
	if err == nil {
		t.Fatalf("expected error but was %d", v)
	}
}

// Test that the slice expands properly
func TestSliceExpand(t *testing.T) {
	s := danstack.New[int]()
	for i := 0; i < 1000; i++ {
		s.Push(i)
	}

	for i := 999; i >= 0; i-- {
		v, err := s.Pop()
		assertNil(t, err)
		assertEqual(t, i, v)
	}
}

// Test pushing and popping unevenly
func TestPushingPopping(t *testing.T) {
	s := danstack.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	v, err := s.Pop()
	assertNil(t, err)
	assertEqual(t, 5, v)

	v, err = s.Pop()
	assertNil(t, err)
	assertEqual(t, 4, v)

	s.Push(6)
	s.Push(7)

	v, err = s.Pop()
	assertNil(t, err)
	assertEqual(t, 7, v)

	v, err = s.Pop()
	assertNil(t, err)
	assertEqual(t, 6, v)
}

func TestEmpty(t *testing.T) {
	s := danstack.New[int]()
	if !s.Empty() {
		t.Fatal("Should be empty but isn't!")
	}
	s.Push(1)
	if s.Empty() {
		t.Fatal("Should not be empty but is!")
	}
	s.Pop()
	if !s.Empty() {
		t.Fatal("Should be empty but isn't!")
	}
}

func FuzzPushing(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	s := danstack.New[string]()
	f.Fuzz(func(t *testing.T, orig string) {
		s.Push(orig)
		v, err := s.Pop()
		assertNil(t, err)
		assertEqualStr(t, orig, v)
	})
}

func FuzzSizes(f *testing.F) {
	testcases := []int{1, 10, 100}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, j int) {
		s := danstack.New[int]()
		for i := 0; i < j; i++ {
			s.Push(i)
			v, err := s.Pop()
			assertNil(t, err)
			assertEqual(t, i, v)
		}
	})
}
