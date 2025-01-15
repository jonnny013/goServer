package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	return Point{}
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / (float64(t.Second())))
}
