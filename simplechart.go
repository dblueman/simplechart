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

type XYValue struct {
   X, Y float64
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

func (sc *SimpleChart) BarMap(fname string, data map[int]int, date bool) error {
   keys := make([]int, 0, len(data))
   for k := range data {
      keys = append(keys, k)
   }

   sort.Ints(keys)
   var values []chart.Value

   for _, key := range keys {
      value := chart.Value{Value: float64(data[key])}

      if date {
         value.Label = time.Unix(int64(key)*60*60*24, 0).Format("Mon 02/01")
      } else {
         value.Label = strconv.Itoa(key)
      }

      values = append(values, value)
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

func (sc *SimpleChart) Line(fname string, xlabel string, ylabel string, datas [][]XYValue) error {
   graph := chart.Chart{
      XAxis: chart.XAxis{
			Name: xlabel,
		},
      YAxis: chart.YAxis{
			Name: ylabel,
		},
   }

   for _, data := range datas {
      series := chart.ContinuousSeries{}

      for _, d := range data {
         series.XValues = append(series.XValues, d.X)
         series.YValues = append(series.YValues, d.Y)
      }

      graph.Series = append(graph.Series, series)
   }

   f, err := os.Create(fname)
   if err != nil {
      return err
   }
   defer f.Close()

   return graph.Render(renderers.NewGoChart(renderers.PDF()), f)
}
