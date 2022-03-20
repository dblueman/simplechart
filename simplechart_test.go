package simplechart

import (
   "testing"
)

const (
   chartWidth = 1024
   font = "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf"
)

func TestPieMap(t *testing.T) {
   sc, err := NewSimpleChart(chartWidth, font)
   if err != nil {
      t.Fatal(err)
   }

   data := map[string]int{"Alice": 10, "Bob": 5, "Charlie": 15}

   err = sc.PieMap("/dev/null", data)
   if err != nil {
      t.Error(err)
   }
}

func TestBarSlice(t *testing.T) {
   sc, err := NewSimpleChart(chartWidth, font)
   if err != nil {
      t.Fatal(err)
   }

   data := []int{10, 5, 15}

   err = sc.BarSlice("/dev/null", data)
   if err != nil {
      t.Error(err)
   }
}

func TestBarMap(t *testing.T) {
   sc, err := NewSimpleChart(chartWidth, font)
   if err != nil {
      t.Fatal(err)
   }

   data := map[int]int{1: 10, 2: 5, 3: 15}

   err = sc.BarMap("/dev/null", data)
   if err != nil {
      t.Error(err)
   }
}