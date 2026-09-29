// Package basic はGoの基礎（変数、制御構文、多値返却、可変長引数、型付き定数）を扱う。
package basic

import "fmt"

// Add はaとbの和を返す。
func Add(a, b int) int {
	return a + b
}

// Sum は与えられた数値すべての合計を返す。
// 引数なしのSum()は0を返す。
func Sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

// FizzBuzz はnが3の倍数なら"Fizz"、5の倍数なら"Buzz"、
// 両方の倍数なら"FizzBuzz"、それ以外はnの10進文字列を返す。
func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}

// DivMod はa / bの商と余りを返す。
// bが0の場合はpanicせずにエラーを返す。
func DivMod(a, b int) (quotient, remainder int, err error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("denominator cannot be zero")
	}
	quotient = a / b
	remainder = a % b
	return quotient, remainder, nil
}

// Season はiotaで定義された季節を表す。
type Season int

const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)

// SeasonName はsに対応する季節名を返す。
// 定義済み定数以外の値には"Unknown"を返す。
func SeasonName(s Season) string {
	switch s {
	case Spring:
		return "Spring"
	case Summer:
		return "Summer"
	case Autumn:
		return "Autumn"
	case Winter:
		return "Winter"
	default:
		return "Unknown"
	}
}
