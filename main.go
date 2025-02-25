package main

import (
	"iter"
)

func main() {
}

func GenMap() map[string]any {
	mp := map[string]any{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}
	return mp
}

func GenIterMap(yield func()) iter.Seq[string] {

}
