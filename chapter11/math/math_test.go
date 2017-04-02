package math

import "testing"

type testpair struct {
  values []float64
  res float64
}

var tests = []testpair {
  { []float64{1,2}, 1.5},
  { []float64{1,1,1,1,1}, 1},
  { []float64{-1, 1}, 0},
  { []float64{}, 0},
}

// func TestAverage(t *testing.T) {
//   var v float64;
//   v = Average([]float64{1,2})
//   if v != 1.5 {
//     t.Error("Expected 1.5, got ", v)
//   }
// }

func TestAverage(t *testing.T) {
  for _, pair := range tests {
    v := Average(pair.values)
    if v != pair.res {
      t.Error(
        "For", pair.values,
        "expected", pair.res,
        "got", v,
      )
    }
  }
}

var minTests = []testpair {
  { []float64{1,2}, 1},
  { []float64{1,1,1,1,1}, 1},
  { []float64{-1, 1}, -1},
  { []float64{}, 0},
}

func TestMin(t *testing.T) {
  for _, pair := range minTests {
    v := Min(pair.values)
    if v != pair.res {
      t.Error(
        "For", pair.values,
        "expected", pair.res,
        "got", v,
      )
    }
  }
}

var maxTests = []testpair {
  { []float64{1,2}, 2},
  { []float64{1,1,1,1,1}, 1},
  { []float64{-1, 1}, 1},
  { []float64{}, 0},
}

func TestMax(t *testing.T) {
  for _, pair := range maxTests {
    v := Max(pair.values)
    if v != pair.res {
      t.Error(
        "For", pair.values,
        "expected", pair.res,
        "got", v,
      )
    }
  }
}
