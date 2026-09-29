package basic

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{-1, 1, 0},
		{0, 0, 0},
	}
	for _, c := range cases {
		if got := Add(c.a, c.b); got != c.want {
			t.Errorf("Add(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestSum(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{nil, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3}, 6},
	}
	for _, c := range cases {
		if got := Sum(c.nums...); got != c.want {
			t.Errorf("Sum(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		n    int
		want string
	}{
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{7, "7"},
	}
	for _, c := range cases {
		if got := FizzBuzz(c.n); got != c.want {
			t.Errorf("FizzBuzz(%d) = %q, want %q", c.n, got, c.want)
		}
	}
}

func TestDivMod(t *testing.T) {
	q, r, err := DivMod(7, 2)
	if err != nil {
		t.Fatalf("DivMod(7, 2) returned unexpected error: %v", err)
	}
	if q != 3 || r != 1 {
		t.Errorf("DivMod(7, 2) = (%d, %d), want (3, 1)", q, r)
	}

	_, _, err = DivMod(1, 0)
	if err == nil {
		t.Fatal("DivMod(1, 0) should return an error")
	}
}

func TestSeasonName(t *testing.T) {
	cases := []struct {
		s    Season
		want string
	}{
		{Spring, "Spring"},
		{Summer, "Summer"},
		{Autumn, "Autumn"},
		{Winter, "Winter"},
		{Season(99), "Unknown"},
	}
	for _, c := range cases {
		if got := SeasonName(c.s); got != c.want {
			t.Errorf("SeasonName(%d) = %q, want %q", c.s, got, c.want)
		}
	}
}
