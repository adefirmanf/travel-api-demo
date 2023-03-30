// Package inmem ...
// Macro for date helper helped by https://github.com/thedevsaddam/gojsonq/issues/31
package inmem

import (
	"fmt"
	"time"
)

const layout = "2006-01-02"

func dateLessOrEqualTo(x, y interface{}) (bool, error) { // date less than or equal query function
	xs, okx := x.(string)
	ys, oky := y.(string)
	if !okx || !oky {
		return false, fmt.Errorf("date support for string only")
	}

	t1, _ := time.Parse(layout, xs)
	t2, _ := time.Parse(layout, ys)

	return t1.Unix() <= t2.Unix(), nil
}

func dateGreaterOrEqualTo(x, y interface{}) (bool, error) { // date greater than or equal query function
	xs, okx := x.(string)
	ys, oky := y.(string)
	if !okx || !oky {
		return false, fmt.Errorf("date support for string only")
	}

	t1, _ := time.Parse(layout, xs) // TODO: check these error too
	t2, _ := time.Parse(layout, ys)

	return t1.Unix() >= t2.Unix(), nil
}
