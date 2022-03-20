package simplechart

import (
   "fmt"
   "os"
   "sort"
   "strconv"
   "time"

   "github.com/golang/freetype/truetype"
   "github.com/tdewolff/canvas/renderers"
   "github.com/wcharczuk/go-chart/v2"
)

type SimpleChart struct {
   width int
   font *truetype.Font
}

func NewSimpleChart(width int, fontName string) (*SimpleChart, error) {
   sc := SimpleChart{width: width}

   if fontName != "" {
      fontData, err := os.ReadFile(fontName)
      if err != nil {
         return nil, err
      }

      sc.font, err = truetype.Parse(fontData)
      if err != nil {
         return nil, err
      }
   }

   return &sc, nil
}

func (sc *SimpleChart) PieMap(fname string, data map[string]int) error {
   var values []chart.Value

   for k, v := range data {
      label := fmt.Sprintf("%s (%d)", k, v)
      values = append(values, chart.Value{Value: float64(v), Label: label})
   }

   graph := chart.DonutChart{
      Width: sc.width,
      Height: sc.width,
      Values: values,
      Font: sc.font,
   }

   f, err := os.Create(fname)
   if err != nil {
      return err
   }
   defer f.Close()

   return graph.Render(renderers.NewGoChart(renderers.PDF()), f)
}

func (sc *SimpleChart) BarSlice(fname string, data []int) error {
   var values []chart.Value

   for i, v := range data {
      label := strconv.Itoa(i)
      values = append(values, chart.Value{Value: float64(v), Label: label})
   }

   graph := chart.BarChart{
      XAxis: chart.Style{
            StrokeColor: chart.ColorBlack,
            StrokeWidth: 1.,
      },
      YAxis: chart.YAxis{
         Style: chart.Style{
            StrokeColor: chart.ColorBlack,
            StrokeWidth: 1.,
         },
      },
      Font: sc.font,
      Width: sc.width,
      Height: sc.width/2,
      BarWidth: sc.width / (len(values)+6),
      Bars: values,
   }

   f, err := os.Create(fname)
   if err != nil {
      return err
   }
   defer f.Close()

   return graph.Render(renderers.NewGoChart(renderers.PDF()), f)
}

func (sc *SimpleChart) BarMap(fname string, data map[int]int) error {
   keys := make([]int, 0, len(data))
   for k := range data {
      keys = append(keys, k)
   }

   sort.Ints(keys)
   var values []chart.Value

   for _, key := range keys {
      label := time.Unix(int64(key)*60*60*24, 0).Format("Mon 02/01")
      values = append(values, chart.Value{Value: float64(data[key]), Label: label})
   }

   graph := chart.BarChart{
      XAxis: chart.Style{
            StrokeColor: chart.ColorBlack,
            StrokeWidth: 1.,
      },
      YAxis: chart.YAxis{
         Style: chart.Style{
            StrokeColor: chart.ColorBlack,
            StrokeWidth: 1.,
         },
      },
      Font: sc.font,
      Width: sc.width,
      Height: sc.width/2,
      BarWidth: sc.width / (len(values)+6),
      Bars: values,
   }

   f, err := os.Create(fname)
   if err != nil {
      return err
   }
   defer f.Close()

   return graph.Render(renderers.NewGoChart(renderers.PDF()), f)
}
