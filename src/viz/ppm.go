package viz

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

const MaxColorValue = 255

func CanvasToPPM(c Canvas) string {
	res := []string{"P3", fmt.Sprintf("%d %d", c.Width, c.Height), fmt.Sprint(MaxColorValue)}
	rows := sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(c.Height)
	for i := 0; i < c.Height; i++ {
		rowI := i
		go func() {
			row := []string{}
			var curStr string
			for j := 0; j < c.Width; j++ {
				r, g, b := ppmScaledColor(c.Pixel(j, rowI))
				for _, v := range []int{r, g, b} {
					vStr := fmt.Sprint(v)
					if len(curStr)+len(vStr) > 70 {
						row = append(row, strings.TrimSpace(curStr))
						curStr = vStr + " "
					} else {
						curStr += vStr + " "
					}
				}
			}
			row = append(row, strings.TrimSpace(curStr))
			rows.Store(rowI, row)
			wg.Done()
		}()
	}
	wg.Wait()
	for i := 0; i < c.Height; i++ {
		v, ok := rows.Load(i)
		if ok {
			if row, ok := v.([]string); ok {
				for _, s := range row {
					res = append(res, s)
				}
			}
		}
		if !ok {
			log.Fatal("Map read failed for snapshot", i)
		}
	}
	return strings.Join(res, "\n") + "\n"
}

func ppmScaledColor(c Color) (int, int, int) {
	return ScaledColorValue256(c.R()), ScaledColorValue256(c.G()), ScaledColorValue256(c.B())
}
